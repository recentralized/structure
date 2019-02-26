package uri

import (
	"errors"
	"fmt"
	"net/url"
	neturl "net/url"
	"path/filepath"
	"regexp"
	"strings"
)

// ParseFile converts the path (abs or rel) to a file:// URI. If the path is
// empty or the input already contains a scheme (even file://) an error is
// returned. The path of the returned URI is normalized via filepath.Clean.
func ParseFile(path string) (Path, error) {
	var err error
	if path, err = cleanPath(path); err != nil {
		return Path{zero}, err
	}
	url := &url.URL{Scheme: "file", Path: path}
	return Path{URI{url: url}}, nil
}

// ParseDir converts the path (abs or rel) to a file:// URI, assuming that the
// path is intended to be a directory. Unlike files, directories always end
// with a slash ('/') in a URI. The path of the returned URI is normalized via
// filepath.Clean.
func ParseDir(path string) (Path, error) {
	var err error
	if path, err = cleanPath(path); err != nil {
		return Path{zero}, err
	}
	// NOTE: ASCII-Only. Is that ok?
	if path[len(path)-1:] != "/" {
		path = path + "/"
	}
	url := &url.URL{Scheme: "file", Path: path}
	return Path{URI{url: url}}, nil
}

func cleanPath(path string) (string, error) {
	path = strings.TrimSpace(path)
	if strings.Contains(path, "://") {
		return "", errors.New("must not include scheme")
	}
	return filepath.Clean(path), nil
}

// Path wraps URI, adding special handling for filesystem paths. To
// convert to standard URI, use path.URI.
type Path struct {
	URI
}

// IsAbs returns true if the path begins at root.
func (u Path) IsAbs() bool {
	url := u.URL()
	if url == nil {
		return false
	}
	return filepath.IsAbs(url.Path)
}

var encodePlus = regexp.MustCompile(`\+`)

// Filepath returns the absolute path for use on a filesystem. If the path
// is not absolute or the URI is not a "file" scheme" an error is returned.
// The resulting path is normalized via filepath.Clean.
func (u Path) Filepath() (string, error) {
	url := u.URL()
	if url == nil {
		return "", fmt.Errorf("missing url")
	}
	if url.Scheme != "file" {
		return "", fmt.Errorf("URI scheme is %s, want file", url.Scheme)
	}
	// If there is a literal "+" in the string, urlencode it,
	// otherwise it will get turned into a space by QueryUnescape.
	ep1 := url.EscapedPath()
	ep2 := encodePlus.ReplaceAllLiteralString(ep1, "%2B")
	path, err := neturl.QueryUnescape(ep2)
	if path == "" {
		return "", fmt.Errorf("path is empty")
	}
	if err != nil {
		return "", err
	}
	path = filepath.Clean(path)
	return path, nil
}
