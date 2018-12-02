package data

import (
	"bytes"
	"testing"

	"github.com/recentralized/structure/cid"
)

func TestNewHash(t *testing.T) {
	if hashFormat != cid.SHA1 {
		t.Fatalf("hash format changed")
	}
	hash, err := NewHash(bytes.NewBufferString("testing 123"))
	if err != nil {
		t.Fatalf("failed to new: %s", err)
	}
	want, _ := ParseHash("b8dfb080bc33fb564249e34252bf143d88fc018f")
	if !hash.Equal(want) {
		t.Fatalf("got %s want %s", hash, want)
	}
}
