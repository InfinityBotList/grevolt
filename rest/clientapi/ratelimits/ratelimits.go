// Ratelimits impl from https://github.com/bwmarrin/discordgo/blob/master/ratelimit.go
package ratelimits

import (
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
)

// customRateLimit holds information for defining a custom rate limit
type CustomRateLimit struct {
	Suffix   string
	Requests int
	Reset    time.Duration
}

// RateLimiter holds all ratelimit buckets
type RateLimiter struct {
	sync.Mutex
	Global           *int64
	Buckets          map[string]*Bucket
	GlobalRateLimit  time.Duration
	CustomRateLimits []*CustomRateLimit
	Logger           *zap.SugaredLogger
}

// NewRatelimiter returns a new RateLimiter
//
// Be sure to set Logger to a valid logger
func NewRatelimiter() *RateLimiter {
	return &RateLimiter{
		Buckets: make(map[string]*Bucket),
		Global:  new(int64),
		/*CustomRateLimits: []*CustomRateLimit{
			{
				suffix:   "//reactions//",
				requests: 1,
				reset:    200 * time.Millisecond,
			},
		},*/
	}
}

// GetBucket retrieves or creates a bucket
func (r *RateLimiter) GetBucket(pkey string) *Bucket {
	r.Lock()
	defer r.Unlock()

	// Buckets on revolt are set based on prefix, so
	bucketName := strings.Split(pkey, "/")

	if len(bucketName) < 1 {
		pkey = "global"
	}

	key := bucketName[0]

	if bucket, ok := r.Buckets[key]; ok {
		return bucket
	}

	b := &Bucket{
		Remaining: 1,
		Key:       key,
		Global:    r.Global,
	}

	// Check if there is a custom ratelimit set for this bucket ID.
	for _, rl := range r.CustomRateLimits {
		if strings.HasSuffix(b.Key, rl.Suffix) {
			b.CustomRateLimit = rl
			break
		}
	}

	r.Logger.Debug("GetBucket: ", key)

	r.Buckets[key] = b
	return b
}

// GetWaitTime returns the duration you should wait for a Bucket
func (r *RateLimiter) GetWaitTime(b *Bucket, minRemaining int) time.Duration {
	// If we ran out of calls and the reset time is still ahead of us
	// then we need to take it easy and relax a little
	if b.Remaining < minRemaining && b.Reset.After(time.Now()) {
		return time.Until(b.Reset)
	}

	// Check for global ratelimits
	sleepTo := time.Unix(0, atomic.LoadInt64(r.Global))
	if now := time.Now(); now.Before(sleepTo) {
		return sleepTo.Sub(now)
	}

	return 0
}

// LockBucket Locks until a request can be made
func (r *RateLimiter) LockBucket(bucketID string) *Bucket {
	return r.LockBucketObject(r.GetBucket(bucketID))
}

// LockBucketObject Locks an already resolved bucket until a request can be made
func (r *RateLimiter) LockBucketObject(b *Bucket) *Bucket {
	b.Lock()

	if wait := r.GetWaitTime(b, 1); wait > 0 {
		r.Logger.Info("LockBucketObject: ", b.Key, " waiting ", wait)
		time.Sleep(wait)
	}

	b.Remaining--
	return b
}

// Bucket represents a ratelimit bucket, each bucket gets ratelimited individually (-global ratelimits)
type Bucket struct {
	sync.Mutex
	Key       string
	Remaining int
	Limit     int
	Reset     time.Time
	Global    *int64

	LastReset       time.Time
	CustomRateLimit *CustomRateLimit
	Userdata        interface{}
}

// Release unlocks the bucket and reads the headers to update the buckets ratelimit info
// and locks up the whole thing in case if there's a global ratelimit.
func (b *Bucket) Release(headers http.Header) error {
	defer b.Unlock()

	// Check if the bucket uses a custom ratelimiter
	if rl := b.CustomRateLimit; rl != nil {
		if time.Since(b.LastReset) >= rl.Reset {
			b.Remaining = rl.Requests - 1
			b.LastReset = time.Now()
		}
		if b.Remaining < 1 {
			b.Reset = time.Now().Add(rl.Reset)
		}
		return nil
	}

	if headers == nil {
		return nil
	}

	remaining := headers.Get("X-RateLimit-Remaining")
	reset := headers.Get("X-RateLimit-Reset") // Discord only, but should be ignored
	global := headers.Get("X-RateLimit-Global")
	resetAfter := headers.Get("X-RateLimit-Reset-After")

	// Update global and per bucket reset time if the proper headers are available
	// If global is set, then it will block all buckets until after Retry-After
	// If Retry-After without global is provided it will use that for the new reset
	// time since it's more accurate than X-RateLimit-Reset.
	// If Retry-After after is not proided, it will update the reset time from X-RateLimit-Reset
	if resetAfter != "" {
		parsedAfter, err := strconv.ParseFloat(resetAfter, 64)
		if err != nil {
			return err
		}

		whole, frac := math.Modf(parsedAfter)
		resetAt := time.Now().Add(time.Duration(whole) * time.Millisecond).Add(time.Duration(frac*1000) * time.Microsecond)

		// Lock either this single bucket or all buckets
		if global != "" {
			atomic.StoreInt64(b.Global, resetAt.UnixNano())
		} else {
			b.Reset = resetAt
		}
	} else if reset != "" {
		// Calculate the reset time by using the date header returned from revolt
		t, err := http.ParseTime(headers.Get("Date"))
		if err != nil {
			return err
		}

		unix, err := strconv.ParseFloat(reset, 64)
		if err != nil {
			return err
		}

		// Calculate the time until reset and add it to the current local time
		// some extra time is added
		//
		// as revolt experiences constant ddos, we set a higher reset time
		whole, frac := math.Modf(unix)
		delta := time.Unix(int64(whole), 0).Add(time.Duration(frac*1000)*time.Millisecond).Sub(t) + time.Millisecond*325
		b.Reset = time.Now().Add(delta)
	}

	// Update remaining if header is present
	if remaining != "" {
		parsedRemaining, err := strconv.ParseInt(remaining, 10, 32)
		if err != nil {
			return err
		}
		b.Remaining = int(parsedRemaining)
	}

	return nil
}
