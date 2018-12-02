package cid

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

// hash is the original content identifier. Use cid.ContentID instead.
type hash string

// hashes are always len() 40.
const legacyHashLen = 40

// newHash calcualtes the original content id. Use cid.New() instead.
func newHash(r io.Reader) (hash, error) {
	sha := sha1.New()
	_, err := io.Copy(sha, r)
	if err != nil {
		return "", err
	}
	shaStr := hex.EncodeToString(sha.Sum(nil))
	return hash(shaStr), nil
}

func (h hash) String() string {
	return string(h)
}
