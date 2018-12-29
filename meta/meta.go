package meta

import (
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/recentralized/structure/data"
)

const (
	// versionV0 is the original implementation.
	versionV0 string = ""

	// versionV1 is the first `structure` implementation.
	versionV1 string = "v1"
)

// Version is the current version of the Meta document. A new version will be
// introduced for backward-incompatible changes.
const Version = versionV1

// ErrWrongVersion means that the parsed meta is not at the current
// version, so its data may be incorrectly interpreted.
var ErrWrongVersion = errors.New("meta is not at a compatible version")

// Meta is all of the potential metadata about content.
type Meta struct {
	Version string
	Type    data.Type
	Size    int64

	// Metadata that came from the content itself.
	Inherent Content

	// Metadata that came from nearby, such as an XMP sidecar file or other
	// source of metadata.
	V0Sidecar Content

	// Metadata that came from the source of the data.
	V0Srcs SrcSpecific
}

// New initializes a new Meta at the current version.
func New() *Meta {
	return &Meta{Version: Version}
}

// ParseJSON loads Meta from JSON. If the loaded data cannot be transparently
// upgraded to the current version then ErrMetaWrongVersion is returned.
func ParseJSON(r io.Reader) (*Meta, error) {
	meta := &Meta{}
	err := json.NewDecoder(r).Decode(meta)
	if err != nil {
		return nil, err
	}
	switch meta.Version {
	case versionV1:
	case versionV0:
		meta.Version = versionV1
	default:
		return nil, ErrWrongVersion
	}
	return meta, nil
}

// DateCreated returns the time that the content was created. It chooses the
// most most likely to be correct time from available sources. If there is no
// time available it returns time.Time's zero value.
func (m *Meta) DateCreated() time.Time {
	times := []time.Time{
		m.V0Sidecar.Created,
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
func (m *Meta) Image() Image {
	return m.Inherent.Image
}

// Content contains all data that describes the content directly.
type Content struct {
	Created time.Time
	Image   Image
	Exif    Exif
}

// Image contains standard fields for all images.
type Image struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// SrcSpecific contains source-specific metadata.
type SrcSpecific struct {
	//Flickr *flickr.FlickrActivity `json:"flickr,omitempty"`
}

func (m Content) isZero() bool {
	return m.Created.IsZero() &&
		m.Image.isZero() &&
		len(m.Exif) == 0
}

func (m Image) isZero() bool {
	return m.Width == 0 && m.Height == 0
}
