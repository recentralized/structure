package uri

import (
	"errors"
	"testing"
)

func errEqual(e1, e2 error) bool {
	if e1 == nil && e2 == nil {
		return true
	}
	if e1 != nil && e2 != nil {
		return e1.Error() == e2.Error()
	}
	return false
}

func TestParseFile(t *testing.T) {
	tests := []struct {
		desc    string
		raw     string
		want    Path
		wantErr error
	}{
		{
			desc: "absolute path",
			raw:  "/tmp/foo",
			want: Path{
				RawPath: "/tmp/foo",
			},
		},
		{
			desc: "path ending in slash",
			raw:  "/tmp/foo/",
			want: Path{
				RawPath: "/tmp/foo",
			},
		},
		{
			desc: "path with extraneous space",
			raw:  "  /tmp/foo  ",
			want: Path{
				RawPath: "/tmp/foo",
			},
		},
		{
			desc: "path with extraneous parts",
			raw:  "/tmp/../foo",
			want: Path{
				RawPath: "/foo",
			},
		},
		{
			desc:    "input with scheme",
			raw:     "file:///tmp/foo",
			wantErr: errors.New("must not include scheme"),
		},
		{
			desc: "relative path",
			raw:  "tmp/foo",
			want: Path{
				RawPath: "tmp/foo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := ParseFile(tt.raw)
			if !errEqual(err, tt.wantErr) {
				t.Fatalf("Err got %v want %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Got %#v want %#v", got, tt.want)
			}
		})
	}
}

func TestParseDir(t *testing.T) {
	tests := []struct {
		desc    string
		raw     string
		want    Path
		wantErr error
	}{
		{
			desc: "absolute path",
			raw:  "/tmp/foo",
			want: Path{
				RawPath: "/tmp/foo",
				IsDir:   true,
			},
		},
		{
			desc: "path ending in slash",
			raw:  "/tmp/foo/",
			want: Path{
				RawPath: "/tmp/foo",
				IsDir:   true,
			},
		},
		{
			desc: "path with extraneous space",
			raw:  "  /tmp/foo  ",
			want: Path{
				RawPath: "/tmp/foo",
				IsDir:   true,
			},
		},
		{
			desc: "path with extraneous parts",
			raw:  "/tmp/../foo",
			want: Path{
				RawPath: "/foo",
				IsDir:   true,
			},
		},
		{
			desc: "input with scheme",
			raw:  "file:///tmp/foo",
			want: Path{
				IsDir: true,
			},
			wantErr: errors.New("must not include scheme"),
		},
		{
			desc: "relative path",
			raw:  "tmp/foo",
			want: Path{
				RawPath: "tmp/foo",
				IsDir:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := ParseDir(tt.raw)
			if !errEqual(err, tt.wantErr) {
				t.Fatalf("Err got %v want %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Got %#v want %#v", got, tt.want)
			}
		})
	}
}
