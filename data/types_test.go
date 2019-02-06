package data

import (
	"errors"
	"fmt"
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
		wantFmtV  string
		wantClass Class
	}{
		{
			desc:      "Unknown",
			typ:       UnknownType,
			wantOk:    false,
			wantStr:   "",
			wantExt:   "",
			wantFmtV:  "unknown",
			wantClass: Unclassified,
		},
		{
			desc:      "other",
			typ:       "foo",
			wantOk:    false,
			wantStr:   "foo",
			wantExt:   ".foo",
			wantFmtV:  "foo",
			wantClass: Unclassified,
		},
		{
			desc:      "jpg",
			typ:       JPG,
			wantOk:    true,
			wantStr:   "jpg",
			wantExt:   ".jpg",
			wantFmtV:  "jpg",
			wantClass: Image,
		},
		{
			desc:      "png",
			typ:       PNG,
			wantOk:    true,
			wantStr:   "png",
			wantExt:   ".png",
			wantFmtV:  "png",
			wantClass: Image,
		},
		{
			desc:      "gif",
			typ:       GIF,
			wantOk:    true,
			wantStr:   "gif",
			wantExt:   ".gif",
			wantFmtV:  "gif",
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
			if got, want := fmt.Sprintf("%v", tt.typ), tt.wantFmtV; got != want {
				t.Errorf("%%v got %q want %q", got, want)
			}
		})
	}
}

func TestEncoding(t *testing.T) {
	tests := []struct {
		desc     string
		enc      Encoding
		wantOk   bool
		wantStr  string
		wantExt  string
		wantFmtV string
	}{
		{
			desc:     "native",
			enc:      Native,
			wantOk:   true,
			wantStr:  "",
			wantExt:  "",
			wantFmtV: "native",
		},
		{
			desc:     "undefined",
			enc:      "foo",
			wantOk:   false,
			wantStr:  "foo",
			wantExt:  ".foo",
			wantFmtV: "foo",
		},
		{
			desc:     "tar",
			enc:      Tar,
			wantOk:   true,
			wantStr:  "tar",
			wantExt:  ".tar",
			wantFmtV: "tar",
		},
		{
			desc:     "gzip",
			enc:      GZip,
			wantOk:   true,
			wantStr:  "gz",
			wantExt:  ".gz",
			wantFmtV: "gz",
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
			if got, want := fmt.Sprintf("%v", tt.enc), tt.wantFmtV; got != want {
				t.Errorf("%%v got %q want %q", got, want)
			}
		})
	}
}

func TestStored(t *testing.T) {
	tests := []struct {
		desc     string
		stored   Stored
		wantZero bool
		wantOk   bool
		wantStr  string
		wantExt  string
		wantFmtV string
	}{
		{
			desc:     "zero value",
			stored:   Stored{},
			wantOk:   false,
			wantZero: true,
			wantStr:  "",
			wantExt:  "",
			wantFmtV: "Stored[type: unknown, encoding: native]",
		},
		{
			desc:     "unknown type",
			stored:   Stored{Encoding: GZip},
			wantOk:   false,
			wantStr:  "gz",
			wantExt:  ".gz",
			wantFmtV: "Stored[type: unknown, encoding: gz]",
		},
		{
			desc:     "type with native encoding",
			stored:   Stored{Type: JPG},
			wantOk:   true,
			wantStr:  "jpg",
			wantExt:  ".jpg",
			wantFmtV: "Stored[type: jpg, encoding: native]",
		},
		{
			desc:     "type with encoding",
			stored:   Stored{PNG, GZip},
			wantOk:   true,
			wantStr:  "png.gz",
			wantExt:  ".png.gz",
			wantFmtV: "Stored[type: png, encoding: gz]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got, want := tt.stored.IsZero(), tt.wantZero; got != want {
				t.Errorf("IsZero() got %t want %t", got, want)
			}
			if got, want := tt.stored.Ok(), tt.wantOk; got != want {
				t.Errorf("Ok() got %t want %t", got, want)
			}
			if got, want := tt.stored.String(), tt.wantStr; got != want {
				t.Errorf("String got %q want %q", got, want)
			}
			if got, want := tt.stored.Ext(), tt.wantExt; got != want {
				t.Errorf("Ext got %q want %q", got, want)
			}
			if got, want := fmt.Sprintf("%v", tt.stored), tt.wantFmtV; got != want {
				t.Errorf("%%v got %q want %q", got, want)
			}
			ext := tt.stored.Ext()
			if tt.wantOk {
				got, err := ParseType(ext)
				if err != nil {
					t.Errorf("ParseType round trip: %s", err)
				}
				if got != tt.stored {
					t.Errorf("ParseType got %v want %v", got, tt.stored)
				}
			} else if ext != "" {
				_, err := ParseType(ext)
				if err == nil {
					t.Errorf("ParseType must fail if not empty and not ok: %s", ext)
				}
			}
		})
	}
}

func TestParseType(t *testing.T) {
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
			desc: "dot",
			ext:  ".",
			want: Stored{},
		},
		{
			desc: "type",
			ext:  "jpg",
			want: Stored{Type: JPG},
		},
		{
			desc: "type extension",
			ext:  ".jpg",
			want: Stored{Type: JPG},
		},
		{
			desc:    "encoding",
			ext:     "gz",
			want:    Stored{Encoding: GZip},
			wantErr: errors.New("data: unknown type: gz"),
		},
		{
			desc:    "encoding extension",
			ext:     ".gz",
			want:    Stored{Encoding: GZip},
			wantErr: errors.New("data: unknown type: gz"),
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
			wantErr: errors.New("data: unknown type: foo"),
		},
		{
			desc:    "unknown encoding",
			ext:     "jpg.foo",
			want:    Stored{JPG, "foo"},
			wantErr: errors.New("data: unknown encoding: foo"),
		},
		{
			desc:    "too many parts (will support this later)",
			ext:     "jpg.tar.gz",
			wantErr: errors.New("data: too many parts in extension \"jpg.tar.gz\""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := ParseType(tt.ext)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Fatalf("ParseType got err %q want %q", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("ParseType got %v want %v", got, tt.want)
			}
		})
	}
}
