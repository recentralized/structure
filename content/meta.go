package content

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

const (
	// metaVersionV0 is the original implementation.
	metaVersionV0 string = ""

	// metaVersionV1 is the first `structure` implementation.
	metaVersionV1 string = "v1"
)

// MetaVersion is the current version of the Meta document. A new version will
// be introduced for backward-incompatible changes.
const MetaVersion = metaVersionV1

// ErrMetaWrongVersion means that the parsed meta is not at the current
// version, so its data may be incorrectly interpreted.
var ErrMetaWrongVersion = errors.New("index is not at a compatible verison")

// Meta is all of the potential metadata about content.
type Meta struct {
	Version     string
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

// NewMeta initializes a new Meta at the current version.
func NewMeta() *Meta {
	return &Meta{Version: MetaVersion}
}

// ParseMetaJSON loads Meta from JSON. If the loaded data cannot be
// transparently upgraded to the current version then ErrMetaWrongVersion is
// returned.
func ParseMetaJSON(r io.Reader) (*Meta, error) {
	meta := &Meta{}
	err := json.NewDecoder(r).Decode(meta)
	if err != nil {
		return nil, err
	}
	switch meta.Version {
	case metaVersionV1:
	case metaVersionV0:
		meta.Version = metaVersionV1
	default:
		return nil, ErrMetaWrongVersion
	}
	return meta, nil
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
