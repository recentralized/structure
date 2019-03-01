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

// Read returns file contents.
var Read readFunc = defaultRead

type readFunc func(string) ([]byte, error)

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

func defaultRead(name string) ([]byte, error) {
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
