package uri

import (
	"net/url"
	"testing"
)

func TestNewFileFromPath(t *testing.T) {
	tests := []struct {
		desc    string
		path    string
		wantStr string
		wantErr bool
	}{
		{
			desc:    "blank string",
			path:    "  ",
			wantStr: "file://.",
		},
		{
			desc:    "relative path",
			path:    "path",
			wantStr: "file://path",
		},
		{
			desc:    "rooted path",
			path:    "/abs/path",
			wantStr: "file:///abs/path",
		},
		{
			desc:    "funky path",
			path:    "/../abs/../path",
			wantStr: "file:///path",
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
		},
	}
	for _, tt := range tests {
		got, err := NewFileFromPath(tt.path)
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
		}
	}
}

func TestNewDirFromPath(t *testing.T) {
	tests := []struct {
		desc    string
		path    string
		wantStr string
		wantErr bool
	}{
		{
			desc:    "blank string",
			path:    "   ",
			wantStr: "file://./",
		},
		{
			desc:    "relative path without trailing slash",
			path:    "path",
			wantStr: "file://path/",
		},
		{
			desc:    "relative path with trailing slash",
			path:    "path/",
			wantStr: "file://path/",
		},
		{
			desc:    "absolute path without trailing slash",
			path:    "/abs/path",
			wantStr: "file:///abs/path/",
		},
		{
			desc:    "absolute path with trailing slash",
			path:    "/abs/path/",
			wantStr: "file:///abs/path/",
		},
		{
			desc:    "funky path",
			path:    "/../abs/../path",
			wantStr: "file:///path/",
		},
		{
			desc:    "file scheme",
			path:    "file:///abs/path",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := NewDirFromPath(tt.path)
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
		}
	}
}
func TestIsAbs(t *testing.T) {
	newURL := func(str string) *url.URL {
		u, err := url.Parse(str)
		if err != nil {
			t.Fatalf("failed to parse URL: %s", err)
		}
		return u
	}
	tests := []struct {
		desc string
		url  *url.URL
		want bool
	}{
		{
			desc: "absolute path",
			url:  newURL("file:///path"),
			want: true,
		},
		{
			desc: "relative path",
			url:  newURL("file://path"),
			want: false,
		},
		{
			desc: "no url",
			url:  nil,
			want: false,
		},
	}
	for _, tt := range tests {
		uri := FileURI{URI{url: tt.url}}
		got := uri.IsAbs()
		if got, want := got, tt.want; got != want {
			t.Errorf("%q IsAbs() got %t want %t", tt.desc, got, want)
		}
	}
}
func TestFilePath(t *testing.T) {
	newURL := func(str string) *url.URL {
		u, err := url.Parse(str)
		if err != nil {
			t.Fatalf("failed to parse URL: %s", err)
		}
		return u
	}
	tests := []struct {
		desc    string
		url     *url.URL
		want    string
		wantErr bool
	}{
		{
			desc: "absolute file",
			url:  newURL("file:///path"),
			want: "/path",
		},
		{
			desc: "absolute dir",
			url:  newURL("file:///path/"),
			want: "/path",
		},
		{
			desc:    "relative path",
			url:     newURL("file://path"),
			wantErr: true,
		},
		{
			desc:    "wrong scheme",
			url:     newURL("http://path"),
			wantErr: true,
		},
		{
			desc:    "no url",
			url:     nil,
			wantErr: true,
		},
		{
			desc: "path has spaces",
			url:  newURL("file:///path with space"),
			want: "/path with space",
		},
		{
			desc: "path is encoded with spaces",
			url:  newURL("file:///path%20with%20space"),
			want: "/path with space",
		},
		{
			desc:    "path has invalid encoding",
			url:     &url.URL{Scheme: "file", Path: "file:///path%2with%20invalid"},
			wantErr: true,
		},
		{
			desc: "path has funky references",
			url:  newURL("file:///a/b/../c//d/./e"),
			want: "/a/c/d/e",
		},
		{
			desc: "path is complex with invalid encoding",
			url:  &url.URL{Scheme: "file", Path: "/Photos Library.photoslibrary/Thumbnails/2015/09/23/20150923-010213/TqFU0duZTV+culxTIy%oVA/thumb_IMG_7220.jpg"},
			want: "/Photos Library.photoslibrary/Thumbnails/2015/09/23/20150923-010213/TqFU0duZTV+culxTIy%oVA/thumb_IMG_7220.jpg",
		},
	}
	for _, tt := range tests {
		uri := FileURI{URI{url: tt.url}}
		got, err := uri.FilePath()
		if tt.wantErr {
			if err == nil {
				t.Errorf("%q FilePath() want error, got none", tt.desc)
			}
		} else {
			if err != nil {
				t.Errorf("%q FilePath() got error, want none: %s", tt.desc, err)
			}
			if got, want := got, tt.want; got != want {
				t.Errorf("%q FilePath()\ngot  %#v\nwant %#v", tt.desc, got, want)
			}
		}
	}
}
