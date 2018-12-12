package data

import (
	"errors"
	"fmt"
)

var (
	// ErrUnknownType is returned if the type cannot be determined.
	ErrUnknownType = errors.New("data: unknown type")
)

// Type is a known type of file such as JPEG or PNG. Only known types of files
// are copied from source to destination.
type Type string

func (t Type) String() string {
	return string(t)
}

// Ext returns the type's file extension.
func (t Type) Ext() string {
	enc := t.Enc()
	if enc == Native {
		return string(t)
	}
	return fmt.Sprintf("%s.%s", t, enc)
}

// Enc returns the type's encoding.
func (t Type) Enc() Encoding {
	return encodingType[t]
}

// Class returns the type's class - image, catalog, etc.
func (t Type) Class() Class {
	return typeClass[t]
}

// Type definitions.
const (
	// UnknownType is the zero value for Type, meaning it is unknown.
	UnknownType Type = ""

	// Image formats.
	JPG = "jpg" // Standard JPG file.
	PNG = "png" // Standard PNG file.
)

// Encoding is the encoding done to the content for storage. Most types are
// stored in their native encoding, but for other types we may want to optimize
// storage by compressing or flattening multi-file structures.
type Encoding string

func (e Encoding) String() string {
	return string(e)
}

// Encoding definitions.
const (
	Native Encoding = ""
	Tar             = "tar"
)

var encodingType = map[Type]Encoding{}

// Class is the category of content types that a specific type belongs to.
// JPG, PNG, GIF are all image, etc.
type Class string

func (c Class) String() string {
	return string(c)
}

// Class definitions.
const (
	Uncategorized Class = ""
	Image               = "image"
)

var typeClass = map[Type]Class{
	UnknownType: Uncategorized,

	// Image formats.
	JPG: Image,
	PNG: Image,
}
