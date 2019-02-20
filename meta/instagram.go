package meta

import "time"

// InstagramMedia is the full set of data available for media on Instagram.
// https://www.instagram.com/developer/
type InstagramMedia struct {

	// Media ID
	ID string `json:"id"`

	// If the media is part of a multi-media post like a carousel, this is
	// the position of this individual piece of media.
	Multiple bool `json:"multiple,omitempty"`
	Position int  `json:"position,omitempty"`

	// Public URL
	Link string `json:"link"`

	// Basic information when posted.
	User      *InstagramUser     `json:"user,omitempty"`
	Location  *InstagramLocation `json:"location,omitempty"`
	Caption   string             `json:"caption,omitempty"`
	CreatedAt time.Time          `json:"created_at"`

	// Additional information when posted.
	Filter      string                `json:"filter,omitempty"`
	Tags        []string              `json:"tags,omitempty"`
	TaggedUsers []InstagramTaggedUser `json:"tagged_users,omitempty"`

	// User activity.
	Likes    int                `json:"likes"`
	Comments []InstagramComment `json:"comments,omitempty"`
}

// InstagramUser describes a user on Instagram.
type InstagramUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name,omitempty"`
}

// InstagramLocation describes geo location of an Instagram post.
type InstagramLocation struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// InstagramComment is a comment on an Instagram post.
type InstagramComment struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	Username string    `json:"username"`
	Date     time.Time `json:"date"`
}

//InstagramTaggedUser is a user tagged in media.
type InstagramTaggedUser struct {
	Username string  `json:"username"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}
