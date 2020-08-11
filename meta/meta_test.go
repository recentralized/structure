package meta

import (
	"reflect"
	"testing"
	"time"

	"github.com/recentralized/structure/index"
)

func TestMetaDateCreated(t *testing.T) {
	tests := []struct {
		desc string
		m    *Meta
		want time.Time
	}{
		{
			desc: "zero value",
			m:    &Meta{},
		},
		{
			desc: "inherent with date",
			m: &Meta{
				Inherent: Content{Created: time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC)},
			},
			want: time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC),
		},
		{
			desc: "oldest of inherent or sidecars",
			m: &Meta{
				Inherent: Content{Created: time.Date(2009, 1, 1, 1, 1, 1, 1, time.UTC)},
				Src: map[index.SrcID]SrcSpecific{
					index.SrcID("a"): {
						Sidecar: &Content{Created: time.Date(2004, 1, 1, 1, 1, 1, 1, time.UTC)},
					},
					index.SrcID("b"): {
						Sidecar: &Content{Created: time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC)},
					},
					index.SrcID("c"): {
						Sidecar: &Content{Created: time.Date(2003, 1, 1, 1, 1, 1, 1, time.UTC)},
					},
				},
			},
			want: time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC),
		},
		{
			desc: "v0: sidecar with date",
			m: &Meta{
				V0Sidecar: Content{Created: time.Date(2001, 2, 1, 1, 1, 1, 1, time.UTC)},
			},
			want: time.Date(2001, 2, 1, 1, 1, 1, 1, time.UTC),
		},
		{
			desc: "v0: prefers sidecar to inherent",
			m: &Meta{
				Inherent:  Content{Created: time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC)},
				V0Sidecar: Content{Created: time.Date(2002, 1, 1, 1, 1, 1, 1, time.UTC)},
			},
			want: time.Date(2002, 1, 1, 1, 1, 1, 1, time.UTC),
		},
	}
	for _, tt := range tests {
		got := tt.m.DateCreated()
		if got, want := got, tt.want; !reflect.DeepEqual(got, want) {
			t.Errorf("%q Meta.DateCreated()\ngot  %s\nwant %s", tt.desc, got, want)
		}
	}
}
func TestMetaImage(t *testing.T) {
	tests := []struct {
		desc string
		m    *Meta
		want Image
	}{
		{
			desc: "zero value",
			m:    &Meta{},
		},
		{
			desc: "inherent image",
			m: &Meta{
				Inherent: Content{Image: Image{Width: 100}},
			},
			want: Image{Width: 100},
		},
	}
	for _, tt := range tests {
		got := tt.m.Image()
		if got, want := got, tt.want; !reflect.DeepEqual(got, want) {
			t.Errorf("%q Meta.Image()\ngot  %#v\nwant %#v", tt.desc, got, want)
		}
	}
}
