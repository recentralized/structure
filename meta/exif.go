package meta

// Exif is all of the Exif data.
type Exif map[string]ExifValue

// ExifValue is an individual Exif value.
type ExifValue struct {
	ID  string      `json:"id"`
	Val interface{} `json:"val"`
}
