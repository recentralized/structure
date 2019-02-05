package data

import (
	"database/sql/driver"
	"encoding/json"
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
	if len(v) > 0 {
		o, err := ParseHash(string(v))
		if err != nil {
			return err
		}
		*h = o
	}
	return nil
}

// MarshalJSON implements json.Marshaler.
func (s Stored) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *Stored) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	o, err := ParseType(str)
	if err != nil {
		return err
	}
	*s = o
	return nil
}

// Value implements database/sql
func (s Stored) Value() (driver.Value, error) {
	return []byte(s.String()), nil
}

// Scan implements database/sql
func (s *Stored) Scan(data interface{}) error {
	v, ok := data.([]byte)
	if !ok {
		return fmt.Errorf("Stored did not get bytes: %#v", data)
	}
	if len(v) > 0 {
		o, err := ParseType(string(v))
		if err != nil {
			return err
		}
		*s = o
	}
	return nil
}
