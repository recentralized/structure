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
	Native Encoding = ""
	Tar             = "tar"
	GZip            = "gz"
)

// Class definitions.
const (
	Uncategorized Class = ""
	Image               = "image"
)

// Type is a known kind of file such as JPEG or PNG.
type Type string

func (t Type) String() string {
	return string(t)
}

// Ext returns the type's standard file extension.
func (t Type) Ext() string {
	return string(t)
}

// Class returns the type's class: image, catalog, etc.
func (t Type) Class() Class {
	return classOfType[t]
}

// ParseExt parses an extension, returning the Stored format.
func ParseExt(str string) (Stored, error) {
	parts := strings.Split(str, ".")
	if len(parts) > 2 {
		return Stored{}, fmt.Errorf("data: too many parts in extension %q", str)
	}
	t := Type(parts[0])
	e := Native
	if len(parts) == 2 {
		e = Encoding(parts[1])
	}
	return Stored{t, e}, nil
}

// Stored is how a type is formatted for storage.
type Stored struct {
	Type     Type
	Encoding Encoding
}

func (s Stored) String() string {
	return s.Ext()
}

// Ext returns the stored data's standard file extension.
func (s Stored) Ext() string {
	if s.Encoding == Native {
		return fmt.Sprintf("%s", s.Type.Ext())
	}
	return fmt.Sprintf("%s.%s", s.Type.Ext(), s.Encoding.Ext())
}

// Encoding is the encoding of the data for storage. Most types are stored in
// their native encoding, but we may want to optimize storage by compressing or
// flattening multi-file structures.
type Encoding string

func (e Encoding) String() string {
	if e == "" {
		return "<native>"
	}
	return string(e)
}

// Ext returns the encoding's standard file extension.
func (e Encoding) Ext() string {
	return string(e)
}

// Class is the category of data that a type belongs to. JPG, PNG, GIF are all
// image, etc.
type Class string

func (c Class) String() string {
	return string(c)
}

var classOfType = map[Type]Class{
	UnknownType: Uncategorized,

	// Image formats.
	JPG: Image,
	PNG: Image,
	GIF: Image,
}
