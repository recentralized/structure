// Package files provides supporting files for destination layouts.
package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type readFunc func(string) ([]byte, error)

// Read returns file contents. It's defined as a variable so that it may be
// replaced with another implementation. For example to support distributing a
// binary where these files don't exist on disk.
var Read readFunc = func(name string) ([]byte, error) {
	path := filepath.Join(getPath(), name)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// List returns a list of the non-go file names contained in this package.
func List() []string {
	out := make([]string, 0)
	files, err := filepath.Glob(path.Join(getPath(), "*"))
	if err != nil {
		panic(fmt.Sprintf("could not glob files: %s", err))
	}
	for _, f := range files {
		if path.Ext(f) != ".go" {
			out = append(out, path.Base(f))
		}
	}
	return out
}

var pkgPath string

func getPath() string {
	if pkgPath == "" {
		_, f, _, ok := runtime.Caller(1)
		if !ok {
			panic("failed to get `files` package path")
		}
		pkgPath = path.Dir(f)
	}
	return pkgPath
}
