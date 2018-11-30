package index

import (
	"encoding/json"
	"time"

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
	ModifiedAt *time.Time `json:"modified_at"`
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
	}
	return nil
}
