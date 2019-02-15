package meta

import "time"

// InstagramMedia is the full set of data available for media on Instagram.
// https://www.instagram.com/developer/
type InstagramMedia struct {

	// Media ID
	ID string

	// If the media is part of a multi-media post like a carousel, this is
	// the position of this individual piece of media.
	Multiple bool
	Position int

	// Public URL
	Link string

	// Basic information when posted.
	User      InstagramUser
	Location  InstagramLocation
	Caption   string
	CreatedAt time.Time

	// Additional information when posted.
	Filter      string
	Tags        []string
	TaggedUsers []InstagramTaggedUser

	// User activity.
	Likes    int
	Comments []InstagramComment
}

// InstagramUser describes a user on Instagram.
type InstagramUser struct {
	ID       string
	Username string
	FullName string
}

// InstagramLocation describes geo location of an Instagram post.
type InstagramLocation struct {
	ID        string
	Name      string
	Latitude  float64
	Longitude float64
}

// InstagramComment is a comment on an Instagram post.
type InstagramComment struct {
	ID   string
	Text string
	Date time.Time
	From InstagramUser
}

//InstagramTaggedUser is a user tagged in media.
type InstagramTaggedUser struct {
	Username string
	X        float64
	Y        float64
}
