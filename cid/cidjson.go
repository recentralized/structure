package cid

import "github.com/ipfs/go-cid"
import "encoding/json"

// MarshalJSON implements json.Marshaler.
func (c ContentID) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

// UnmarshalJSON implements json.Unmarshaler.
func (c *ContentID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if len(s) == legacyHashLen {
		hash := Hash(s)
		cid := ContentID{hash: &hash}
		*c = cid
		return nil
	}
	decoded, err := cid.Decode(s)
	if err != nil {
		return err
	}
	*c = ContentID{cid: &decoded}
	return nil
}
