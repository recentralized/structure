package data

import (
	"bytes"
	"reflect"
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

func TestHashEqual(t *testing.T) {
	h1, _ := NewHash(bytes.NewBufferString("a"))
	h2, _ := NewHash(bytes.NewBufferString("a"))
	h3, _ := NewHash(bytes.NewBufferString("b"))

	if !h1.Equal(h2) {
		t.Errorf("same data must be equal")
	}
	if h1.Equal(h3) {
		t.Errorf("different data must NOT be equal")
	}
}
func TestHashIsZero(t *testing.T) {
	tests := []struct {
		desc     string
		hash     Hash
		wantZero bool
	}{
		{
			desc:     "unspecified is zero",
			wantZero: true,
		},
		{
			desc:     "undef is zero",
			hash:     undefHash,
			wantZero: true,
		},
		{
			desc:     "value is not zero",
			hash:     LiteralHash("ok"),
			wantZero: false,
		},
	}
	for _, tt := range tests {
		if got, want := tt.hash.IsZero(), tt.wantZero; !reflect.DeepEqual(got, want) {
			t.Errorf("%q IsZero got %t want %t", tt.desc, got, want)
		}
	}
}
