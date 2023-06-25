package compat

import (
	"errors"
	"time"

	"github.com/vmihailenco/msgpack/v5"
)

// Since msgpack decodes an int and json gives a time.Time...
//
// Timestamp decodes json -> time.Time and messagepack through a custom
// unmarshaler to time.Time because modern problems require modern solutions.
type Timestamp struct {
	time.Time
}

func (t *Timestamp) UnmarshalMsgpack(b []byte) error {
	var i int64

	err := msgpack.Unmarshal(b, &i)

	if err != nil {
		// Try to decode using normal
		var ts time.Time

		err1 := msgpack.Unmarshal(b, &ts)

		if err != nil {
			return errors.New("failed to unmarshal msgpack: " + err.Error() + " and " + err1.Error())
		}

		t.Time = ts

		return nil
	}

	t.Time = time.UnixMilli(i)

	return nil
}
