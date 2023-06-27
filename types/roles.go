package types

// FieldsRole : Optional fields on server object
type FieldsRole string

// List of FieldsRole
const (
	COLOUR_FieldsRole FieldsRole = "Colour"
)

// Representation of a server role
type Role struct {
	// Role name
	Name string `json:"name"`
	// Permissions available to this role
	Permissions *PermissionOverrideField `json:"permissions"`
	// Colour used for this role  This can be any valid CSS colour
	Colour string `json:"colour,omitempty"`
	// Whether this role should be shown separately on the member sidebar
	Hoist bool `json:"hoist,omitempty"`
	// Ranking of this role
	Rank uint64 `json:"rank,omitempty"`
}

// Data needed to create a role
type DataCreateRole struct {
	// Role name
	Name string `json:"name"`
	// Ranking position  Smaller values take priority.
	Rank uint64 `json:"rank,omitempty"`
}

// Data needed to edit a role
type DataEditRole struct {
	// Role name
	Name string `json:"name,omitempty"`
	// Role colour
	Colour string `json:"colour,omitempty"`
	// Whether this role should be displayed separately
	Hoist bool `json:"hoist,omitempty"`
	// Ranking position  Smaller values take priority.
	Rank uint64 `json:"rank,omitempty"`
	// Fields to remove from role object
	Remove []FieldsRole `json:"remove,omitempty"`
}

// Response upon creating a role
type NewRoleResponse struct {
	// Id of the role
	Id string `json:"id"`
	// New role
	Role *Role `json:"role"`
}
