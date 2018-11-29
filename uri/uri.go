package uri

import (
	"fmt"
	"net/url"
	"strings"
)

var (
	// Empty is an empty URI. A good default when an empty URI is ok.
	Empty = uri{}
)

// URI represents all item locations across all sources and destinations. It is
// immutable, but can be modified using ResolveReference.
type URI interface {

	// String returns the string representation of the URI.
	String() string

	// Equal compares the string representation of another URI.
	Equal(URI) bool

	// URL returns a url.URL representation. It may be nil if the URI
	// cannot be handled by url.URL. Modifying the returned URL will *not*
	// alter this URI.
	URL() *url.URL

	// ResolveReference appends another URI, returning the resolved path.
	// Resolving uri.Empty with uri.Empty results in uri.Empty. Resolving
	// any non-valid URI results in an error.
	ResolveReference(URI) (URI, error)
}

// New parses str and returns a URI. If the input is an empty string or blank,
// the package variable url.Empty is returned, which you can do an equality
// test against. If there is a problem parsing the input, the error is
// returned, but so is a URI. You may choose whether you wish to proceed.
//
// uri, err := uri.New("http://www.example.com")
//
func New(str string) (URI, error) {
	str = strings.TrimSpace(str)
	if str == "" {
		return Empty, nil
	}
	url, err := url.Parse(str)
	if err != nil {
		return uri{rawStr: str}, err
	}
	return uri{url: url}, nil
}

// NewFromURL converts the URL to a URI. You can use this in tandem with
// URI.URL() to modify the URL and then create a new URI. Passing a nil URL
// will result in uri.Empty.
//
// uri, _ := uri.New("http://www.example.com")
// url := uri.URL()
// url.Scheme = "https:"
// secureURI := uri.NewFromURL(url)
//
func NewFromURL(url *url.URL) URI {
	if url == nil {
		return Empty
	}
	// Ignore error, assuming url.URL is always round-trippable.
	uri, _ := New(url.String())
	return uri
}

// uri implements URI. The zero value is safe to use.
//
// This implementation is a light wrapper around url.URL. If the URI cannot be
// represented by url.URL, it can be represented by rawStr but not all methods
// will work.
type uri struct {

	// url represents the URI for most cases.
	url *url.URL

	// rawStr is the original input if it could not be parsed into a URL.
	rawStr string
}

func (u uri) String() string {
	if u.url != nil {
		return u.url.String()
	}
	return u.rawStr
}
func (u uri) Equal(ref URI) bool {
	return u.String() == ref.String()
}
func (u uri) URL() *url.URL {
	if u.url == nil {
		return nil
	}
	url := *u.url
	return &url
}
func (u uri) ResolveReference(ref URI) (URI, error) {
	refURL := ref.URL()
	if u == Empty && ref == Empty {
		return Empty, nil
	}
	if u.url != nil && refURL != nil {
		url := u.url.ResolveReference(refURL)
		if url.String() == "" {
			return Empty, nil
		}
		return uri{url: url}, nil
	}
	return nil, fmt.Errorf("cannot resolve %q and %q", u, ref)
}
