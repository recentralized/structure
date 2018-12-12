package data

import (
	"database/sql/driver"
	"fmt"
)

// MarshalJSON implements json.Marshaler.
func (h Hash) MarshalJSON() ([]byte, error) {
	return h.cid.MarshalJSON()
}

// UnmarshalJSON implements json.Unmarshaler.
func (h *Hash) UnmarshalJSON(data []byte) error {
	return h.cid.UnmarshalJSON(data)
}

// Value implements database/sql
func (h Hash) Value() (driver.Value, error) {
	return []byte(h.String()), nil
}

// Scan implements database/sql
func (h *Hash) Scan(data interface{}) error {
	v, ok := data.([]byte)
	if !ok {
		return fmt.Errorf("Hash did not get bytes: %#v", data)
	}
	o, err := ParseHash(string(v))
	if err != nil {
		return err
	}
	*h = o
	return nil
}
