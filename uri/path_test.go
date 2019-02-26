package uri

import (
	"errors"
	"net/url"
	"reflect"
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
			desc:    "relative path",
			raw:     "tmp/foo",
			want:    Path{},
			wantErr: errors.New("must be absolute"),
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
				IsDir: true,
			},
			wantErr: errors.New("must be absolute"),
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

func TestPathURI(t *testing.T) {
	tests := []struct {
		desc          string
		path          Path
		wantURI       URI
		wantErr       error
		wantURIString string
	}{
		{
			desc: "file path",
			path: Path{
				RawPath: "/tmp/file",
			},
			wantURI: URI{
				url: &url.URL{
					Scheme: "file",
					Path:   "/tmp/file",
				},
			},
			wantURIString: "file:///tmp/file",
		},
		{
			desc: "dir path",
			path: Path{
				RawPath: "/tmp/file",
				IsDir:   true,
			},
			wantURI: URI{
				url: &url.URL{
					Scheme: "file",
					Path:   "/tmp/file/",
				},
			},
			wantURIString: "file:///tmp/file/",
		},
		{
			desc: "encoded path",
			path: Path{
				RawPath: "/tmp/file%20with%20space",
			},
			wantURI: URI{
				url: &url.URL{
					Scheme: "file",
					Path:   "/tmp/file with space",
				},
			},
			wantURIString: "file:///tmp/file%20with%20space",
		},
		{
			desc: "badly encoded path",
			path: Path{
				RawPath: "/tmp/file%2with%20space",
			},
			wantURI: URI{
				rawStr: "file:///tmp/file%2with%20space",
			},
			wantURIString: "file:///tmp/file%2with%20space",
		},
		{
			desc: "path is complex with invalid encoding",
			path: Path{
				RawPath: "/Photos Library.photoslibrary/Thumbnails/2015/09/23/20150923-010213/TqFU0duZTV+culxTIy%oVA/thumb_IMG_7220.jpg",
			},
			wantURI: URI{
				rawStr: "file:///Photos Library.photoslibrary/Thumbnails/2015/09/23/20150923-010213/TqFU0duZTV+culxTIy%oVA/thumb_IMG_7220.jpg",
			},
			wantURIString: "file:///Photos Library.photoslibrary/Thumbnails/2015/09/23/20150923-010213/TqFU0duZTV+culxTIy%oVA/thumb_IMG_7220.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := tt.path.URI()
			if !errEqual(err, tt.wantErr) {
				t.Fatalf("Err got %v want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.wantURI) {
				t.Errorf("URL\ngot  %#v\nwant %#v", got, tt.wantURI)
			}
			if got, want := got.String(), tt.wantURIString; got != want {
				t.Errorf("URL String\ngot  %q\nwant %q", got, want)
			}
		})
	}
}

func TestPathFilepath(t *testing.T) {
	tests := []struct {
		desc string
		path Path
		want string
	}{
		{
			desc: "file path",
			path: Path{
				RawPath: "/tmp/file",
			},
			want: "/tmp/file",
		},
		{
			desc: "dir path",
			path: Path{
				RawPath: "/tmp/dir",
				IsDir:   true,
			},
			want: "/tmp/dir",
		},
		{
			desc: "path is complex with invalid encoding",
			path: Path{
				RawPath: "/Photos Library.photoslibrary/Thumbnails/2015/09/23/20150923-010213/TqFU0duZTV+culxTIy%oVA/thumb_IMG_7220.jpg",
			},
			want: "/Photos Library.photoslibrary/Thumbnails/2015/09/23/20150923-010213/TqFU0duZTV+culxTIy%oVA/thumb_IMG_7220.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := tt.path.Filepath()
			if got, want := got, tt.want; got != want {
				t.Errorf("Filepath\ngot  %q\nwant %q", got, want)
			}
		})
	}
}
