// Package files provides supporting files for destination layouts.
package files

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// Read returns a file contained in this package.
func Read(name string) ([]byte, error) {
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
