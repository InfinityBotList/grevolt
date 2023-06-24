package permissions

// A simple way to create permissions
type PermissionManager struct {
	perms []Permission
}

// Creates a new PermissionMaker
func NewPermissionManager(perms ...Permission) *PermissionManager {
	return &PermissionManager{
		perms: perms,
	}
}

// Adds a permission to the maker
func (p *PermissionManager) Add(perm Permission) *PermissionManager {
	p.perms = append(p.perms, perm)
	return p
}

// Returns the permissions as a uint64
func (p *PermissionManager) Build() uint64 {
	var perms uint64
	for _, perm := range p.perms {
		perms |= uint64(perm)
	}
	return perms
}
