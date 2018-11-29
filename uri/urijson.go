package uri

import "encoding/json"

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
