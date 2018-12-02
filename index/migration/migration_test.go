package migration

import (
	"os"
	"reflect"
	"testing"

	"github.com/recentralized/structure/index"
)

func loadIndex(t *testing.T, path string) *index.Index {
	fi, err := os.Open("_data/" + path)
	if err != nil {
		t.Fatalf("could not open file: %s", err)
		return nil
	}
	defer fi.Close()
	idx, err := index.ParseJSON(fi)
	if err != nil {
		t.Fatalf("could not parse json: %s (%s)", err, path)
		return nil
	}
	return idx
}

// Verify that a v0 (unversioned) index can be transparently migrated to v1.
func TestIndexV0toV1(t *testing.T) {
	v0 := loadIndex(t, "v0.json")
	v1 := loadIndex(t, "v1.json")
	if got, want := v0, v1; !reflect.DeepEqual(got, want) {
		t.Fatalf("not equal:\nv0: %+v\nv1: %+v", got, want)
	}
}
