// Note that while types.go is autogenerated, this file is NOT and is manually maintained
//
// Revolt does not provide bitflags so this file is used to provide a similar interface
package flags

type UserFlag uint64

const (
	// User has been suspended from the platform
	SUSPENDED_UserFlag = 1
	// User has deleted their account
	DELETED_UserFlag = 2
	// User was banned off the platform
	BANNED_UserFlag = 4
	// User was marked as spam and removed from platform
	SPAM_UserFlag = 8
)

type BotFlag = int64

const (
	VERIFIED_BotFlag BotFlag = 1
	OFFICIAL_BotFlag BotFlag = 2
)

type ServerFlag = int64

const (
	VERIFIED_ServerFlag ServerFlag = 1
	OFFICIAL_ServerFlag ServerFlag = 2
)

// Returns whether the user has the given flag
//
// E.g. HasFlag(user.Flags, SUSPENDED_UserFlag)
func HasFlag(flags int64, flag int64) bool {
	return flags&flag != 0
}
