package uri

import (
	"fmt"
	"net/url"
	"strings"
)

// URI represents all item locations across all sources and destinations. It is
// immutable, but can be modified using ResolveReference.
//
// The implementation is a light wrapper around url.URL. If the URI cannot be
// represented by url.URL, it can be represented by rawStr but not all methods
// will work.
type URI struct {

	// url represents the URI for most cases.
	url *url.URL

	// rawStr is the original input if it could not be parsed into a URL.
	rawStr string
}

// Error is the type of error returned by URI operations.
type Error struct {
	Msg string
	Err error
}

func (e Error) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("uri: %s", e.Msg)
	}
	return fmt.Sprintf("uri: %s (%s)", e.Msg, e.Err)
}

// ErrEmptyInput is returned if the input is empty. If a URI is returned along
// with this error, the URI is functional, but probably not what you want.
var ErrEmptyInput = Error{Msg: "input is empty"}

// ErrBadInput is returned if the input is invalid. If a URI is returned along
// with this error, the URI is functional but cannot be converted to a url.URL.
var ErrBadInput = Error{Msg: "input is invalid"}

// New parses str and returns a URI. If the input is an empty or blank string,
// it returns ErrEmptyInput. If any other problem occurs while parsing str it
// returns ErrBadInput. You may choose to ignore these errors, see the comments
// on the error values.
//
// uri, err := uri.New("http://www.example.com")
//
func New(str string) (URI, error) {
	str = strings.TrimSpace(str)
	if str == "" {
		return URI{url: &url.URL{}}, ErrEmptyInput
	}
	u, err := url.Parse(str)
	if err != nil {
		return URI{rawStr: str}, ErrBadInput
	}
	return URI{url: u}, nil
}

var zero = URI{}

// TrustedNew calls New and ignores any error. ONLY use this if you trust the
// input.
func TrustedNew(str string) URI {
	uri, _ := New(str)
	return uri
}

// NewFromURL converts the URL to a URI. You can use this in tandem with
// URI.URL() to modify the URL and then create a new URI. Passing a nil URL
// will result in the zero value.
//
// uri, _ := uri.New("http://www.example.com")
// url := uri.URL()
// url.Scheme = "https:"
// secureURI := uri.NewFromURL(url)
//
func NewFromURL(url *url.URL) URI {
	if url == nil {
		return zero
	}
	// Ignore error, assuming url.URL is always round-trippable.
	uri, _ := New(url.String())
	return uri
}

// String returns the string representation of the URI.
func (u URI) String() string {
	if u.url != nil {
		return u.url.String()
	}
	return u.rawStr
}

type uriStringer interface {
	uriString() string
}

func (u URI) uriString() string {
	return u.String()
}

// IsZero returns true if this URI is its zero value.
func (u URI) IsZero() bool {
	if u.url != nil {
		return u.url.String() == ""
	}
	return u.rawStr == ""
}

// Equal compares the string representation of another URI.
func (u URI) Equal(ref uriStringer) bool {
	return u.uriString() == ref.uriString()
}

// URL returns a url.URL representation. It may be nil if the URI could not be
// parsed. Modifying the returned URL will *not* alter this URI.
func (u URI) URL() *url.URL {
	if u.url == nil {
		return nil
	}
	url := *u.url
	return &url
}

// ResolveReference appends another URI, returning the resolved path.
// Resolving uri.Empty with uri.Empty results in uri.Empty. Resolving
// any non-valid URI results in an error.
func (u URI) ResolveReference(ref URI) (URI, error) {
	a := u.URL()
	b := ref.URL()
	if a == nil || b == nil {
		return zero, ErrBadInput
	}
	return URI{url: a.ResolveReference(b)}, nil
}
