package structure

import (
	"reflect"
	"testing"

	"github.com/recentralized/structure/uri"
)

func TestNewDst(t *testing.T) {
	tests := []struct {
		indexURI uri.URI
		dataURI  uri.URI
		metaURI  uri.URI
		want     Dst
	}{
		{
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
	for i, test := range tests {
		got := NewDst(test.indexURI, test.dataURI, test.metaURI)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d NewDst() got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestNewDstIDFromURIs(t *testing.T) {
	tests := []struct {
		indexURI uri.URI
		dataURI  uri.URI
		metaURI  uri.URI
		want     DstID
	}{
		{
			indexURI: uri.TrustedNew("file:///tmp/"),
			dataURI:  uri.TrustedNew("file:///tmp/data/"),
			metaURI:  uri.TrustedNew("file:///tmp/meta/"),
			want:     DstID("0b1bccaf-e8e2-5693-be3c-320031f4850a"),
		},
	}
	for i, test := range tests {
		got := newDstIDFromURIs(test.indexURI, test.dataURI, test.metaURI)
		if got != test.want {
			t.Errorf("%d newDstIDFromURIs got %#v, want %#v", i, got, test.want)
		}
	}
}
