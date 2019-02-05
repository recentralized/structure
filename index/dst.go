package index

import (
	"fmt"
	"strconv"
	"time"

	"github.com/recentralized/structure/data"
	"github.com/recentralized/structure/uri"
	"github.com/satori/go.uuid"
)

// DstID is the unique ID for each user-defined storage destination.
type DstID string

// Dst is a distinct location that content has been stored.
type Dst struct {

	// DstID is a unique identifier for this set of storage URIs.
	DstID DstID

	// IndexURI is a unique identifier for the storage location of this
	// destination's ref index.
	IndexURI uri.URI

	// DataURI is a unique identifier for the storage location of this
	// destination's data. If a DstItem's DataURI is relative, this URI can
	// be used to resolve it.
	DataURI uri.URI

	// MetaURI is a unique identifier for the storage location of this
	// destination's metadata. If a DstItem's MetaURI is relative, this URI
	// can be used to resolve it.
	MetaURI uri.URI
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

// NewDstAllAt initializes a Dst with its index, data, and meta all at the
// given URI.
func NewDstAllAt(baseURI uri.URI) Dst {
	return NewDst(baseURI, baseURI, baseURI)
}

func (d Dst) String() string {
	return fmt.Sprintf("<Dst %s index:%s data:%s meta:%s>", d.DstID, strconv.Quote(d.IndexURI.String()), strconv.Quote(d.DataURI.String()), strconv.Quote(d.MetaURI.String()))
}

func (id DstID) String() string {
	return string(id)
}

// DstItem is the storage location of an item in a destination. This record is
// immutable in the index.
type DstItem struct {

	// DstID is the Dst that this item was stored in.
	DstID DstID

	// DataURI is a unique identifier for the data of this item. It is
	// typically a URL pointing to the storage location of the raw data.
	// The URI is always relative, resolved to absolute using Dst.DataURI.
	DataURI uri.URI

	// MetaURI is a unique identifier for the metadata of this item. It is
	// typically a URL pointing to the storage location of the metadata.
	// The URI is always relative resolved to absolute using Dst.MetaURI.
	MetaURI uri.URI

	// DataType is the type of data that's stored. This field is useful to
	// filter on a type of data without accessing the meta. It is normally
	// the type of the original data but could be otherwise if the data has
	// been compressed on storage, for example.
	DataType data.Stored

	// DataSize is the size of the stored data in bytes. This field is
	// useful to calculate things like storage and transfer costs. It will
	// normally equal the size of the content, but may differ if the
	// content is compressed on storage, for example. This field should be
	// considered immutable; not changed after the original value.
	DataSize int64

	// Metaize is the size of the stored data in bytes. This field is
	// useful to calculate things like storage and transfer costs. It will
	// normally equal the size of the content, but may differ if the
	// content is compressed on storage, for example. This field should be
	// updated each time the metadata changes.
	MetaSize int64

	// StoredAt is the time that the item was originally stored. This field
	// should be considered immutable; not changed after the original value.
	StoredAt time.Time

	// UpdatedAt is the time that the item was updated. This typically
	// means metadata updates since data is immutable.
	UpdatedAt time.Time
}

// EqualKey determines if two DstItem have the same primary key.
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

// Equal determines if two DstItem are completely identical.
func (d DstItem) Equal(dd DstItem) bool {
	switch {
	case d.DstID != dd.DstID:
	case !d.DataURI.Equal(dd.DataURI):
	case !d.MetaURI.Equal(dd.MetaURI):
	case d.DataSize != dd.DataSize:
	case d.MetaSize != dd.MetaSize:
	case !d.StoredAt.Equal(dd.StoredAt):
	case !d.UpdatedAt.Equal(dd.UpdatedAt):
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
