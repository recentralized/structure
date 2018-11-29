package uri

import (
	"net/url"
	"testing"
)

func TestNew(t *testing.T) {
	empty1, _ := New("")
	empty2, _ := New("   ")
	if empty1 != Empty {
		t.Errorf("empty URI must equal the constant")
	}
	if empty2 != Empty {
		t.Errorf("blank URI must equal the constant")
	}
	tests := []struct {
		str     string
		wantStr string
		wantErr bool
	}{
		{
			// Empty string is ok.
			str:     "",
			wantStr: "",
		},
		{
			// Blank string is ok.
			str:     "  ",
			wantStr: "",
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
			str:     "%",
			wantStr: "%",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := New(tt.str)
		if tt.wantErr {
			if err == nil {
				t.Errorf("%q New() want error, got none", tt.str)
			}
		} else {
			if err != nil {
				t.Errorf("%q New() got Err, want none: %s", tt.str, err)
			}
		}
		if gotStr := got.String(); gotStr != tt.wantStr {
			t.Errorf("%q New() String() got %#v, want %#v", tt.str, gotStr, tt.wantStr)
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
			uri:  uri{url: newURL("http://example.com")},
		},
		{
			desc: "nil url",
			url:  nil,
			uri:  Empty,
		},
	}
	for _, tt := range tests {
		got := NewFromURL(tt.url)
		if got, want := got, tt.uri; !got.Equal(want) {
			t.Errorf("%q NewFromURL()\ngot  %#v\nwant %#v", tt.desc, got, want)
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
		uri := uri{url: tt.url, rawStr: tt.rawStr}
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
			a:    uri{url: newURL("http://example.com")},
			b:    uri{url: newURL("http://example.com")},
			want: true,
		},
		{
			desc: "unequal url",
			a:    uri{url: newURL("http://example.com")},
			b:    uri{url: newURL("https://example.com")},
			want: false,
		},
		{
			desc: "equal rawStr",
			a:    uri{rawStr: "/path"},
			b:    uri{rawStr: "/path"},
			want: true,
		},
		{
			desc: "unequal rawStr",
			a:    uri{rawStr: "/path"},
			b:    uri{rawStr: "/paths"},
			want: false,
		},
		{
			desc: "equal rawStr and url",
			a:    uri{rawStr: "/path"},
			b:    uri{url: newURL("/path")},
			want: true,
		},
		{
			desc: "unequal rawStr and url",
			a:    uri{rawStr: "/paths"},
			b:    uri{url: newURL("/path")},
			want: false,
		},
		{
			desc: "spaces are not ignored",
			a:    uri{rawStr: "/path"},
			b:    uri{rawStr: " /path"},
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
			uri:    uri{url: newURL("http://example.com")},
			hasURL: true,
		},
		{
			desc:   "without url",
			uri:    uri{rawStr: "http://example.com"},
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
	e := uri{url: newURL("")}
	re, _ := e.ResolveReference(e)
	if re != Empty {
		t.Errorf("Resolving to Empty must == Empty")
	}
	tests := []struct {
		desc    string
		base    URI
		ref     URI
		want    URI
		wantErr bool
	}{
		{
			desc: "append url path",
			base: uri{url: newURL("http://example.com/")},
			ref:  uri{url: newURL("/path")},
			want: uri{url: newURL("http://example.com/path")},
		},
		{
			desc: "append absolute file path",
			base: uri{url: newURL("file:///root/a/")},
			ref:  uri{url: newURL("/path")},
			want: uri{url: newURL("file:///path")},
		},
		{
			desc: "append relative file path",
			base: uri{url: newURL("file:///root/a/")},
			ref:  uri{url: newURL("path")},
			want: uri{url: newURL("file:///root/a/path")},
		},
		{
			desc: "append to empty",
			base: uri{url: newURL("")},
			ref:  uri{url: newURL("/path")},
			want: uri{url: newURL("/path")},
		},
		{
			desc: "append empty",
			base: uri{url: newURL("http://example.com/")},
			ref:  uri{url: newURL("")},
			want: uri{url: newURL("http://example.com/")},
		},
		{
			desc: "append empty to empty",
			base: uri{url: newURL("")},
			ref:  uri{url: newURL("")},
			want: uri{url: newURL("")},
		},
		{
			desc:    "append valid url to invalid url",
			base:    uri{rawStr: "/something"},
			ref:     uri{url: newURL("/path")},
			wantErr: true,
		},
		{
			desc:    "append invalid url to valid url",
			base:    uri{url: newURL("/path")},
			ref:     uri{rawStr: "/something"},
			wantErr: true,
		},
		{
			desc:    "append invalid url to invalid url",
			base:    uri{rawStr: "/a"},
			ref:     uri{rawStr: "/b"},
			wantErr: true,
		},
		{
			desc: "append empty to empty",
			base: Empty,
			ref:  Empty,
			want: Empty,
		},
	}
	for _, tt := range tests {
		got, err := tt.base.ResolveReference(tt.ref)
		if tt.wantErr {
			if err == nil {
				t.Errorf("%q wants error, got none", tt.desc)
			}
		} else {
			if err != nil {
				t.Errorf("%q wants error no error, got: %s", tt.desc, err)
			}
		}
		if err == nil {
			if !got.Equal(tt.want) {
				t.Errorf("%q ResolveReference() got %#v, want %#v", tt.desc, got, tt.want)
			}
		}
	}
}
