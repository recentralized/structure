package uri

import (
	"net/url"
	"testing"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		desc    string
		path    string
		wantStr string
		wantErr bool
		wantAbs bool
	}{
		{
			desc:    "blank string",
			path:    "  ",
			wantStr: "file://.",
			wantAbs: false,
		},
		{
			desc:    "single part path",
			path:    "path",
			wantStr: "file://path",
			wantAbs: false,
		},
		{
			desc:    "relative path",
			path:    "path/to/thing",
			wantStr: "file://path/to/thing",
			wantAbs: false,
		},
		{
			desc:    "rooted path",
			path:    "/abs/path",
			wantStr: "file:///abs/path",
			wantAbs: true,
		},
		{
			desc:    "funky path",
			path:    "/../abs/../path",
			wantStr: "file:///path",
			wantAbs: true,
		},
		{
			desc:    "file scheme",
			path:    "file:///abs/path",
			wantErr: true,
		},
		{
			desc:    "other scheme",
			path:    "http:///abs/path",
			wantErr: true,
		},
		{
			desc:    "path with encoding problems",
			path:    "/Photos Library.photoslibrary/Thumbnails/2015/09/23/20150923-010213/TqFU0duZTV+culxTIy%oVA/thumb_IMG_7220.jpg",
			wantStr: "file:///Photos%20Library.photoslibrary/Thumbnails/2015/09/23/20150923-010213/TqFU0duZTV+culxTIy%25oVA/thumb_IMG_7220.jpg",
			wantAbs: true,
		},
	}
	for _, tt := range tests {
		got, err := ParseFile(tt.path)
		if tt.wantErr {
			if err == nil {
				t.Errorf("%q NewFileFromPath() want error, got none", tt.desc)
			}
		} else {
			if err != nil {
				t.Errorf("%q NewFileFromPath() got err, want none: %s", tt.desc, err)
				continue
			}
			if gotStr := got.String(); gotStr != tt.wantStr {
				t.Errorf("%q NewFileFromPath() String()\ngot  %#v\nwant %#v", tt.desc, gotStr, tt.wantStr)
			}
			if got, want := got.IsAbs(), tt.wantAbs; got != want {
				t.Errorf("%q IsAbs got %t want %t", tt.desc, got, want)
			}
		}
	}
}

func TestParseDir(t *testing.T) {
	tests := []struct {
		desc    string
		path    string
		wantStr string
		wantErr bool
		wantAbs bool
	}{
		{
			desc:    "blank string",
			path:    "   ",
			wantStr: "file://./",
			wantAbs: false,
		},
		{
			desc:    "single part path",
			path:    "path",
			wantStr: "file://path/",
			wantAbs: false,
		},
		{
			desc:    "relative path without trailing slash",
			path:    "path/to/thing",
			wantStr: "file://path/to/thing/",
			wantAbs: false,
		},
		{
			desc:    "relative path with trailing slash",
			path:    "path/to/",
			wantStr: "file://path/to/",
			wantAbs: false,
		},
		{
			desc:    "absolute path without trailing slash",
			path:    "/abs/path",
			wantStr: "file:///abs/path/",
			wantAbs: true,
		},
		{
			desc:    "absolute path with trailing slash",
			path:    "/abs/path/",
			wantStr: "file:///abs/path/",
			wantAbs: true,
		},
		{
			desc:    "funky path",
			path:    "/../abs/../path/to",
			wantStr: "file:///path/to/",
			wantAbs: true,
		},
		{
			desc:    "file scheme",
			path:    "file:///abs/path",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := ParseDir(tt.path)
		if tt.wantErr {
			if err == nil {
				t.Errorf("%q NewFileDir() want error, got none", tt.desc)
			}
		} else {
			if err != nil {
				t.Errorf("%q NewFileDir() got error, want none: %s", tt.desc, err)
				continue
			}
			if gotStr := got.String(); gotStr != tt.wantStr {
				t.Errorf("%q NewFileDir() String() got %#v, want %#v", tt.desc, gotStr, tt.wantStr)
			}
			if got, want := got.IsAbs(), tt.wantAbs; got != want {
				t.Errorf("%q IsAbs got %t want %t", tt.desc, got, want)
			}
		}
	}
}
func TestIsAbs(t *testing.T) {
	tests := []struct {
		desc string
		path Path
		want bool
	}{
		{
			desc: "absolute path",
			path: Path{TrustedNew("file:///path")},
			want: true,
		},
		{
			desc: "relative path",
			path: Path{TrustedNew("file://path")},
			want: false,
		},
		{
			desc: "no url",
			path: Path{URI{}},
			want: false,
		},
	}
	for _, tt := range tests {
		got := tt.path.IsAbs()
		if got, want := got, tt.want; got != want {
			t.Errorf("%q IsAbs() got %t want %t", tt.desc, got, want)
		}
	}
}
func TestFilepath(t *testing.T) {
	tests := []struct {
		desc    string
		path    Path
		want    string
		wantErr bool
	}{
		{
			desc: "absolute file",
			path: Path{TrustedNew("file:///path")},
			want: "/path",
		},
		{
			desc: "absolute dir",
			path: Path{TrustedNew("file:///path/")},
			want: "/path",
		},
		{
			desc:    "relative path",
			path:    Path{TrustedNew("file://path")},
			wantErr: true,
		},
		{
			desc:    "wrong scheme",
			path:    Path{TrustedNew("http://path")},
			wantErr: true,
		},
		{
			desc:    "no url",
			path:    Path{URI{}},
			wantErr: true,
		},
		{
			desc: "path has spaces",
			path: Path{TrustedNew("file:///path with space")},
			want: "/path with space",
		},
		{
			desc: "path is encoded with spaces",
			path: Path{TrustedNew("file:///path%20with%20space")},
			want: "/path with space",
		},
		{
			desc:    "path has invalid encoding",
			path:    Path{URI{url: &url.URL{Scheme: "file", Path: "file:///path%2with%20invalid"}}},
			wantErr: true,
		},
		{
			desc: "path has funky references",
			path: Path{TrustedNew("file:///a/b/../c//d/./e")},
			want: "/a/c/d/e",
		},
		{
			desc: "path is complex with invalid encoding",
			path: Path{URI{url: &url.URL{Scheme: "file", Path: "/Photos Library.photoslibrary/Thumbnails/2015/09/23/20150923-010213/TqFU0duZTV+culxTIy%oVA/thumb_IMG_7220.jpg"}}},
			want: "/Photos Library.photoslibrary/Thumbnails/2015/09/23/20150923-010213/TqFU0duZTV+culxTIy%oVA/thumb_IMG_7220.jpg",
		},
	}
	for _, tt := range tests {
		got, err := tt.path.Filepath()
		if tt.wantErr {
			if err == nil {
				t.Errorf("%q Filepath() want error, got none", tt.desc)
			}
		} else {
			if err != nil {
				t.Errorf("%q Filepath() got error, want none: %s", tt.desc, err)
			}
			if got, want := got, tt.want; got != want {
				t.Errorf("%q Filepath()\ngot  %#v\nwant %#v", tt.desc, got, want)
			}
		}
	}
}
