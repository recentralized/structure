package data

import (
	"testing"
)

func TestType(t *testing.T) {
	tests := []struct {
		desc      string
		typ       Type
		wantStr   string
		wantExt   string
		wantClass Class
	}{
		{
			desc:      "jpg",
			typ:       JPG,
			wantStr:   "jpg",
			wantExt:   "jpg",
			wantClass: Image,
		},
		{
			desc:      "png",
			typ:       PNG,
			wantStr:   "png",
			wantExt:   "png",
			wantClass: Image,
		},
		{
			desc:      "gif",
			typ:       GIF,
			wantStr:   "gif",
			wantExt:   "gif",
			wantClass: Image,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got, want := tt.typ.String(), tt.wantStr; got != want {
				t.Errorf("String got %q want %q", got, want)
			}
			if got, want := tt.typ.Ext(), tt.wantExt; got != want {
				t.Errorf("Ext got %q want %q", got, want)
			}
			if got, want := tt.typ.Class(), tt.wantClass; got != want {
				t.Errorf("Class got %q want %q", got, want)
			}
		})
	}
}

func TestStored(t *testing.T) {
	tests := []struct {
		desc      string
		stored    Stored
		wantStr   string
		wantExt   string
		wantClass Class
	}{
		{
			desc:    "native encoding",
			stored:  Stored{JPG, Native},
			wantStr: "jpg",
			wantExt: "jpg",
		},
		{
			desc:    "with encoding",
			stored:  Stored{PNG, GZip},
			wantStr: "png.gz",
			wantExt: "png.gz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got, want := tt.stored.String(), tt.wantStr; got != want {
				t.Errorf("String got %q want %q", got, want)
			}
			if got, want := tt.stored.Ext(), tt.wantExt; got != want {
				t.Errorf("Ext got %q want %q", got, want)
			}
		})
	}
}
