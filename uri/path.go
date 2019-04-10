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

// ParsePath parses raw as a directory if it ends in "/", else as a file. If
// you're parsing user input and intend it to be a directory, prefer ParseDir.
func ParsePath(raw string) (Path, error) {
	if isDir(raw) {
		return ParseDir(raw)
	}
	return ParseFile(raw)
}

// ParseFileURI converts a URI into a filesystem path.
func ParseFileURI(u URI) (Path, error) {
	url := u.URL()
	if url != nil {
		if url.Scheme != "file" {
			return Path{}, fmt.Errorf("scheme must be file")
		}
		if url.Host != "" {
			return Path{}, fmt.Errorf("host must be empty")
		}
		return ParsePath(url.Path)
	}
	str := u.String()
	str = strings.TrimPrefix(str, "file://")
	return ParsePath(str)
}

// TrustedFile returns a Path. Only use this if you trust the input.
func TrustedFile(raw string) Path {
	path, _ := ParseFile(raw)
	return path
}

// TrustedDir returns a Path. Only use this if you trust input.
func TrustedDir(raw string) Path {
	path, _ := ParseDir(raw)
	return path
}

// TrustedFileURI returns a Path. Only use this if you trust the input.
func TrustedFileURI(u URI) Path {
	path, _ := ParseFileURI(u)
	return path
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

// IsZero returns true if the Path has no value.
func (p Path) IsZero() bool {
	return p.RawPath == ""
}

// URI returns the path as a URI. A directory path will have "/" appended.
func (p Path) URI() (URI, error) {
	path := p.RawPath
	if path == "" {
		return zero, nil
	}
	if p.IsDir {
		if !isDir(path) {
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
	raw := p.RawPath
	un, err := url.PathUnescape(raw)
	if err == nil {
		return filepath.Clean(un)
	}
	return filepath.Clean(raw)
}

func cleanAbsPath(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", errors.New("path is empty")
	}
	if strings.Contains(raw, "://") {
		return "", errors.New("must not include scheme")
	}
	if !path.IsAbs(raw) {
		return "", errors.New("must be absolute")
	}
	return filepath.Clean(raw), nil
}

func isDir(raw string) bool {
	if raw == "" {
		return false
	}
	// NOTE: ASCII-Only. Is that ok?
	return raw[len(raw)-1:] == "/"
}
