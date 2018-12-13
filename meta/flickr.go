package meta

import "time"

// FlickrMedia is the full set of data available for media on Flickr.
// https://www.flickr.com/services/api/
// https://www.flickr.com/services/api/flickr.photos.getInfo.html
type FlickrMedia struct {
	ID           string               `json:"id"`
	UserID       string               `json:"user_id,omitempty"`
	Username     string               `json:"username,omitempty"`
	Title        string               `json:"title,omitempty"`
	Description  string               `json:"description,omitempty"`
	PostedAt     *time.Time           `json:"posted_at,omitempty"`
	TakenAt      *time.Time           `json:"taken_at,omitempty"`
	LastUpdateAt *time.Time           `json:"last_update_at,omitempty"`
	URL          string               `json:"url,omitempty"`
	Visibility   string               `json:"visibility,omitempty"`
	License      string               `json:"license,omitempty"`
	Geo          *FlickrMediaGeo      `json:"geo,omitempty"`
	Views        int                  `json:"views,omitempty"`
	Faves        []FlickrMediaFave    `json:"faves,omitempty"`
	Tags         []FlickrMediaTag     `json:"tags,omitempty"`
	People       []FlickrMediaPerson  `json:"people,omitempty"`
	Notes        []FlickrMediaNote    `json:"notes,omitempty"`
	Sets         []FlickrMediaInSet   `json:"sets,omitempty"`
	Pools        []FlickrMediaInPool  `json:"pools,omitempty"`
	Comments     []FlickrMediaComment `json:"comments,omitempty"`
}

// FlickrMediaFave is who favorited an image on Flickr.
// https://www.flickr.com/services/api/flickr.photos.getFavorites.html
type FlickrMediaFave struct {
	UserID   string    `json:"user_id"`
	Username string    `json:"username,omitempty"`
	Date     time.Time `json:"date,omitempty"`
}

// FlickrMediaGeo is geo data on a Flickr image.
// https://www.flickr.com/services/api/flickr.photos.getExif.html
type FlickrMediaGeo struct {
	Latitude  int `json:"latitide,omitempty"`
	Longitude int `json:"longitude,omitempty"`
	Accuracy  int `json:"accuracy,omitempty"`
}

// FlickrMediaTag is a tag applied to an image on Flickr.
// https://www.flickr.com/services/api/flickr.tags.getListPhoto.html
// https://www.flickr.com/services/api/flickr.photos.getInfo.html
type FlickrMediaTag struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Tag      string `json:"tag,omitempty"`
	RawTag   string `json:"raw_tag,omitempty"`
}

// FlickrMediaPerson is a person tagged in a Flickr image.
// https://www.flickr.com/services/api/flickr.photos.people.getList.html
type FlickrMediaPerson struct {
	UserID        string `json:"user_id"`
	Username      string `json:"username,omitempty"`
	X             int    `json:"x,omitempty"`
	Y             int    `json:"y,omitempty"`
	W             int    `json:"w,omitempty"`
	H             int    `json:"h,omitempty"`
	AddedByUserID string `json:"added_by_user_id,omitempty"`
}

// FlickrMediaNote is a note on a Flickr image.
// https://www.flickr.com/services/api/flickr.photos.getInfo.html
type FlickrMediaNote struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	X        int    `json:"x,omitempty"`
	Y        int    `json:"y,omitempty"`
	W        int    `json:"w,omitempty"`
	H        int    `json:"h,omitempty"`
}

// FlickrMediaInSet is a set a Flickr image is part of.
// https://www.flickr.com/services/api/flickr.photos.getAllContexts.html
// order: https://www.flickr.com/services/api/flickr.photosets.getPhotos.html
type FlickrMediaInSet struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Order int    `json:"order,omitempty"`
}

// FlickrMediaInPool is a pool that a Flickr image is part of.
// https://www.flickr.com/services/api/flickr.photos.getAllContexts.html
type FlickrMediaInPool struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// FlickrMediaComment is a comment on a Flickr image.
// https://www.flickr.com/services/api/flickr.photos.comments.getList.html
type FlickrMediaComment struct {
	ID       string     `json:"id"`
	UserID   string     `json:"user_id,omitempty"`
	Username string     `json:"username,omitempty"`
	Text     string     `json:"text,omitempty"`
	PostedAt *time.Time `json:"posted_at,omitempty"`
	URL      string     `json:"url,omitempty"`
}
