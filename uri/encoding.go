package uri

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// MarshalJSON implements json.Marshaler.
func (u URI) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *URI) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != "" {
		// Ignore error, we'll get a rawStr uri back and that's fine.
		newURI, _ := New(s)
		*u = newURI
	}
	return nil
}

// MarshalJSON implements json.Marshaler.
func (p Path) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

// UnmarshalJSON implements json.Unmarshaler.
func (p *Path) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != "" {
		uri, _ := New(s)
		path, _ := ParseFileURI(uri)
		*p = path
	}
	return nil
}

// Value implements database/sql
func (u URI) Value() (driver.Value, error) {
	return []byte(u.String()), nil
}

// Scan implements database/sql
func (u *URI) Scan(data interface{}) error {
	v, ok := data.([]byte)
	if !ok {
		return fmt.Errorf("URI did not get bytes: %#v", data)
	}
	if len(v) != 0 {
		// Ignore error, we'll get a rawStr uri back and that's fine.
		newURI, _ := New(string(v))
		*u = newURI
	}
	return nil
}

// Value implements database/sql
func (p Path) Value() (driver.Value, error) {
	return []byte(p.String()), nil
}

// Scan implements database/sql
func (p *Path) Scan(data interface{}) error {
	v, ok := data.([]byte)
	if !ok {
		return fmt.Errorf("URI did not get bytes: %#v", data)
	}
	if len(v) != 0 {
		uri, _ := New(string(v))
		path, _ := ParseFileURI(uri)
		*p = path
	}
	return nil
}
