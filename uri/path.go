package uri

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

// ParseFile parses raw as a filesystem file.
func ParseFile(raw string) (Path, error) {
	path, err := cleanAbsPath(raw)
	if err != nil {
		return Path{"", false}, err
	}
	return Path{path, false}, nil
}

// ParseDir parses raw as a filesystem directory.
func ParseDir(raw string) (Path, error) {
	path, err := cleanAbsPath(raw)
	if err != nil {
		return Path{"", true}, err
	}
	return Path{path, true}, nil
}

// ParsePath converts a URI into a filesystem path.
func ParsePath(u URI) (Path, error) {
	url := u.URL()
	if url != nil {
		if url.Scheme != "file" {
			return Path{}, fmt.Errorf("scheme must be file")
		}
		if url.Host != "" {
			return Path{}, fmt.Errorf("host must be empty")
		}
		return parsePath(url.EscapedPath())
	}
	str := u.String()
	str = strings.TrimPrefix(str, "file://")
	return parsePath(str)
}

// Path is a filesystem path.
type Path struct {
	RawPath string
	IsDir   bool
}

func (p Path) String() string {
	u, _ := p.URI()
	return u.String()
}

// URI returns the path as a URI. A directory path will have "/" appended.
func (p Path) URI() (URI, error) {
	path := p.RawPath
	if p.IsDir {
		// NOTE: ASCII-Only. Is that ok?
		if path[len(path)-1:] != "/" {
			path = path + "/"
		}
	}
	sc := fmt.Sprintf("file://%s", path)
	uri, err := New(sc)
	if ee, ok := err.(Error); ok {
		if ee.IsInvalid() {
			return uri, nil
		}
		return uri, err
	}
	return uri, nil
}

// URL returns the path as a url.URL. It might be nil if the path cannot be
// represented by url.URL.
func (p Path) URL() *url.URL {
	u, _ := p.URI()
	return u.URL()
}

// Filepath returns a clean, absolute path on the filesystem.
func (p Path) Filepath() string {
	return filepath.Clean(p.RawPath)
}

func parsePath(raw string) (Path, error) {
	if raw[len(raw)-1:] == "/" {
		return ParseDir(raw)
	}
	return ParseFile(raw)
}

func cleanAbsPath(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if strings.Contains(raw, "://") {
		return "", errors.New("must not include scheme")
	}
	if !path.IsAbs(raw) {
		return "", errors.New("must be absolute")
	}
	return filepath.Clean(raw), nil
}
