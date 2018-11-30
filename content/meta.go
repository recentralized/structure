package content

import (
	"time"
)

// Meta is all of the potential metadata about content.
type Meta struct {
	ContentType Type
	Size        int64

	// Metadata that came from the content itself.
	Inherent MetaContent

	// Metadata that came from nearby, such as an XMP sidecar file or other
	// source of metadata.
	Sidecar MetaContent

	// Metadata that came from the source of the data.
	Srcs SrcSpecific
}

// DateCreated returns the time that the content was created. It chooses the
// most most likely to be correct time from available sources. If there is no
// time available it returns time.Time's zero value.
func (m *Meta) DateCreated() time.Time {
	times := []time.Time{
		m.Sidecar.Created,
		m.Inherent.Created,
	}
	for _, t := range times {
		if !t.IsZero() {
			return t
		}
	}
	var t time.Time
	return t
}

// Image returns the inherent image data.
func (m *Meta) Image() MetaImage {
	return m.Inherent.Image
}

// MetaContent contains all data that describes the content directly.
type MetaContent struct {
	Created time.Time
	Image   MetaImage
	Exif    Exif
}

// MetaImage contains standard fields for all images.
type MetaImage struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// SrcSpecific contains source-specific metadata.
type SrcSpecific struct {
	//Flickr *flickr.FlickrActivity `json:"flickr,omitempty"`
}

func (m MetaContent) isZero() bool {
	return m.Created.IsZero() &&
		m.Image.isZero() &&
		len(m.Exif) == 0
}

func (m MetaImage) isZero() bool {
	return m.Width == 0 && m.Height == 0
}
