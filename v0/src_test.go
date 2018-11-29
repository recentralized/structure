package structure

import (
	"reflect"
	"testing"

	"github.com/recentralized/structure/uri"
)

func TestNewSrc(t *testing.T) {
	tests := []struct {
		srcURI uri.URI
		want   Src
	}{
		{
			srcURI: uri.TrustedNew("file:///tmp/"),
			want: Src{
				SrcID:  SrcID("901e9ee6-64b8-5d44-941c-7e22c7cb4af3"),
				SrcURI: uri.TrustedNew("file:///tmp/"),
			},
		},
	}
	for i, test := range tests {
		got := NewSrc(test.srcURI)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d NewSrc() got %#v, want %#v", i, got, test.want)
		}
	}
}

func TestNewSrcIDFromURI(t *testing.T) {
	tests := []struct {
		srcURI uri.URI
		want   SrcID
	}{
		{
			srcURI: uri.TrustedNew("file:///tmp/"),
			want:   SrcID("901e9ee6-64b8-5d44-941c-7e22c7cb4af3"),
		},
	}
	for i, test := range tests {
		got := newSrcIDFromURI(test.srcURI)
		if got != test.want {
			t.Errorf("%d newSrcIDFromURIs got %#v, want %#v", i, got, test.want)
		}
	}
}
