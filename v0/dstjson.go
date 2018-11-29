package structure

import (
	"encoding/json"
	"time"
)

type itemCopy DstItem

type dstItemJSON struct {
	itemCopy
	StoredAt *time.Time `json:"stored_at"`
}

// MarshalJSON converts DstItem to JSON.
func (d DstItem) MarshalJSON() ([]byte, error) {
	j := dstItemJSON{itemCopy: itemCopy(d)}
	if !d.StoredAt.IsZero() {
		j.StoredAt = &d.StoredAt
	}
	return json.Marshal(j)
}

// UnmarshalJSON converts DstItem from JSON.
func (d *DstItem) UnmarshalJSON(b []byte) error {
	j := dstItemJSON{}
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	*d = DstItem(j.itemCopy)
	if j.StoredAt != nil {
		d.StoredAt = *j.StoredAt
	}
	return nil
}
