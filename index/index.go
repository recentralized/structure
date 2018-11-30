package index

// Index is the list of sources, destinations, and refs.
type Index struct {
	Srcs []Src   `json:"srcs,omitempty"`
	Dsts []Dst   `json:"dsts,omitempty"`
	Refs []*URef `json:"refs,omitempty"`
}

// AddRef implements the logic to add a Ref to the Index. It's idempotent; if
// either the Ref's Src or Dst was actually added, the method returns true.
// Storage implementations that write directly to a database should probably
// implement this logic directly in the db.
func (i *Index) AddRef(ref Ref) bool {
	var uref *URef
	for _, u := range i.Refs {
		if u.Hash.EqualHash(ref.Hash) {
			uref = u
			break
		}
	}
	if uref == nil {
		uref = &URef{Hash: ref.Hash}
		i.Refs = append(i.Refs, uref)
	}
	addSrc := uref.AddSrc(ref.Src)
	addDst := uref.AddDst(ref.Dst)
	return addSrc || addDst
}
