package index

import (
	"fmt"
	"strconv"
	"time"

	"github.com/recentralized/structure/uri"
	"github.com/satori/go.uuid"
)

// SrcID is the unique ID for each user-defined source.
type SrcID string

// Src is a distinct location that content has been found.
type Src struct {

	// SrcID is a unique identifier for the source URI.
	SrcID SrcID

	// SrcURI is a unique identifier for the source location. It is
	// typically a URL describing the type and location to be searched for
	// content.
	SrcURI uri.URI
}

// SrcItem describes where content was originally found.
type SrcItem struct {

	// SrcID is the Src that this item was found in.
	SrcID SrcID

	// DataURI is a unique identifier for the data of this item. The URI is
	// typically a URL that points to the raw data.
	DataURI uri.URI

	// MetaURI is a unique identifier for the metadata of this item. The
	// URI is typically a URL that points to record containing additional
	// information about the data.
	MetaURI uri.URI

	// ModifiedAt is the last known time that the content was updated. This
	// timestamp should be populated from a representative field in the
	// source. If no "modified" time is available, the "created" time may
	// be used; however if so consumers of this record won't be able to
	// differentiate content changes over time.
	ModifiedAt time.Time
}

// NewSrc initializes a source location. All sources initialized with
// equivalent URIs are equivalent.
//
// Examples
//
//	NewSrc(uri.TrustedNew("file:///Users/rcarver/Pictures/"))
//	NewSrc(uri.TrustedNew("https://instagram.com/rcarver"))
//
func NewSrc(srcURI uri.URI) Src {
	return Src{
		SrcID:  newSrcIDFromURI(srcURI),
		SrcURI: srcURI,
	}
}

func (s Src) String() string {
	return fmt.Sprintf("<Src %s src:%q>", s.SrcID, s.SrcURI.String())
}

func (s SrcItem) String() string {
	return fmt.Sprintf("<srcs.Item %s>", strconv.Quote(s.DataURI.String()))
}

// EqualKey determines if two SrcItem have the same primary key.
func (s SrcItem) EqualKey(ss SrcItem) bool {
	switch {
	case s.SrcID != ss.SrcID:
	case !s.DataURI.Equal(ss.DataURI):
	case !s.MetaURI.Equal(ss.MetaURI):
	default:
		return true
	}
	return false
}

// Equal determines if two SrcItem are completely identical.
func (s SrcItem) Equal(ss SrcItem) bool {
	switch {
	case s.SrcID != ss.SrcID:
	case !s.DataURI.Equal(ss.DataURI):
	case !s.MetaURI.Equal(ss.MetaURI):
	case !s.ModifiedAt.Equal(ss.ModifiedAt):
	default:
		return true
	}
	return false
}

var uuidNamespaceSrc = uuid.NewV5(uuid.Nil, "Storage Source ID")

// newSrcIDFromURI creates a SrcID by generating a UUID from the URI.
// This ID will be consistent given the same URI.
func newSrcIDFromURI(srcURI uri.URI) SrcID {
	// Generate a UUID in the URL namespace for the src URI.
	a := uuid.NewV5(uuid.NamespaceURL, srcURI.String())
	// Generate a UUID in our custom namespace as the concatenation of the url UUIDs.
	r := uuid.NewV5(uuidNamespaceSrc, a.String())
	return SrcID(r.String())
}
