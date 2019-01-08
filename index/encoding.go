package index

import (
	"encoding/json"
	"time"

	"github.com/recentralized/structure/data"
	"github.com/recentralized/structure/uri"
)

type srcJSON struct {
	SrcID  SrcID   `json:"src_id"`
	SrcURI uri.URI `json:"src_uri"`
}

// MarshalJSON implements json.Marshaler.
func (s Src) MarshalJSON() ([]byte, error) {
	// Store other fields to the local representation. We can't use struct
	// embedding because it would trigger recursive marshal.
	sj := srcJSON{
		SrcID:  s.SrcID,
		SrcURI: s.SrcURI,
	}
	return json.Marshal(sj)
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *Src) UnmarshalJSON(data []byte) error {
	// Unmarshal basic fields.
	var sj srcJSON
	if err := json.Unmarshal(data, &sj); err != nil {
		return nil
	}
	// Set fields on Src.
	s.SrcID = sj.SrcID
	s.SrcURI = sj.SrcURI
	return nil
}

// srcItemJSON converts time.Time to *time.Time so the value is `null` in JSON.
type srcItemJSON struct {
	SrcID      SrcID      `json:"src_id"`
	DataURI    uri.URI    `json:"data_uri"`
	MetaURI    uri.URI    `json:"meta_uri"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	Modified   *time.Time `json:"modified,omitempty"`
}

// MarshalJSON implements json.Marshaler.
func (s SrcItem) MarshalJSON() ([]byte, error) {
	sj := srcItemJSON{
		SrcID:   s.SrcID,
		DataURI: s.DataURI,
		MetaURI: s.MetaURI,
	}
	if !s.ModifiedAt.IsZero() {
		sj.ModifiedAt = &s.ModifiedAt
	}
	return json.Marshal(sj)
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *SrcItem) UnmarshalJSON(data []byte) error {
	var sj srcItemJSON
	if err := json.Unmarshal(data, &sj); err != nil {
		return err
	}
	s.SrcID = sj.SrcID
	s.DataURI = sj.DataURI
	s.MetaURI = sj.MetaURI
	if sj.ModifiedAt != nil {
		s.ModifiedAt = *sj.ModifiedAt
	} else {
		// migrate: v0 to v1
		if sj.Modified != nil {
			s.ModifiedAt = *sj.Modified
		}
	}
	return nil
}

type dstJSON struct {
	DstID    DstID   `json:"dst_id"`
	IndexURI uri.URI `json:"index_uri"`
	DataURI  uri.URI `json:"data_uri"`
	MetaURI  uri.URI `json:"meta_uri"`
}

// MarshalJSON implements json.Marshaler.
func (s Dst) MarshalJSON() ([]byte, error) {
	sj := dstJSON{
		DstID:    s.DstID,
		IndexURI: s.IndexURI,
		DataURI:  s.DataURI,
		MetaURI:  s.MetaURI,
	}
	return json.Marshal(sj)
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *Dst) UnmarshalJSON(data []byte) error {
	var dj dstJSON
	if err := json.Unmarshal(data, &dj); err != nil {
		return nil
	}
	s.DstID = dj.DstID
	s.IndexURI = dj.IndexURI
	s.DataURI = dj.DataURI
	s.MetaURI = dj.MetaURI
	return nil
}

type dstItemJSON struct {
	DstID     DstID      `json:"dst_id"`
	DataURI   uri.URI    `json:"data_uri"`
	MetaURI   uri.URI    `json:"meta_uri"`
	DataSize  int64      `json:"data_size,omitempty"`
	MetaSize  int64      `json:"meta_size,omitempty"`
	StoredAt  *time.Time `json:"stored_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// MarshalJSON converts DstItem to JSON.
func (d DstItem) MarshalJSON() ([]byte, error) {
	j := dstItemJSON{
		DstID:    d.DstID,
		DataURI:  d.DataURI,
		MetaURI:  d.MetaURI,
		DataSize: d.DataSize,
		MetaSize: d.MetaSize,
	}
	if !d.StoredAt.IsZero() {
		j.StoredAt = &d.StoredAt
	}
	if !d.UpdatedAt.IsZero() {
		j.UpdatedAt = &d.UpdatedAt
	}
	return json.Marshal(j)
}

// UnmarshalJSON converts DstItem from JSON.
func (d *DstItem) UnmarshalJSON(b []byte) error {
	dj := dstItemJSON{}
	if err := json.Unmarshal(b, &dj); err != nil {
		return err
	}
	d.DstID = dj.DstID
	d.DataURI = dj.DataURI
	d.MetaURI = dj.MetaURI
	d.DataSize = dj.DataSize
	d.MetaSize = dj.MetaSize
	if dj.StoredAt != nil {
		d.StoredAt = *dj.StoredAt
	}
	if dj.UpdatedAt != nil {
		d.UpdatedAt = *dj.UpdatedAt
	}
	return nil
}

type uRefJSON struct {
	Hash data.Hash `json:"hash"`
	Srcs []SrcItem `json:"srcs"`
	Dsts []DstItem `json:"dsts"`
}

// MarshalJSON implements json.Marshaler.
func (r URef) MarshalJSON() ([]byte, error) {
	var rj uRefJSON
	rj.Hash = r.Hash
	if len(r.Srcs) == 0 {
		rj.Srcs = make([]SrcItem, 0)
	} else {
		rj.Srcs = r.Srcs
	}
	if len(r.Dsts) == 0 {
		rj.Dsts = make([]DstItem, 0)
	} else {
		rj.Dsts = r.Dsts
	}
	return json.Marshal(rj)
}

// UnmarshalJSON implements json.Unmarshaler.
func (r *URef) UnmarshalJSON(data []byte) error {
	var rj uRefJSON
	if err := json.Unmarshal(data, &rj); err != nil {
		return err
	}
	r.Hash = rj.Hash
	if len(rj.Srcs) > 0 {
		r.Srcs = rj.Srcs
	}
	if len(rj.Dsts) > 0 {
		r.Dsts = rj.Dsts
	}
	return nil
}
