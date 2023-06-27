package types

// Data for creating a group
type DataCreateGroup struct {
	// Group name
	Name string `json:"name"`
	// Group description
	Description string `json:"description,omitempty"`
	// Array of user IDs to add to the group.
	//
	// Must be friends with these users.
	Users []string `json:"users"`
	// Whether this group is age-restricted
	Nsfw bool `json:"nsfw,omitempty"`
}
