package structure

import (
	"reflect"
	"testing"

	"github.com/recentralized/structure/uri"
)

func TestNewSrc(t *testing.T) {
	tests := []struct {
		desc   string
		srcURI uri.URI
		want   Src
	}{
		{
			desc:   "consistent value",
			srcURI: uri.TrustedNew("file:///tmp/"),
			want: Src{
				SrcID:  SrcID("901e9ee6-64b8-5d44-941c-7e22c7cb4af3"),
				SrcURI: uri.TrustedNew("file:///tmp/"),
			},
		},
	}
	for _, tt := range tests {
		got := NewSrc(tt.srcURI)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q NewSrc() got %#v, want %#v", tt.desc, got, tt.want)
		}
	}
}

func TestNewSrcIDFromURI(t *testing.T) {
	tests := []struct {
		desc   string
		srcURI uri.URI
		want   SrcID
	}{
		{
			desc:   "consistent value",
			srcURI: uri.TrustedNew("file:///tmp/"),
			want:   SrcID("901e9ee6-64b8-5d44-941c-7e22c7cb4af3"),
		},
	}
	for _, tt := range tests {
		got := newSrcIDFromURI(tt.srcURI)
		if got != tt.want {
			t.Errorf("%q newSrcIDFromURIs got %#v, want %#v", tt.desc, got, tt.want)
		}
	}
}
