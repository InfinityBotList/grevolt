package types

// RelationshipStatus : User's relationship with another user (or themselves)
type RelationshipStatus string

// List of RelationshipStatus
const (
	NONE_RelationshipStatus          RelationshipStatus = "None"
	USER_RelationshipStatus          RelationshipStatus = "User"
	FRIEND_RelationshipStatus        RelationshipStatus = "Friend"
	OUTGOING_RelationshipStatus      RelationshipStatus = "Outgoing"
	INCOMING_RelationshipStatus      RelationshipStatus = "Incoming"
	BLOCKED_RelationshipStatus       RelationshipStatus = "Blocked"
	BLOCKED_OTHER_RelationshipStatus RelationshipStatus = "BlockedOther"
)

type MutualResponse struct {
	// Array of mutual user IDs that both users are friends with
	Users []string `json:"users"`
	// Array of mutual server IDs that both users are in
	Servers []string `json:"servers"`
}
