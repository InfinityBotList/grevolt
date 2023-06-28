package types

// Common types on revolt

// Rate limit struct
type RateLimit struct {
	RetryAfter int64 `json:"retry_after"`
}

// Representation of a File on Revolt Generated by Autumn
type File struct {
	// Unique Id
	Id string `json:"_id"`
	// Tag / bucket this file was uploaded to
	Tag string `json:"tag"`
	// Original filename
	Filename string `json:"filename"`
	// Parsed metadata of this file
	Metadata *FileMetadata `json:"metadata"`
	// Raw content type of this file
	ContentType string `json:"content_type"`
	// Size of this file (in bytes)
	Size uint64 `json:"size"`
	// Whether this file was deleted
	Deleted bool `json:"deleted,omitempty"`
	// Whether this file was reported
	Reported bool `json:"reported,omitempty"`
	// Message Id this file is associated with
	MessageId string `json:"message_id,omitempty"`
	// User Id this file is associated with
	UserId string `json:"user_id,omitempty"`
	// Server Id this file is associated with
	ServerId string `json:"server_id,omitempty"`
	// Id of the object this file is associated with
	ObjectId string `json:"object_id,omitempty"`
}

type FileMetadata struct {
	// Type of this file
	Type string `json:"type"`

	// Width of the video
	Width int `json:"width"`

	// Height of the video
	Height int `json:"height"`
}

// API error, not properly generated
type APIError map[string]any

func (e APIError) Type() string {
	typ, ok := e["type"]

	if !ok {
		return "Unknown"
	}

	// Check if typ is a string
	if _, ok := typ.(string); !ok {
		return "Unknown"
	}

	return typ.(string)
}