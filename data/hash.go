package data

import (
	"io"

	"github.com/recentralized/structure/cid"
)

// Hash is the unique identifier for a piece of data. A fingerprint calculated
// by its bytes. We use Hash to identify each unique piece of data stored by
// structure.
type Hash struct {
	cid cid.ContentID
}

var undefHash = Hash{}

const hashFormat = cid.SHA1

// ParseHash converts a string to a Hash.
func ParseHash(s string) (Hash, error) {
	cid, err := cid.Parse(s)
	if err != nil {
		return undefHash, err
	}
	return Hash{cid}, nil
}

// NewHash generates a Hash from the data in reader.
func NewHash(r io.Reader) (Hash, error) {
	cid, err := cid.New(r, hashFormat)
	if err != nil {
		return undefHash, err
	}
	return Hash{cid}, nil
}

// LiteralHash constructs a Hash whose value is literally the input. This is
// intended only for tests.
func LiteralHash(s string) Hash {
	return Hash{cid.NewLiteral(s)}
}

// Equal returns true if the two hashes represent the same data.
func (h Hash) Equal(hh Hash) bool {
	// NOTE: consider using EqualHash if we allow different cid formats.
	return h.cid.Equal(hh.cid)

}
func (h Hash) String() string {
	return h.cid.String()
}
