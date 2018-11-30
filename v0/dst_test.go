package structure

import (
	"reflect"
	"testing"

	"github.com/recentralized/structure/uri"
)

func TestNewDst(t *testing.T) {
	tests := []struct {
		desc     string
		indexURI uri.URI
		dataURI  uri.URI
		metaURI  uri.URI
		want     Dst
	}{
		{
			desc:     "consistent value",
			indexURI: uri.TrustedNew("file:///tmp/"),
			dataURI:  uri.TrustedNew("file:///tmp/data/"),
			metaURI:  uri.TrustedNew("file:///tmp/meta/"),
			want: Dst{
				DstID:    DstID("0b1bccaf-e8e2-5693-be3c-320031f4850a"),
				IndexURI: uri.TrustedNew("file:///tmp/"),
				DataURI:  uri.TrustedNew("file:///tmp/data/"),
				MetaURI:  uri.TrustedNew("file:///tmp/meta/"),
			},
		},
	}
	for _, tt := range tests {
		got := NewDst(tt.indexURI, tt.dataURI, tt.metaURI)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q NewDst() got %#v, want %#v", tt.desc, got, tt.want)
		}
	}
}

func TestNewDstIDFromURIs(t *testing.T) {
	tests := []struct {
		desc     string
		indexURI uri.URI
		dataURI  uri.URI
		metaURI  uri.URI
		want     DstID
	}{
		{
			desc:     "conistent value",
			indexURI: uri.TrustedNew("file:///tmp/"),
			dataURI:  uri.TrustedNew("file:///tmp/data/"),
			metaURI:  uri.TrustedNew("file:///tmp/meta/"),
			want:     DstID("0b1bccaf-e8e2-5693-be3c-320031f4850a"),
		},
	}
	for _, tt := range tests {
		got := newDstIDFromURIs(tt.indexURI, tt.dataURI, tt.metaURI)
		if got != tt.want {
			t.Errorf("%q newDstIDFromURIs got %#v, want %#v", tt.desc, got, tt.want)
		}
	}
}
