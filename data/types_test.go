package data

import (
	"errors"
	"reflect"
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

func TestEncoding(t *testing.T) {
	tests := []struct {
		desc    string
		enc     Encoding
		wantStr string
		wantExt string
	}{
		{
			desc:    "native",
			enc:     Native,
			wantStr: "<native>",
			wantExt: "",
		},
		{
			desc:    "tar",
			enc:     Tar,
			wantStr: "tar",
			wantExt: "tar",
		},
		{
			desc:    "gzip",
			enc:     GZip,
			wantStr: "gz",
			wantExt: "gz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got, want := tt.enc.String(), tt.wantStr; got != want {
				t.Errorf("String got %q want %q", got, want)
			}
			if got, want := tt.enc.Ext(), tt.wantExt; got != want {
				t.Errorf("Ext got %q want %q", got, want)
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
			got, err := ParseExt(tt.stored.String())
			if err != nil {
				t.Errorf("ParseExt round trip: %s", err)
			}
			if got != tt.stored {
				t.Errorf("ParseExt got %s want %s", got, tt.stored)
			}
		})
	}
}

func TestParseExt(t *testing.T) {
	tests := []struct {
		desc    string
		ext     string
		want    Stored
		wantErr error
	}{
		{

			desc: "just type",
			ext:  "jpg",
			want: Stored{JPG, Native},
		},
		{
			desc: "type and encoding",
			ext:  "jpg.gz",
			want: Stored{JPG, GZip},
		},
		{
			// NOTE: being generous for now. could definitely reconsider this.
			desc: "unknown type is allowed",
			ext:  "foo.gz",
			want: Stored{"foo", GZip},
		},
		{
			// NOTE: being generous for now. could definitely reconsider this.
			desc: "unknown encoding is allowed",
			ext:  "jpg.foo",
			want: Stored{JPG, "foo"},
		},
		{
			desc:    "too many parts (will support this later)",
			ext:     "jpg.tar.gz",
			wantErr: errors.New("data: too many parts in extension \"jpg.tar.gz\""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := ParseExt(tt.ext)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Fatalf("ParseExt got err %q want %q", got, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("ParseExt got %q want %q", got, tt.want)
			}
		})
	}
}
