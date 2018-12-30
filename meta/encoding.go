package meta

import (
	"encoding/json"
	"time"

	"github.com/recentralized/structure/data"
	"github.com/recentralized/structure/index"
)

type metaJSON struct {
	Version     string    `json:"version"`
	Type        data.Type `json:"type"`
	ContentType data.Type `json:"content_type,omitempty"`
	Size        int64     `json:"size"`

	Inherent *Content                    `json:"inherent,omitempty"`
	Src      map[index.SrcID]SrcSpecific `json:"src,omitempty"`

	// V0 Fields
	Sidecar       *Content `json:"sidecar,omitempty"`
	V0SrcSpecific          // Embedded fields.
}

// MarshalJSON converts Meta to JSON.
func (m Meta) MarshalJSON() ([]byte, error) {
	j := metaJSON{
		Version:       m.Version,
		Type:          m.Type,
		Size:          m.Size,
		V0SrcSpecific: m.V0Srcs,
	}
	if !m.Inherent.isZero() {
		j.Inherent = &m.Inherent
	}
	if !m.V0Sidecar.isZero() {
		j.Sidecar = &m.V0Sidecar
	}
	if m.Src != nil {
		j.Src = make(map[index.SrcID]SrcSpecific)
		for k, v := range m.Src {
			j.Src[k] = v
		}
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
	m.Type = j.Type
	if m.Type == "" && j.ContentType != "" {
		m.Type = j.ContentType
	}
	m.Size = j.Size
	if j.Src != nil {
		m.Src = make(map[index.SrcID]SrcSpecific)
		for k, v := range j.Src {
			m.Src[k] = v
		}
	}
	if j.Inherent != nil {
		m.Inherent = *j.Inherent
	}
	if j.Sidecar != nil {
		m.V0Sidecar = *j.Sidecar
	}
	m.V0Srcs = j.V0SrcSpecific
	return nil
}

type metaContentJSON struct {
	Created *time.Time `json:"created,omitempty"`
	Image   *Image     `json:"image,omitempty"`
	Exif    Exif       `json:"exif,omitempty"`
}

// MarshalJSON converts MetaContent to JSON.
func (m Content) MarshalJSON() ([]byte, error) {
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
func (m *Content) UnmarshalJSON(data []byte) error {
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
