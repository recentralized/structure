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
		wantOk    bool
		wantStr   string
		wantExt   string
		wantClass Class
	}{
		{
			desc:      "Unknown",
			typ:       UnknownType,
			wantOk:    false,
			wantStr:   "data:unknown",
			wantExt:   "",
			wantClass: Unclassified,
		},
		{
			desc:      "other",
			typ:       "foo",
			wantOk:    false,
			wantStr:   "foo",
			wantExt:   "foo",
			wantClass: Unclassified,
		},
		{
			desc:      "jpg",
			typ:       JPG,
			wantOk:    true,
			wantStr:   "jpg",
			wantExt:   "jpg",
			wantClass: Image,
		},
		{
			desc:      "png",
			typ:       PNG,
			wantOk:    true,
			wantStr:   "png",
			wantExt:   "png",
			wantClass: Image,
		},
		{
			desc:      "gif",
			typ:       GIF,
			wantOk:    true,
			wantStr:   "gif",
			wantExt:   "gif",
			wantClass: Image,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got, want := tt.typ.Ok(), tt.wantOk; got != want {
				t.Errorf("Ok() got %t want %t", got, want)
			}
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
		wantOk  bool
		wantStr string
		wantExt string
	}{
		{
			desc:    "native",
			enc:     Native,
			wantOk:  true,
			wantStr: "data:native",
			wantExt: "",
		},
		{
			desc:    "undefined",
			enc:     "foo",
			wantOk:  false,
			wantStr: "foo",
			wantExt: "foo",
		},
		{
			desc:    "tar",
			enc:     Tar,
			wantOk:  true,
			wantStr: "tar",
			wantExt: "tar",
		},
		{
			desc:    "gzip",
			enc:     GZip,
			wantOk:  true,
			wantStr: "gz",
			wantExt: "gz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got, want := tt.enc.Ok(), tt.wantOk; got != want {
				t.Errorf("Ok() got %t want %t", got, want)
			}
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
		wantOk    bool
		wantStr   string
		wantExt   string
		wantClass Class
	}{
		{
			desc:    "zero value",
			stored:  Stored{},
			wantOk:  false,
			wantStr: "data:unknown",
			wantExt: "",
		},
		{
			desc:    "unknown type",
			stored:  Stored{Encoding: GZip},
			wantOk:  false,
			wantStr: "gz",
			wantExt: "gz",
		},
		{
			desc:    "type with native encoding",
			stored:  Stored{Type: JPG},
			wantOk:  true,
			wantStr: "jpg",
			wantExt: "jpg",
		},
		{
			desc:    "type with encoding",
			stored:  Stored{PNG, GZip},
			wantOk:  true,
			wantStr: "png.gz",
			wantExt: "png.gz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got, want := tt.stored.Ok(), tt.wantOk; got != want {
				t.Errorf("Ok() got %t want %t", got, want)
			}
			if got, want := tt.stored.String(), tt.wantStr; got != want {
				t.Errorf("String got %q want %q", got, want)
			}
			if got, want := tt.stored.Ext(), tt.wantExt; got != want {
				t.Errorf("Ext got %q want %q", got, want)
			}
			ext := tt.stored.Ext()
			if tt.wantOk {
				got, err := ParseExt(ext)
				if err != nil {
					t.Errorf("ParseExt round trip: %s", err)
				}
				if got != tt.stored {
					t.Errorf("ParseExt got %v want %v", got, tt.stored)
				}
			} else if ext != "" {
				_, err := ParseExt(ext)
				if err == nil {
					t.Errorf("ParseExt must fail if not empty and not ok: %s", ext)
				}
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
			desc: "empty string",
			ext:  "",
			want: Stored{},
		},
		{
			desc: "just type",
			ext:  "jpg",
			want: Stored{Type: JPG},
		},
		{
			desc:    "just encoding",
			ext:     "gz",
			want:    Stored{Encoding: GZip},
			wantErr: errors.New("data: unknown type: \"gz\""),
		},
		{
			desc: "type and encoding",
			ext:  "jpg.gz",
			want: Stored{JPG, GZip},
		},
		{
			desc:    "unknown type",
			ext:     "foo.gz",
			want:    Stored{"foo", GZip},
			wantErr: errors.New("data: unknown type: \"foo\""),
		},
		{
			desc:    "unknown encoding",
			ext:     "jpg.foo",
			want:    Stored{JPG, "foo"},
			wantErr: errors.New("data: unknown encoding: \"foo\""),
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
				t.Fatalf("ParseExt got err %q want %q", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("ParseExt got %v want %v", got, tt.want)
			}
		})
	}
}
