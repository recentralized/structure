package dst

import (
	"fmt"
	"io"

	"github.com/recentralized/structure/cid"
	"github.com/recentralized/structure/content"
	"github.com/recentralized/structure/uri"
)

// Locator is the interface for defining the locations of data in a
// destination.
type Locator interface {

	// NewHash generates the hash for data.
	NewHash(io.Reader) (cid.ContentID, error)

	// IndexURI returns the document that stores the index.
	IndexURI() uri.URI

	// RefsURI returns the document that should store this ref. This is
	// generally the same as IndexURI, but returning a different value
	// would allow you to shard the refs.
	RefsURI(cid.ContentID) uri.URI

	// DataURI returns the location that this data should be stored.
	DataURI(cid.ContentID, *content.Meta) uri.URI

	// MetaURI returns the location that this meta should be stored.
	MetaURI(cid.ContentID, *content.Meta) uri.URI
}

// NewFilesystemLocator initializes the standard locator for use on filesystems
// and filesystem-like storage media such as AWS S3.
func NewFilesystemLocator() Locator {
	return fsLocator{
		indexFile: "index.json",
		classToCategory: map[content.Class]string{
			content.Image: "media",
		},
		unknownCategory: "unknown",
		zeroDateDir:     "Undated",
	}
}

type fsLocator struct {
	indexFile       string
	classToCategory map[content.Class]string
	unknownCategory string
	zeroDateDir     string
}

func (l fsLocator) NewHash(r io.Reader) (cid.ContentID, error) {
	return cid.New(r)
}

func (l fsLocator) IndexURI() uri.URI {
	return uri.TrustedNew(l.indexFile)
}

func (l fsLocator) RefsURI(cid cid.ContentID) uri.URI {
	return uri.TrustedNew(l.indexFile)
}

// media/2006/2006-01-02/<cid>.<ext>
// media/Undated/hash(<cid>)/<cid>.<ext>
// <category>/hash(<cid>)/<cid>.<ext>
func (l fsLocator) DataURI(cid cid.ContentID, meta *content.Meta) uri.URI {
	var (
		key string
		ext = meta.ContentType.Ext()
		cls = meta.ContentType.Class()
	)

	// Categorize by the class of content.
	category := l.classToCategory[cls]
	if category == "" {
		category = l.unknownCategory
	}

	// Customize the path location for each category.
	switch category {

	// "media" category names files by cid and organized by date.
	case "media":
		t := meta.DateCreated()
		if t.IsZero() {
			key = fmt.Sprintf("%s/%s/%s.%s", category, l.zeroDateDir, l.dirs(cid), ext)
			return uri.TrustedNew(key)
		}
		year := t.Format("2006")
		date := t.Format("2006-01-02")
		key = fmt.Sprintf("%s/%s/%s/%s.%s", category, year, date, cid.String(), ext)
		return uri.TrustedNew(key)

	// Otherwise organize by breaking down the cid.
	default:
		if ext == "" {
			key = fmt.Sprintf("%s/%s", category, l.dirs(cid))
		} else {
			key = fmt.Sprintf("%s/%s.%s", category, l.dirs(cid), ext)
		}
		return uri.TrustedNew(key)
	}
}

// meta/hash(<cid>)/<cid>.json
func (l fsLocator) MetaURI(cid cid.ContentID, meta *content.Meta) uri.URI {
	key := fmt.Sprintf("meta/%s.%s", l.dirs(cid), "json")
	return uri.TrustedNew(key)
}

func (l fsLocator) dirs(cid cid.ContentID) string {
	s := cid.String()
	if len(s) > 4 {
		return fmt.Sprintf("%s/%s/%s", s[0:2], s[2:4], s[4:])
	}
	return s
}
