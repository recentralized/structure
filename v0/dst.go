package structure

import (
	"fmt"
	"strconv"
	"time"

	"github.com/recentralized/structure/uri"
	"github.com/satori/go.uuid"
)

// DstID is the unique ID for each user-defined storage destination.
type DstID string

// Dst is a distinct location that content has been stored.
type Dst struct {

	// DstID is a unique identifier for this set of storage URIs.
	DstID DstID `json:"dst_id"`

	// IndexURI is a unique identifier for the storage location of this
	// destination's ref index.
	IndexURI uri.URI `json:"index_uri"`

	// DataURI is a unique identifier for the storage location of this
	// destination's data. If a DstItem's DataURI is relative, this URI can
	// be used to resolve it.
	DataURI uri.URI `json:"data_uri"`

	// MetaURI is a unique identifier for the storage location of this
	// destination's metadata. If a DstItem's MetaURI is relative, this URI
	// can be used to resolve it.
	MetaURI uri.URI `json:"meta_uri"`
}

// NewDst initializes a storage destination. All destinations initialized with
// the equivalent URIs are equivalent.
//
// Examples
//
//	NewDst(uri.TrustedNew("s3://bucket/"),
//	       uri.TrustedNew("s3://bucket/data"),
//	       uri.TrustedNew("s3://bucket/meta"))
//
//	NewDst(uri.TrustedNew("file:///Users/rcarver/Pictures/"),
//	       uri.TrustedNew("s3://bucket/data/"),
//	       uri.TrustedNew("s3://bucket/meta/"))
//
func NewDst(indexURI, dataURI, metaURI uri.URI) Dst {
	return Dst{
		DstID:    newDstIDFromURIs(indexURI, dataURI, metaURI),
		IndexURI: indexURI,
		DataURI:  dataURI,
		MetaURI:  metaURI,
	}
}

func (d Dst) String() string {
	return fmt.Sprintf("<Dst %s index:%s data:%s meta:%s>", d.DstID, strconv.Quote(d.IndexURI.String()), strconv.Quote(d.DataURI.String()), strconv.Quote(d.MetaURI.String()))
}

func (id DstID) String() string {
	return string(id)
}

// DstItem is the storage location of an item in a destination.
type DstItem struct {
	DstID DstID `json:"dst_id"`

	// DataURI is a unique identifier for the data of this item. It is
	// typically a URL pointing to the storage location of the raw data.
	// The URI may be relative or absolute. If relative, it should resolve
	// to absolute using Dst.DataURI.
	DataURI uri.URI `json:"data_uri"`

	// MetaURI is a unique identifier for the metadata of this item. It is
	// typically a URL pointing to the storage location of the metadata.
	// The URI may be relative or absolute. If relative, it should resolve
	// to absolute using Dst.MetaURI.
	//
	// Metadata can be versioned. For some destinations, this URI may point
	// to a metadata index from which any version can be retrieved. For
	// others, it may point to the current metadata. See
	// content.MetaReader.
	MetaURI uri.URI `json:"meta_uri"`

	// StoredAt is the time that the item was stored.
	StoredAt time.Time `json:"stored_at"`
}

// EqualKey determins if two DstItem have the same primary key.
func (d DstItem) EqualKey(dd DstItem) bool {
	switch {
	case d.DstID != dd.DstID:
	case !d.DataURI.Equal(dd.DataURI):
	case !d.MetaURI.Equal(dd.MetaURI):
	default:
		return true
	}
	return false
}

var uuidNamespaceDst = uuid.NewV5(uuid.Nil, "Storage Destination ID")

// newDstIDFromURIs creates a DstID by generating a UUID from the URIs.  This
// ID will be consistent given the same URIs.
func newDstIDFromURIs(indexURI, dataURI, metaURI uri.URI) DstID {
	// Generate a UUID in the URL namespace for the index, data, and meta URIs.
	a := uuid.NewV5(uuid.NamespaceURL, indexURI.String())
	b := uuid.NewV5(uuid.NamespaceURL, dataURI.String())
	c := uuid.NewV5(uuid.NamespaceURL, metaURI.String())
	// Generate a UUID in our custom namespace as the concatenation of the url UUIDs.
	r := uuid.NewV5(uuidNamespaceDst, fmt.Sprintf("%s|%s|%s", a, b, c))
	return DstID(r.String())
}
