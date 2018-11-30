package cid

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

// Hash is the original content identifier. Use cid.ContentID instead.
type Hash string

// newHash calcualtes the original content id. Use cid.New() instead.
func newHash(r io.Reader) (Hash, error) {
	sha := sha1.New()
	_, err := io.Copy(sha, r)
	if err != nil {
		return "", err
	}
	shaStr := hex.EncodeToString(sha.Sum(nil))
	return Hash(shaStr), nil
}

func (h Hash) String() string {
	return string(h)
}
