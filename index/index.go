package index

import "github.com/recentralized/structure/cid"

// Index is the set of sources, destinations, and refs.
//
// The methods it provides are convenience, suitable for small in-memory
// implementations. Since index could be implemented in any number of ways,
// such as a relational dataase, the methods here serve as documentation of the
// algorithms to add and retrieve data.
type Index struct {
	Srcs []Src   `json:"srcs,omitempty"`
	Dsts []Dst   `json:"dsts,omitempty"`
	Refs []*URef `json:"refs,omitempty"`
}

// AddSrc adds a source to the index. It's idempotent, returning true if the
// index was modified.
func (i *Index) AddSrc(src Src) bool {
	for _, s := range i.Srcs {
		if s.SrcID == src.SrcID {
			return false
		}
	}
	i.Srcs = append(i.Srcs, src)
	return true
}

// AddDst adds a destination to the index. It's idempotent, returning true if
// the index was modified.
func (i *Index) AddDst(dst Dst) bool {
	for _, d := range i.Dsts {
		if d.DstID == dst.DstID {
			return false
		}
	}
	i.Dsts = append(i.Dsts, dst)
	return true
}

// GetSrc returns the source with srcID. It returns false if no source was
// found.
func (i *Index) GetSrc(srcID SrcID) (Src, bool) {
	for _, src := range i.Srcs {
		if src.SrcID == srcID {
			return src, true
		}
	}
	return Src{}, false
}

// GetDst returns the destination with dstID. It returns false if no
// destination was found.
func (i *Index) GetDst(dstID DstID) (Dst, bool) {
	for _, dst := range i.Dsts {
		if dst.DstID == dstID {
			return dst, true
		}
	}
	return Dst{}, false
}

// AddRef adds a ref to the index. A ref is a hash with source and destination.
// It's idempotent, returning true if the index was modified.
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

// GetRef retrieves a URef from the index. A URef is a hash with all sources
// and destinations that have been added. If you're only interested in one
// source or destination see GetSrcItem and GetDstItem.
func (i *Index) GetRef(hash cid.ContentID) (*URef, bool) {
	for _, uref := range i.Refs {
		if uref.Hash.EqualHash(hash) {
			return uref, true
		}
	}
	return nil, false
}

// GetSrcItem returns a SrcItem for content with hash and stored in srcID. If
// no SrcItem is found, it returns false.
func (i *Index) GetSrcItem(hash cid.ContentID, srcID SrcID) (SrcItem, bool) {
	ref, ok := i.GetRef(hash)
	if ok {
		return ref.GetSrc(srcID)
	}
	return SrcItem{}, false
}

// GetDstItem returns a DstItem for content with hash and stored in dstID. If
// no DstItem is found, it returns false.
func (i *Index) GetDstItem(hash cid.ContentID, dstID DstID) (DstItem, bool) {
	ref, ok := i.GetRef(hash)
	if ok {
		return ref.GetDst(dstID)
	}
	return DstItem{}, false
}
