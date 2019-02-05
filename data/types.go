package data

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrUnknownType is returned if the type cannot be determined.
	ErrUnknownType = errors.New("data: unknown type")
)

// Type definitions.
const (
	// UnknownType is the zero value for Type, meaning it is unknown.
	UnknownType Type = ""

	// Image formats.
	JPG = "jpg" // Standard JPG file.
	PNG = "png" // Standard PNG file.
	GIF = "gif" // Standard GIF file.
)

// Encoding definitions.
const (
	// Native is the zero value for Encoding, meaning no encoding.
	Native Encoding = ""

	Tar  = "tar"
	GZip = "gz"
)

// Class definitions.
const (
	Unclassified Class = ""
	Image              = "image"
)

// Type is a known kind of file such as JPEG or PNG.
type Type string

func (t Type) String() string {
	if t == "" {
		return "data:unknown"
	}
	return string(t)
}

// Ext returns the type's standard file extension.
func (t Type) Ext() string {
	if t == "" {
		return ""
	}
	return fmt.Sprintf(".%s", t)
}

// Class returns the type's class: image, catalog, etc.
func (t Type) Class() Class {
	return classOfType[t]
}

// Ok return true if the type is defined.
func (t Type) Ok() bool {
	_, ok := types[t]
	return ok
}

// ParseExt parses an extension, returning the Stored format.
func ParseExt(str string) (Stored, error) {
	str = strings.TrimPrefix(str, ".")
	if str == "" {
		return Stored{}, nil
	}
	parts := strings.Split(str, ".")
	if len(parts) > 2 {
		return Stored{}, fmt.Errorf("data: too many parts in extension %q", str)
	}
	if len(parts) == 1 {
		if t := Type(parts[0]); t.Ok() {
			return Stored{Type: t}, nil
		}
		e := Encoding(parts[0])
		return Stored{Encoding: e}, fmt.Errorf("data: unknown type: %q", str)
	}
	t := Type(parts[0])
	e := Native
	if len(parts) == 2 {
		e = Encoding(parts[1])
	}
	s := Stored{t, e}
	if !t.Ok() {
		return s, fmt.Errorf("data: unknown type: %q", t)
	}
	if !e.Ok() {
		return s, fmt.Errorf("data: unknown encoding: %q", e)
	}
	return s, nil
}

// Stored is how a type is formatted for storage.
type Stored struct {
	Type     Type
	Encoding Encoding
}

func (s Stored) String() string {
	e := strings.TrimPrefix(s.Ext(), ".")
	if e == "" {
		return "data:unknown"
	}
	return e
}

// Ext returns the stored data's standard file extension.
func (s Stored) Ext() string {
	switch {
	case s.Type == UnknownType && s.Encoding == Native:
		return ""
	case s.Type == UnknownType:
		return s.Encoding.Ext()
	case s.Encoding == Native:
		return s.Type.Ext()
	default:
		return fmt.Sprintf("%s%s", s.Type.Ext(), s.Encoding.Ext())
	}
}

// Ok return true if the storage format is defined.
func (s Stored) Ok() bool {
	return s.Type.Ok() && s.Encoding.Ok()
}

// Format implements fmt.Formatter.
func (s Stored) Format(f fmt.State, c rune) {
	switch c {
	case 's':
		f.Write([]byte(s.String()))
	case 'v':
		f.Write([]byte("Stored["))
		f.Write([]byte("type: "))
		f.Write([]byte(s.Type.String()))
		f.Write([]byte(", encoding: "))
		f.Write([]byte(s.Encoding.String()))
		f.Write([]byte("]"))
	}
}

// Encoding is the encoding of the data for storage. Most types are stored in
// their native encoding, but we may want to optimize storage by compressing or
// flattening multi-file structures.
type Encoding string

func (e Encoding) String() string {
	if e == "" {
		return "data:native"
	}
	return string(e)
}

// Ext returns the encoding's standard file extension.
func (e Encoding) Ext() string {
	if e == "" {
		return ""
	}
	return fmt.Sprintf(".%s", e)
}

// Ok return true if the encoding is defined.
func (e Encoding) Ok() bool {
	_, ok := encodings[e]
	return ok
}

// Class is the category of data that a type belongs to. JPG, PNG, GIF are all
// image, etc.
type Class string

func (c Class) String() string {
	return string(c)
}

var types = map[Type]bool{
	JPG: true,
	PNG: true,
	GIF: true,
}

var encodings = map[Encoding]bool{
	Native: true,
	Tar:    true,
	GZip:   true,
}

var classOfType = map[Type]Class{
	UnknownType: Unclassified,

	// Image formats.
	JPG: Image,
	PNG: Image,
	GIF: Image,
}
