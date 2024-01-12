package types

// FieldsServer : Optional fields on server object
type FieldsServer string

// List of FieldsServer
const (
	DESCRIPTION_FieldsServer     FieldsServer = "Description"
	CATEGORIES_FieldsServer      FieldsServer = "Categories"
	SYSTEM_MESSAGES_FieldsServer FieldsServer = "SystemMessages"
	ICON_FieldsServer            FieldsServer = "Icon"
	BANNER_FieldsServer          FieldsServer = "Banner"
)

// Representation of a server on Revolt
type Server struct {
	// Unique Id
	Id string `json:"_id"`
	// User id of the owner
	Owner string `json:"owner"`
	// Name of the server
	Name string `json:"name"`
	// Description for the server
	Description string `json:"description,omitempty"`
	// Channels within this server
	Channels []string `json:"channels"`
	// Categories for this server
	Categories []*Category `json:"categories,omitempty"`
	// Configuration for sending system event messages
	SystemMessages *SystemMessageChannels `json:"system_messages,omitempty"`
	// Roles for this server
	Roles map[string]*Role `json:"roles,omitempty"`
	// Default set of server and channel permissions
	DefaultPermissions uint64 `json:"default_permissions"`
	// Icon attachment
	Icon *File `json:"icon,omitempty"`
	// Banner attachment
	Banner *File `json:"banner,omitempty"`
	// Bitfield of server flags
	Flags uint64 `json:"flags,omitempty"`
	// Whether this server is flagged as not safe for work
	Nsfw bool `json:"nsfw,omitempty"`
	// Whether to enable analytics
	Analytics bool `json:"analytics,omitempty"`
	// Whether this server should be publicly discoverable
	Discoverable bool `json:"discoverable,omitempty"`
}

// Data to create a server
type DataCreateServer struct {
	// Server name
	Name string `json:"name"`
	// Server description
	Description string `json:"description,omitempty"`
	// Whether this server is age-restricted
	Nsfw bool `json:"nsfw,omitempty"`
}

// Response from creating a server
type CreateServerResponse struct {
	// Server object
	Server *Server `json:"server"`
	// Default channels
	Channels []*Channel `json:"channels"`
}

// Data to edit a server
type DataEditServer struct {
	// Server name
	Name string `json:"name,omitempty"`
	// Server description
	Description string `json:"description,omitempty"`
	// Attachment Id for icon
	Icon string `json:"icon,omitempty"`
	// Attachment Id for banner
	Banner string `json:"banner,omitempty"`
	// Category structure for server
	Categories []*Category `json:"categories,omitempty"`
	// System message configuration
	SystemMessages *SystemMessageChannels `json:"system_messages,omitempty"`
	// Bitfield of server flags
	Flags uint64 `json:"flags,omitempty"`
	// Whether this server is public and should show up on [Revolt Discover](https://rvlt.gg)
	Discoverable bool `json:"discoverable,omitempty"`
	// Whether analytics should be collected for this server  Must be enabled in order to show up on [Revolt Discover](https://rvlt.gg).
	Analytics bool `json:"analytics,omitempty"`
	// Fields to remove from server object
	Remove []*FieldsServer `json:"remove,omitempty"`
}
