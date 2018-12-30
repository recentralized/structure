package meta

import "time"

type flickrVisibility string

const (
	// FlickrPublic means that the photo is visible to the public.
	FlickrPublic flickrVisibility = "public"

	// FlickrPrivate means that the photo is not visible to the public,
	// It may also be visible to friends, family, or both.
	FlickrPrivate = "private"
)

// FlickrMedia is the full set of data available for media on Flickr.
// https://www.flickr.com/services/api/
// https://www.flickr.com/services/api/flickr.photos.getInfo.html
type FlickrMedia struct {
	ID            string               `json:"id"`
	UserID        string               `json:"user_id,omitempty"`
	Username      string               `json:"username,omitempty"`
	Title         string               `json:"title,omitempty"`
	Description   string               `json:"description,omitempty"`
	PostedAt      *time.Time           `json:"posted_at,omitempty"`
	TakenAt       *time.Time           `json:"taken_at,omitempty"`
	LastUpdateAt  *time.Time           `json:"last_update_at,omitempty"`
	URL           string               `json:"url,omitempty"`
	Visibility    flickrVisibility     `json:"visibility,omitempty"`
	IsPublic      bool                 `json:"is_public,omitempty"`
	IsFriendsOnly bool                 `json:"is_friends_only,omitempty"`
	IsFamilyOnly  bool                 `json:"is_family_only,omitempty"`
	License       string               `json:"license,omitempty"`
	LicenseURL    string               `json:"license_url,omitempty"`
	Geo           *FlickrMediaGeo      `json:"geo,omitempty"`
	Views         int                  `json:"views,omitempty"`
	Faves         []FlickrMediaFave    `json:"faves,omitempty"`
	Tags          []FlickrMediaTag     `json:"tags,omitempty"`
	People        []FlickrMediaPerson  `json:"people,omitempty"`
	Notes         []FlickrMediaNote    `json:"notes,omitempty"`
	Sets          []FlickrMediaInSet   `json:"sets,omitempty"`
	Pools         []FlickrMediaInPool  `json:"pools,omitempty"`
	Comments      []FlickrMediaComment `json:"comments,omitempty"`
}

// FlickrMediaFave is who favorited an image on Flickr.
// https://www.flickr.com/services/api/flickr.photos.getFavorites.html
type FlickrMediaFave struct {
	UserID   string     `json:"user_id"`
	Username string     `json:"username,omitempty"`
	Date     *time.Time `json:"date,omitempty"`
}

// FlickrGeoContext describes context of the location.
type FlickrGeoContext string

// Values for FlickrMediaGeo.Context
const (
	FlickrGeoContextNone    FlickrGeoContext = ""
	FlickrGeoContextInside                   = "inside"
	FlickrGeoContextOutside                  = "outside"
)

// FlickrMediaGeo is geo data on a Flickr image.
// https://www.flickr.com/services/api/flickr.photos.geo.getLocation.html
// https://www.flickr.com/services/api/flickr.photos.geo.setLocation.html
type FlickrMediaGeo struct {
	// The latitude whose valid range is -90 to 90.
	Latitude float64 `json:"latitude"`

	// The longitude whose valid range is -180 to 180.
	Longitude float64 `json:"longitude"`

	// Recorded accuracy level of the location information. World level is
	// 1, Country is ~3, Region ~6, City ~11, Street ~16. Current range is
	// 1-16.
	Accuracy int `json:"accuracy,omitempty"`

	// Context is a numeric value representing the photo's geotagginess
	// beyond latitude and longitude. For example, you may wish to indicate
	// that a photo was taken "inside" or "outside".
	Context FlickrGeoContext `json:"context,omitempty"`

	// Where On Earth ID
	WoeID string `json:"woe_id,omitempty"`

	// Details about the tagged places.
	Places []FlickrPlace `json:"places,omitempty"`
}

// FlickrPlaceType describes the type of place.
type FlickrPlaceType string

// Values for FlickrMediaPlace.Type
const (
	FlickrPlaceNone         FlickrPlaceType = ""
	FlickrPlaceNeighborhood                 = "neighborhood"
	FlickrPlaceLocality                     = "locality"
	FlickrPlaceCounty                       = "county"
	FlickrPlaceRegion                       = "region"
	FlickrPlaceCountry                      = "country"
	FlickrPlaceContinent                    = "continent"
)

// FlickrPlace contains more information about the location information on
// a Flickr photo.
// https://www.flickr.com/services/api/flickr.places.getInfo.html
type FlickrPlace struct {
	WoeID     string          `json:"woe_id"`
	Name      string          `json:"name,omitempty"`
	Type      FlickrPlaceType `json:"type,omitempty"`
	Latitude  float64         `json:"latitude"`
	Longitude float64         `json:"longitude"`
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
	AddedByUserID string `json:"added_by_user_id,omitempty"`
}

// FlickrMediaNote is a note on a Flickr image.
// https://www.flickr.com/services/api/flickr.photos.getInfo.html
type FlickrMediaNote struct {
	ID       string           `json:"id"`
	UserID   string           `json:"user_id,omitempty"`
	Username string           `json:"username,omitempty"`
	Text     string           `json:"text"`
	Coords   NormalizedCoords `json:"coords"`
}

// NormalizedCoords are normalized coordinates - 0-1. Flickr stores them as pixels
// assuming 500px wide image.
type NormalizedCoords struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}

// FlickrMediaInSet is a set a Flickr image is part of.
// https://www.flickr.com/services/api/flickr.photos.getAllContexts.html
// position: https://www.flickr.com/services/api/flickr.photosets.getPhotos.html
type FlickrMediaInSet struct {
	ID       string `json:"id"`
	Name     string `json:"name,omitempty"`
	Position int    `json:"position,omitempty"`
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
	UserID   string     `json:"user_id"`
	Username string     `json:"username,omitempty"`
	Text     string     `json:"text"`
	Date     *time.Time `json:"date,omitempty"`
	URL      string     `json:"url,omitempty"`
}
