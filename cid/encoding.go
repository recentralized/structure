package cid

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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

// Value implements database/sql
func (c ContentID) Value() (driver.Value, error) {
	return []byte(c.String()), nil
}

// Scan implements database/sql
func (c *ContentID) Scan(data interface{}) error {
	v, ok := data.([]byte)
	if !ok {
		return fmt.Errorf("ContentID did not get bytes: %#v", data)
	}
	o, err := Parse(string(v))
	if err != nil {
		return err
	}
	*c = o
	return nil
}
