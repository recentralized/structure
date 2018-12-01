package cid

import (
	"encoding/json"
	"strings"
)

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
	if strings.TrimSpace(s) == "" {
		return nil
	}
	cid, err := Parse(s)
	if err != nil {
		return err
	}
	*c = cid
	return nil
}
