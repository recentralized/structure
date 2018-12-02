package data

import (
	"io"

	"github.com/recentralized/structure/cid"
)

// Hash is the unique identifier for a piece of content. A fingerprint
// calculated by its bytes.
type Hash cid.ContentID

var undefHash = Hash{}

const hashFormat = cid.SHA1

// NewHash generates a Hash from the data in reader.
func NewHash(r io.Reader) (Hash, error) {
	cid, err := cid.New(r, hashFormat)
	if err != nil {
		return undefHash, err
	}
	return Hash(cid), nil
}

// ParseHash converts a string to a Hash.
func ParseHash(s string) (Hash, error) {
	cid, err := cid.Parse(s)
	if err != nil {
		return undefHash, err
	}
	return Hash(cid), nil
}

// Equal returns true if the two hashes are the same, meaning they represent
// the same data.
func (h Hash) Equal(hh Hash) bool {
	return cid.ContentID(h).Equal(cid.ContentID(hh))

}
func (h Hash) String() string {
	return cid.ContentID(h).String()
}
