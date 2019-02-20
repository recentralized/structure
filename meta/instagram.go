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
	URL string `json:"url,omitempty"`

	// User that posted it.
	UserID       string `json:"user_id,omitempty"`
	Username     string `json:"username,omitempty"`
	UserFullName string `json:"user_full_name,omitempty"`

	// Basic information when posted.
	Caption   string     `json:"caption,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Additional information when posted.
	Filter      string                `json:"filter,omitempty"`
	Tags        []string              `json:"tags,omitempty"`
	TaggedUsers []InstagramTaggedUser `json:"tagged_users,omitempty"`
	Location    *InstagramLocation    `json:"location,omitempty"`

	// User activity.
	LikesCount    int                `json:"likes_count,omitempty"`
	CommentsCount int                `json:"comments_count,omitempty"`
	Comments      []InstagramComment `json:"comments,omitempty"`
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
	ID       string     `json:"id"`
	Text     string     `json:"text"`
	Username string     `json:"username,omitempty"`
	Date     *time.Time `json:"date,omitempty"`
}

//InstagramTaggedUser is a user tagged in media.
type InstagramTaggedUser struct {
	Username string  `json:"username"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}
