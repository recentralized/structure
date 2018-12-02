package meta

import (
	"reflect"
	"testing"
	"time"
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
			desc: "sidecar with date",
			m: &Meta{
				Sidecar: Content{Created: time.Date(2001, 2, 1, 1, 1, 1, 1, time.UTC)},
			},
			want: time.Date(2001, 2, 1, 1, 1, 1, 1, time.UTC),
		},
		{
			desc: "prefers sidecar to inherent",
			m: &Meta{
				Inherent: Content{Created: time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC)},
				Sidecar:  Content{Created: time.Date(2001, 2, 1, 1, 1, 1, 1, time.UTC)},
			},
			want: time.Date(2001, 2, 1, 1, 1, 1, 1, time.UTC),
		},
	}
	for _, tt := range tests {
		got := tt.m.DateCreated()
		if got, want := got, tt.want; !reflect.DeepEqual(got, want) {
			t.Errorf("%q Meta.DateCreated()\ngot  %#v\nwant %#v", tt.desc, got, want)
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
