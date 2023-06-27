package types

import "time"

// Type of embed
type EmbedType string

const (
	WEBSITE_EmbedType = "Website"
	IMAGE_EmbedType   = "Image"
	VIDEO_EmbedType   = "Video"
	TEXT_EmbedType    = "Text"
)

// Type of special embed content (remote content)
type MessageEmbedSpecialType string

const (
	NONE_MessageEmbedSpecialType       = "None"
	GIF_MessageEmbedSpecialType        = "GIF"
	YOUTUBE_MessageEmbedSpecialType    = "YouTube"
	LIGHTSPEED_MessageEmbedSpecialType = "Lightspeed"
	TWITCH_MessageEmbedSpecialType     = "Twitch"
	SPOTIFY_MessageEmbedSpecialType    = "Spotify"
	SOUNDCLOUD_MessageEmbedSpecialType = "Soundcloud"
	BANDCAMP_MessageEmbedSpecialType   = "Bandcamp"
	STREAMABLE_MessageEmbedSpecialType = "Streamable"
)

// Representation of a text embed before it is sent.
type SendableEmbed struct {
	IconUrl     string `json:"icon_url,omitempty"`
	Url         string `json:"url,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Media       string `json:"media,omitempty"`
	Colour      string `json:"colour,omitempty"`
}

// Representation of an message embed on Revolt.
type MessageEmbed struct {
	// Type of embed
	Type EmbedType `json:"type,omitempty"`
	// Only on Text embeds, the embeds icon url
	IconUrl string `json:"icon_url,omitempty"`
	// Available on most embed types, the url of the embed
	Url string `json:"url,omitempty"`
	// Only sent on Website embeds, the original url of the embed
	OriginalUrl string `json:"original_url,omitempty"`
	// Any special remote content, only sent on Website embeds
	Special *MessageEmbedSpecial `json:"special,omitempty"`
	// Sent on all event types other than Video/Image embeds (?)
	Title string `json:"title,omitempty"`
	// Sent on all event types other than Video/Image embeds (?)
	Description string `json:"description,omitempty"`
	// Image of the website. Only sent on Website embeds
	Image *MessageEmbedImage `json:"image,omitempty"`
	// Video embed of the website. Only sent on Website embeds
	Video *MessageEmbedVideo `json:"video,omitempty"`
	// Site name, only sent on Website embeds
	SiteName string `json:"site_name,omitempty"`
	// ID of uploaded autumn file
	//
	// <does not appear to be sent at this time, may be sent in the future>
	Media *File `json:"media,omitempty"`
	// Colour of the embed (CSS colour)
	Colour string `json:"colour,omitempty"`
}

type MessageEmbedSpecial struct {
	Type MessageEmbedSpecialType `json:"type,omitempty"`

	// The ID of the content on the remote service
	ID string `json:"id,omitempty"`

	// The title of the content
	Timestamp time.Time `json:"timestamp,omitempty"`

	// Identifies the type of content for types: Lightspeed, Twitch, Spotify, and Bandcamp
	ContentType string `json:"content_type,omitempty"`
}

type MessageEmbedImage struct {
	// Size of the image
	Size string `json:"size"`

	// URL of the image
	URL string `json:"url"`

	// Width of the image
	Width int `json:"width"`

	// Height of the image
	Height int `json:"height"`
}

type MessageEmbedVideo struct {
	// URL to the original video
	Url string `json:"url"`

	// Width of the video
	Width int `json:"width"`

	// Height of the video
	Height int `json:"height"`
}

// TwitchType : Type of remote Twitch content
//
// <currently unused?>
type TwitchType string

// List of TwitchType
const (
	CHANNEL_TwitchType TwitchType = "Channel"
	VIDEO_TwitchType   TwitchType = "Video"
	CLIP_TwitchType    TwitchType = "Clip"
)

// LightspeedType : Type of remote Lightspeed.tv content
//
// <currently unused?>
type LightspeedType string

// List of LightspeedType
const (
	CHANNEL_LightspeedType LightspeedType = "Channel"
)

// BandcampType : Type of remote Bandcamp content
//
// <currently unused?>
type BandcampType string

// List of BandcampType
const (
	ALBUM_BandcampType BandcampType = "Album"
	TRACK_BandcampType BandcampType = "Track"
)

// ImageSize : Image positioning and size
//
// <currently unused?>
type ImageSize string

// List of ImageSize
const (
	LARGE_ImageSize   ImageSize = "Large"
	PREVIEW_ImageSize ImageSize = "Preview"
)
