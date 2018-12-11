package meta

// FlickrActivity is the full set of data available for media on Flickr.
type FlickrActivity struct {
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Views       int      `json:"views,omitempty"`
}

// NewFlickrActivity initializes an empty FlickrActivity.
func NewFlickrActivity() *FlickrActivity {
	return &FlickrActivity{
		Tags: make([]string, 0),
	}
}
