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
	hash *Hash
	cid  *cid.Cid
}

type format int

const (
	hash format = iota
	cidV0
	cidV1
)

const defaultFormat = hash

// New calculates a ContentID from data.
func New(r io.Reader) (ContentID, error) {
	return newInFormat(r, defaultFormat)
}

func newInFormat(r io.Reader, fmt format) (ContentID, error) {
	var builder cid.Builder
	switch fmt {
	case hash:
		hash, err := newHash(r)
		if err != nil {
			return ContentID{cid: &cid.Undef}, err
		}
		return ContentID{hash: &hash}, nil
	case cidV0:
		builder = cid.V0Builder{}
	case cidV1:
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
