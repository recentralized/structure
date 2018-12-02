package migration

import (
	"os"
	"reflect"
	"testing"

	"github.com/recentralized/structure/content"
)

func loadMeta(t *testing.T, path string) *content.Meta {
	fi, err := os.Open("_data/" + path)
	if err != nil {
		t.Fatalf("could not open file: %s", err)
		return nil
	}
	defer fi.Close()
	meta, err := content.ParseMetaJSON(fi)
	if err != nil {
		t.Fatalf("could not parse json: %s (%s)", err, path)
		return nil
	}
	return meta
}

// Verify that a v0 (unversioned) meta can be transparently migrated to v1.
func TestMetaV0toV1(t *testing.T) {
	v0 := loadMeta(t, "meta-v0.json")
	v1 := loadMeta(t, "meta-v1.json")
	if got, want := v0, v1; !reflect.DeepEqual(got, want) {
		t.Fatalf("not equal:\nv0: %+v\nv1: %+v", got, want)
	}
}
