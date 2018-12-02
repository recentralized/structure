package dst

import (
	"testing"
	"time"

	"github.com/recentralized/structure/cid"
	"github.com/recentralized/structure/data"
	"github.com/recentralized/structure/meta"
)

func TestFilesystemLocator(t *testing.T) {
	tests := []struct {
		desc        string
		hash        cid.ContentID
		meta        *meta.Meta
		wantDataURI string
		wantMetaURI string
	}{
		{
			desc: "dated media",
			hash: cid.NewLiteral("abcdefg"),
			meta: &meta.Meta{
				Type: data.JPG,
				Inherent: meta.Content{
					Created: time.Date(2015, 1, 2, 9, 9, 9, 9, time.UTC),
				},
			},
			wantDataURI: "media/2015/2015-01-02/abcdefg.jpg",
			wantMetaURI: "meta/ab/cd/efg.json",
		},
		{
			desc: "undated media",
			hash: cid.NewLiteral("abcdefg"),
			meta: &meta.Meta{
				Type: data.JPG,
			},
			wantDataURI: "media/Undated/ab/cd/efg.jpg",
			wantMetaURI: "meta/ab/cd/efg.json",
		},
		{
			desc: "unknown class",
			hash: cid.NewLiteral("abcdefg"),
			meta: &meta.Meta{
				Type: data.UnknownType,
			},
			wantDataURI: "unknown/ab/cd/efg",
			wantMetaURI: "meta/ab/cd/efg.json",
		},
	}
	for _, tt := range tests {
		locator := NewFilesystemLocator()
		got := locator.DataURI(tt.hash, tt.meta)
		if got, want := got.String(), tt.wantDataURI; got != want {
			t.Errorf("%q DataURI()\ngot  %s\nwant %s", tt.desc, got, want)
		}
		got = locator.MetaURI(tt.hash, tt.meta)
		if got, want := got.String(), tt.wantMetaURI; got != want {
			t.Errorf("%q MetaURI()\ngot  %s\nwant %s", tt.desc, got, want)
		}
	}
}
