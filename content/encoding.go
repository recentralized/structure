package content

import (
	"encoding/json"
	"time"
)

type metaJSON struct {
	Version     string       `json:"version"`
	ContentType Type         `json:"content_type"`
	Size        int64        `json:"size"`
	Inherent    *MetaContent `json:"inherent,omitempty"`
	Sidecar     *MetaContent `json:"sidecar,omitempty"`
	SrcSpecific              // Embedded fields.
}

// MarshalJSON converts Meta to JSON.
func (m Meta) MarshalJSON() ([]byte, error) {
	j := metaJSON{
		Version:     m.Version,
		ContentType: m.ContentType,
		Size:        m.Size,
		SrcSpecific: m.Srcs,
	}
	if !m.Inherent.isZero() {
		j.Inherent = &m.Inherent
	}
	if !m.Sidecar.isZero() {
		j.Sidecar = &m.Sidecar
	}
	return json.Marshal(j)
}

// UnmarshalJSON converts JSON to Meta.
func (m *Meta) UnmarshalJSON(data []byte) error {
	var j metaJSON
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	m.Version = j.Version
	m.ContentType = j.ContentType
	m.Size = j.Size
	if j.Inherent != nil {
		m.Inherent = *j.Inherent
	}
	if j.Sidecar != nil {
		m.Sidecar = *j.Sidecar
	}
	m.Srcs = j.SrcSpecific
	return nil
}

type metaContentJSON struct {
	Created *time.Time `json:"created,omitempty"`
	Image   *MetaImage `json:"image,omitempty"`
	Exif    Exif       `json:"exif,omitempty"`
}

// MarshalJSON converts MetaContent to JSON.
func (m MetaContent) MarshalJSON() ([]byte, error) {
	j := metaContentJSON{}
	if !m.Created.IsZero() {
		j.Created = &m.Created
	}
	if !m.Image.isZero() {
		j.Image = &m.Image
	}
	if len(m.Exif) != 0 {
		j.Exif = m.Exif
	}
	return json.Marshal(j)
}

// UnmarshalJSON converts JSON to MetaContent.
func (m *MetaContent) UnmarshalJSON(data []byte) error {
	var j metaContentJSON
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	if j.Created != nil {
		m.Created = *j.Created
	}
	if j.Image != nil {
		m.Image = *j.Image
	}
	if j.Exif != nil {
		m.Exif = j.Exif
	}
	return nil
}
