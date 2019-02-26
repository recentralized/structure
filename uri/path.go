package uri

import (
	"fmt"
	"net/url"
	"path"
	"path/filepath"
)

// ParseFile parses raw as a filesystem file.
func ParseFile(raw string) (Path, error) {
	path, err := cleanPath(raw)
	if err != nil {
		return Path{"", false}, err
	}
	return Path{path, false}, nil
}

// ParseDir parses raw as a filesystem directory.
func ParseDir(raw string) (Path, error) {
	path, err := cleanPath(raw)
	if err != nil {
		return Path{"", true}, err
	}
	return Path{path, true}, nil
}

// Path is a filesystem path.
type Path struct {
	RawPath string
	IsDir   bool
}

// IsAbs returns true if the path starts at root.
func (p Path) IsAbs() bool {
	return path.IsAbs(p.RawPath)
}

// URI returns the path as a URI.
func (p Path) URI() URI {
	return URI{url: p.URL()}
}

// URL returns the path as a url.URL. A directory path will have "/" appended.
func (p Path) URL() *url.URL {
	path := p.RawPath
	if p.IsDir {
		// NOTE: ASCII-Only. Is that ok?
		if path[len(path)-1:] != "/" {
			path = path + "/"
		}
	}
	return &url.URL{Scheme: "file", Path: path}
}

// Filepath returns a clean, absolute path on the filesystem.
func (p Path) Filepath() (string, error) {
	if !p.IsAbs() {
		return "", fmt.Errorf("path is not absolute: %s", p.RawPath)
	}
	return filepath.Clean(p.RawPath), nil
}
