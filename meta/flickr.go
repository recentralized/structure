package meta

import "time"

// FlickrMedia is the full set of data available for media on Flickr.
type FlickrMedia struct {
	ID          string               `json:"id"`
	Title       string               `json:"title,omitempty"`
	Description string               `json:"description,omitempty"`
	Privacy     string               `json:"privacy,omitempty"`
	License     string               `json:"license,omitempty"`
	URL         string               `json:"url,omitempty"`
	Geo         *FlickrMediaGeo      `json:"geo,omitempty"`
	Views       int                  `json:"views,omitempty"`
	Faves       []FlickrMediaFave    `json:"faves,omitempty"`
	Tags        []FlickrMediaTag     `json:"tags,omitempty"`
	People      []FlickrMediaPerson  `json:"people,omitempty"`
	Notes       []FlickrMediaNote    `json:"note,omitempty"`
	Sets        []FlickrMediaInSet   `json:"sets,omitempty"`
	Pools       []FlickrMediaInPool  `json:"pools,omitempty"`
	Comments    []FlickrMediaComment `json:"comments,omitempty"`
}

// FlickrMediaFave is who favorited an image on Flickr.
// https://www.flickr.com/services/api/flickr.photos.getFavorites.html
type FlickrMediaFave struct {
	NSID     string    `json:"nsid,omitempty"`
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
	ID       string `json:"id,omitempty"`
	NSID     string `json:"nsid,omitempty"`
	Username string `json:"username,omitempty"`
	Tag      string `json:"tag,omitempty"`
	RawTag   string `json:"raw_tag,omitempty"`
}

// FlickrMediaPerson is a person tagged in a Flickr image.
// https://www.flickr.com/services/api/flickr.photos.people.getList.html
type FlickrMediaPerson struct {
	NSID        string `json:"nsid,omitempty"`
	Username    string `json:"username,omitempty"`
	X           int    `json:"x,omitempty"`
	Y           int    `json:"y,omitempty"`
	W           int    `json:"w,omitempty"`
	H           int    `json:"h,omitempty"`
	AddedByNSID string `json:"added_by_nsid,omitempty"`
}

// FlickrMediaNote is a note on a Flickr image.
// https://www.flickr.com/services/api/flickr.photos.getInfo.html
type FlickrMediaNote struct {
	ID       string `json:"id,omitempty"`
	NSID     string `json:"nsid,omitempty"`
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
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Order int    `json:"order,omitempty"`
}

// FlickrMediaInPool is a pool that a Flickr image is part of.
// https://www.flickr.com/services/api/flickr.photos.getAllContexts.html
type FlickrMediaInPool struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// FlickrMediaComment is a comment on a Flickr image.
// https://www.flickr.com/services/api/flickr.photos.comments.getList.html
type FlickrMediaComment struct {
	ID       string    `json:"id,omitempty"`
	NSID     string    `json:"nsid,omitempty"`
	Username string    `json:"username,omitempty"`
	Text     string    `json:"text,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	URL      string    `json:"url,omitempty"`
}
