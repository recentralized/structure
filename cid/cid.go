package cid

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
)

// ContentID is the identifier of a unique piece of content. Under the hood we
// use ipfs/cid which encodes additional format and versioning information in
// addition to a multihash of the data.
//
// Spec: https://github.com/ipld/cid
// Impl: https://github.com/ipfs/go-cid/
//
type ContentID struct {
	hash *hash
	cid  *cid.Cid
}

// Format is the kind of CID to generate.
type Format int

// Format options
const (
	Hash Format = iota
	CidV0
	CidV1
)

const defaultFormat = Hash

// Parse converts a string to a ContentID.
func Parse(s string) (ContentID, error) {
	if len(s) <= legacyHashLen {
		hash := hash(s)
		return ContentID{hash: &hash}, nil
	}
	decoded, err := cid.Decode(s)
	if err != nil {
		return ContentID{}, err
	}
	return ContentID{cid: &decoded}, nil
}

// New calculates a ContentID from data.
func New(r io.Reader) (ContentID, error) {
	return NewInFormat(r, defaultFormat)
}

// NewInFormat calculates a ContentID from data, in the given format.
func NewInFormat(r io.Reader, fmt Format) (ContentID, error) {
	var builder cid.Builder
	switch fmt {
	case Hash:
		hash, err := newHash(r)
		if err != nil {
			return ContentID{cid: &cid.Undef}, err
		}
		return ContentID{hash: &hash}, nil
	case CidV0:
		builder = cid.V0Builder{}
	case CidV1:
		builder = cid.V1Builder{
			Codec:  cid.Raw,
			MhType: mh.SHA2_256,
		}
	}
	// NOTE: is there any way to stream data in? Why isn't this a problem for IPFS?
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return ContentID{cid: &cid.Undef}, err
	}
	v1, err := builder.Sum(data)
	if err != nil {
		return ContentID{cid: &cid.Undef}, err
	}
	return ContentID{cid: &v1}, nil
}

// MustNewString creates a ContentID from string and panics if it fails.  This
// is intended only for tests.
func MustNewString(s string) ContentID {
	cid, err := New(bytes.NewBufferString(s))
	if err != nil {
		panic(fmt.Sprintf("failed to create content id: %s", err))
	}
	return cid
}

// NewLiteral constructs a ContentID whose value is literally the input. This
// is intended only for tests.
func NewLiteral(s string) ContentID {
	if len(s) >= legacyHashLen {
		panic(fmt.Sprintf("literal ContentID must be less than %d bytes", legacyHashLen))
	}
	h := hash(s)
	return ContentID{hash: &h}
}

// Equal determines if the two ContentID's are the same.
func (c ContentID) Equal(cc ContentID) bool {
	return c.String() == cc.String()
}

// EqualHash determines if the two ContentID's have the same content hash, even
// if they have different formats or other encodings.`
func (c ContentID) EqualHash(cc ContentID) bool {
	return c.Hash() == cc.Hash()
}

// String is the full ContentID string.
func (c ContentID) String() string {
	if c.cid != nil {
		return c.cid.String()
	}
	if c.hash != nil {
		return c.hash.String()
	}
	return ""
}

// Hash is the content hash string.
func (c ContentID) Hash() string {
	if c.cid != nil {
		hash := c.cid.Hash()
		return hash.B58String()
	}
	if c.hash != nil {
		return c.hash.String()
	}
	return ""
}
