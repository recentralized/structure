package index

// Index is the list of sources, destinations, and refs.
type Index struct {
	Srcs []Src   `json:"srcs,omitempty"`
	Dsts []Dst   `json:"dsts,omitempty"`
	Refs []*URef `json:"refs,omitempty"`
}
