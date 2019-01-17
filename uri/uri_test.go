package uri

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	_, parseErr := url.Parse("%")
	tests := []struct {
		str        string
		wantStr    string
		wantErr    error
		wantZero   bool
		wantNilURL bool
	}{
		{
			// Empty string is an error.
			str:      "",
			wantStr:  "",
			wantZero: true,
			wantErr:  errEmpty,
		},
		{
			// Blank string is an error.
			str:      "  ",
			wantStr:  "",
			wantZero: true,
			wantErr:  errEmpty,
		},
		{
			// Path only.
			str:     "/path",
			wantStr: "/path",
		},
		{
			// HTTP host no path.
			str:     "http://example.com",
			wantStr: "http://example.com",
		},
		{
			// HTTP host and path.
			str:     "http://example.com/",
			wantStr: "http://example.com/",
		},
		{
			// Normalized.
			str:     "HTTP://example.com/",
			wantStr: "http://example.com/",
		},
		{
			// Query
			str:     "http://example.com/path?xyz=789&abc=123",
			wantStr: "http://example.com/path?xyz=789&abc=123",
		},
		{
			// Parse error.
			str:        "%",
			wantStr:    "%",
			wantErr:    Error{Err: parseErr, invalid: true},
			wantNilURL: true,
		},
		{
			// AWS
			str:     "arn:aws:rds:eu-west-1:123456789012:db:mysql-db",
			wantStr: "arn:aws:rds:eu-west-1:123456789012:db:mysql-db",
		},
		{
			// Generic
			str:     "news:comp.infosystems.www.servers.unix",
			wantStr: "news:comp.infosystems.www.servers.unix",
		},
		{
			// Mail
			str:     "mailto:John.Doe@example.com",
			wantStr: "mailto:John.Doe@example.com",
		},
	}
	for _, tt := range tests {
		got, err := New(tt.str)
		if !reflect.DeepEqual(err, tt.wantErr) {
			t.Errorf("%q New() want %q, got %q", tt.str, err, tt.wantErr)
		}
		if gotStr := got.String(); gotStr != tt.wantStr {
			t.Errorf("%q New() String() got %#v, want %#v", tt.str, gotStr, tt.wantStr)
		}
		if tt.wantZero {
			if !got.IsZero() {
				t.Errorf("%q IsZero() got %t, want %t", tt.str, got.IsZero(), tt.wantZero)
			}
		} else {
			if got.IsZero() {
				t.Errorf("%q IsZero() got %t, want %t", tt.str, got.IsZero(), tt.wantZero)
			}
		}
		if tt.wantNilURL {
			if got.URL() != nil {
				t.Errorf("%q URL() must be nil", tt.str)
			}
		} else {
			if got.URL() == nil {
				t.Errorf("%q URL() must not be nil", tt.str)
			}
		}
	}
}
func TestNewFromURL(t *testing.T) {
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
		uri  URI
	}{
		{
			desc: "simple url",
			url:  newURL("http://example.com"),
			uri:  URI{url: newURL("http://example.com")},
		},
		{
			desc: "nil url",
			url:  nil,
			uri:  zero,
		},
	}
	for _, tt := range tests {
		got := NewFromURL(tt.url)
		if got, want := got, tt.uri; !got.Equal(want) {
			t.Errorf("%q NewFromURL()\ngot  %#v\nwant %#v", tt.desc, got, want)
		}
	}
}
func TestIsZero(t *testing.T) {
	newURL := func(str string) *url.URL {
		u, err := url.Parse(str)
		if err != nil {
			t.Fatalf("failed to parse URL: %s", err)
		}
		return u
	}
	tests := []struct {
		desc     string
		uri      URI
		wantZero bool
	}{
		{
			desc:     "zero is zero",
			wantZero: true,
		},
		{
			desc:     "zero var is zero",
			uri:      zero,
			wantZero: true,
		},
		{
			desc:     "empty url is zero",
			uri:      URI{url: newURL("")},
			wantZero: true,
		},
		{
			desc:     "non-empty url is non-zero",
			uri:      URI{url: newURL("ok")},
			wantZero: false,
		},
		{
			desc:     "any rawStr is non-zero",
			uri:      URI{rawStr: "ok"},
			wantZero: false,
		},
	}
	for _, tt := range tests {
		if got, want := tt.uri.IsZero(), tt.wantZero; got != want {
			t.Errorf("%q IsZero() got %t want %t", tt.desc, got, want)
		}
	}
}
func TestString(t *testing.T) {
	newURL := func(str string) *url.URL {
		u, err := url.Parse(str)
		if err != nil {
			t.Fatalf("failed to parse URL: %s", err)
		}
		return u
	}
	tests := []struct {
		desc   string
		rawStr string
		url    *url.URL
		want   string
	}{
		{
			desc:   "nothing is present",
			url:    nil,
			rawStr: "",
			want:   "",
		},
		{
			desc: "the url is present",
			url:  newURL("/path"),
			want: "/path",
		},
		{
			desc:   "the url is not present",
			rawStr: "%invalid url%",
			want:   "%invalid url%",
		},
	}
	for _, tt := range tests {
		uri := URI{url: tt.url, rawStr: tt.rawStr}
		if got, want := uri.String(), tt.want; got != want {
			t.Errorf("%q String() got %q, want %q", tt.desc, got, want)
		}
	}
}
func TestEqual(t *testing.T) {
	newURL := func(str string) *url.URL {
		u, err := url.Parse(str)
		if err != nil {
			t.Fatalf("failed to parse URL: %s", err)
		}
		return u
	}
	tests := []struct {
		desc string
		a    URI
		b    URI
		want bool
	}{
		{
			desc: "equal url",
			a:    URI{url: newURL("http://example.com")},
			b:    URI{url: newURL("http://example.com")},
			want: true,
		},
		{
			desc: "unequal url",
			a:    URI{url: newURL("http://example.com")},
			b:    URI{url: newURL("https://example.com")},
			want: false,
		},
		{
			desc: "equal rawStr",
			a:    URI{rawStr: "/path"},
			b:    URI{rawStr: "/path"},
			want: true,
		},
		{
			desc: "unequal rawStr",
			a:    URI{rawStr: "/path"},
			b:    URI{rawStr: "/paths"},
			want: false,
		},
		{
			desc: "equal rawStr and url",
			a:    URI{rawStr: "/path"},
			b:    URI{url: newURL("/path")},
			want: true,
		},
		{
			desc: "unequal rawStr and url",
			a:    URI{rawStr: "/paths"},
			b:    URI{url: newURL("/path")},
			want: false,
		},
		{
			desc: "spaces are not ignored",
			a:    URI{rawStr: "/path"},
			b:    URI{rawStr: " /path"},
			want: false,
		},
	}
	for _, tt := range tests {
		got := tt.a.Equal(tt.b)
		if got != tt.want {
			t.Errorf("%q Equal() got %t, want %t", tt.desc, got, tt.want)
		}
	}
}
func TestURL(t *testing.T) {
	newURL := func(str string) *url.URL {
		u, err := url.Parse(str)
		if err != nil {
			t.Fatalf("failed to parse URL: %s", err)
		}
		return u
	}
	tests := []struct {
		desc   string
		uri    URI
		hasURL bool
	}{
		{
			desc:   "with url",
			uri:    URI{url: newURL("http://example.com")},
			hasURL: true,
		},
		{
			desc:   "without url",
			uri:    URI{rawStr: "http://example.com"},
			hasURL: false,
		},
	}
	for _, tt := range tests {
		str := tt.uri.String()
		url := tt.uri.URL()

		if !tt.hasURL {
			if url != nil {
				t.Errorf("%q expects no url", tt.desc)
			}
			continue
		}
		if url == nil {
			t.Errorf("%q expects url", tt.desc)
			continue
		}
		if got, want := url.String(), str; got != want {
			t.Errorf("%q expect string to match got %s want %s", tt.desc, got, want)
		}
		// TEST IMMUTABILITY
		url.Scheme = "changed:"
		if got, want := url.String(), str; got == want {
			t.Errorf("%q expect mutated URL to change its string. got %s want %s", tt.desc, got, want)
		}
		if got, want := tt.uri.String(), str; got != want {
			t.Errorf("%q expect mutated URL not to change URI. got %s want %s", tt.desc, got, want)
		}
	}
}
func TestResolveReference(t *testing.T) {
	newURL := func(str string) *url.URL {
		u, err := url.Parse(str)
		if err != nil {
			t.Fatalf("failed to parse URL: %s", err)
		}
		return u
	}
	tests := []struct {
		desc    string
		base    URI
		ref     URI
		want    URI
		wantErr error
	}{
		{
			desc: "append url path",
			base: URI{url: newURL("http://example.com/")},
			ref:  URI{url: newURL("/path")},
			want: URI{url: newURL("http://example.com/path")},
		},
		{
			desc: "append absolute file path",
			base: URI{url: newURL("file:///root/a/")},
			ref:  URI{url: newURL("/path")},
			want: URI{url: newURL("file:///path")},
		},
		{
			desc: "append relative file path",
			base: URI{url: newURL("file:///root/a/")},
			ref:  URI{url: newURL("path")},
			want: URI{url: newURL("file:///root/a/path")},
		},
		{
			desc: "append to empty",
			base: URI{url: newURL("")},
			ref:  URI{url: newURL("/path")},
			want: URI{url: newURL("/path")},
		},
		{
			desc: "append empty",
			base: URI{url: newURL("http://example.com/")},
			ref:  URI{url: newURL("")},
			want: URI{url: newURL("http://example.com/")},
		},
		{
			desc: "append empty to empty",
			base: URI{url: newURL("")},
			ref:  URI{url: newURL("")},
			want: URI{url: newURL("")},
		},
		{
			desc:    "append valid url to invalid url",
			base:    URI{rawStr: "/something"},
			ref:     URI{url: newURL("/path")},
			wantErr: errInvalid,
		},
		{
			desc:    "append invalid url to valid url",
			base:    URI{url: newURL("/path")},
			ref:     URI{rawStr: "/something"},
			wantErr: errInvalid,
		},
		{
			desc:    "append invalid url to invalid url",
			base:    URI{rawStr: "/a"},
			ref:     URI{rawStr: "/b"},
			wantErr: errInvalid,
		},
		{
			desc:    "append empty to empty",
			base:    zero,
			ref:     zero,
			wantErr: errInvalid,
		},
	}
	for _, tt := range tests {
		got, err := tt.base.ResolveReference(tt.ref)
		if err != tt.wantErr {
			t.Errorf("%q ResolveReference got err %q want %q", tt.desc, err, tt.wantErr)
		}
		if !got.Equal(tt.want) {
			t.Errorf("%q ResolveReference() got %#v, want %#v", tt.desc, got, tt.want)
		}
	}
}
