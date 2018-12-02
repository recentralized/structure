package cid

import (
	"crypto/sha1"
	"encoding/hex"
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
	str *string
	cid *cid.Cid
}

// Format is the kind of CID to generate.
type Format int

// Format options
const (
	SHA1 Format = iota
	CidV0
	CidV1
)

// Parse converts a string to a ContentID.
func Parse(s string) (ContentID, error) {
	if len(s) <= sha1Length {
		return ContentID{str: &s}, nil
	}
	decoded, err := cid.Decode(s)
	if err != nil {
		return ContentID{}, err
	}
	return ContentID{cid: &decoded}, nil
}

// New calculates a ContentID from data, in the given format.
func New(r io.Reader, fmt Format) (ContentID, error) {
	var builder cid.Builder
	switch fmt {
	case SHA1:
		hash, err := newSHA1(r)
		if err != nil {
			return ContentID{}, err
		}
		return ContentID{str: &hash}, nil
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
		return ContentID{}, err
	}
	v1, err := builder.Sum(data)
	if err != nil {
		return ContentID{}, err
	}
	return ContentID{cid: &v1}, nil
}

// NewLiteral constructs a ContentID whose value is literally the input. This
// is intended only for tests.
func NewLiteral(s string) ContentID {
	if len(s) >= sha1Length {
		panic(fmt.Sprintf("literal ContentID must be less than %d bytes", sha1Length))
	}
	return ContentID{str: &s}
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
	if c.str != nil {
		return *c.str
	}
	return ""
}

// Hash is the content hash string.
func (c ContentID) Hash() string {
	if c.cid != nil {
		hash := c.cid.Hash()
		return hash.B58String()
	}
	if c.str != nil {
		return *c.str
	}
	return ""
}

const sha1Length = 40

// newSHA1 calcualtes the original Hash content id.
func newSHA1(r io.Reader) (string, error) {
	sha := sha1.New()
	_, err := io.Copy(sha, r)
	if err != nil {
		return "", err
	}
	shaStr := hex.EncodeToString(sha.Sum(nil))
	return shaStr, nil
}
