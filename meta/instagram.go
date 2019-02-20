package meta

import "time"

// InstagramPostType defines the type of post.
type InstagramPostType string

// InstagramMediaType defines the type of media in the post.
type InstagramMediaType string

// Types of Instagram posts.
const (
	InstagramImagePost    InstagramPostType = "image"
	InstagramVideoPost                      = "video"
	InstagramCarouselPost                   = "carousel"
)

// Types of Instagram media.
const (
	InstagramImageMedia InstagramMediaType = "image"
	InstagramVideoMedia                    = "video"
)

// InstagramPost is the full set of data available for media on Instagram.
// https://www.instagram.com/developer/
type InstagramPost struct {

	// Media ID
	ID   string            `json:"id"`
	Type InstagramPostType `json:"type,omitempty"`

	// Public URL
	URL string `json:"url,omitempty"`

	// Describe an individual piece of media within a multi-media post.
	Multiple  bool               `json:"multiple,omitempty"`
	MediaType InstagramMediaType `json:"media_type,omitempty"`
	Position  int                `json:"position"`

	// User that posted it.
	UserID       string `json:"user_id,omitempty"`
	Username     string `json:"username,omitempty"`
	UserFullName string `json:"user_full_name,omitempty"`

	// Basic information when posted.
	Caption  string     `json:"caption,omitempty"`
	PostedAt *time.Time `json:"posted_at,omitempty"`

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

// InstagramTaggedUser is a user tagged in media.
type InstagramTaggedUser struct {
	Username string  `json:"username"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}
