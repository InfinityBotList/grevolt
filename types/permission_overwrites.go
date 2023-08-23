package types

// A permission update
type PermissionUpdate struct {
	Permissions uint64 `json:"permissions"`
}

// Representation of a single permission override as it appears on models and in the database
type PermissionOverrideField struct {
	// Allow bit flags
	A uint64 `json:"a"`
	// Disallow bit flags
	D uint64 `json:"d"`
}

// Representation of a single permission override
type PermissionOverride struct {
	// Allow bit flags
	Allow uint64 `json:"allow"`
	// Disallow bit flags
	Deny uint64 `json:"deny"`
}

type PermissionsPatchOverrideField struct {
	Permissions *PermissionOverride `json:"permissions"`
}
