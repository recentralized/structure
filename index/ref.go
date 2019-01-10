package index

import (
	"fmt"

	"github.com/recentralized/structure/data"
)

// Ref is a distinct piece of content (represented by Hash), where it was found
// (SrcItem) and where it was stored (DstItem). Refs are generally created
// while searching a Src and copying to a Dst.
type Ref struct {
	// Hash is a fingerprint of the data. The fingerprint is of the data
	// itself, and may differ from the data stored, for example if
	// compressed for storage.
	Hash data.Hash

	// Src is details about where the content was found.
	Src SrcItem

	// Dst is details about whree the content was stored.
	Dst DstItem
}

func (r Ref) String() string {
	return fmt.Sprintf(
		"<Ref %s srcID:%q dstID:%q srcData:%q srcMeta:%q dstData:%q dstMeta:%q modified:%s stored:%s>",
		r.Hash,
		r.Src.SrcID,
		r.Dst.DstID,
		r.Src.DataURI,
		r.Src.MetaURI,
		r.Dst.DataURI,
		r.Dst.MetaURI,
		r.Src.ModifiedAt.String(),
		r.Dst.StoredAt.String())
}

// URef is the universal ref - all of the sources and destinations that a
// distinct piece of content has been found and put. Any number of Ref can be
// combined into a URef.
//
// The methods it provides are convenience, suitable for small in-memory
// implementations. Since refs could be implemented in any number of ways, such
// as a relational database, the methods here serve as documentation of the
// algorithms to add and retrieve data.
type URef struct {
	Hash data.Hash
	Srcs []SrcItem
	Dsts []DstItem
}

// NewURefFromRef converts a Ref to a URef.
func NewURefFromRef(ref Ref) *URef {
	r := &URef{Hash: ref.Hash}
	r.AddSrc(ref.Src)
	r.AddDst(ref.Dst)
	return r
}

func (r URef) String() string {
	return fmt.Sprintf("<URef %s srcs:%d dsts:%d>", r.Hash, len(r.Srcs), len(r.Dsts))
}

// DecomposeRefs returns an array of individual Refs.
func (r URef) DecomposeRefs() []Ref {
	refs := make([]Ref, 0, len(r.Srcs)*len(r.Dsts))
	for _, s := range r.Srcs {
		for _, d := range r.Dsts {
			refs = append(refs, Ref{
				Hash: r.Hash,
				Src:  s,
				Dst:  d,
			})
		}
	}
	return refs
}

// AddSrc adds a SrcItem to the ref. If a matching SrcItem exists, it's mutable
// attributes will be updated. The method returns true if any changes to the
// URef or existing SrcItem occurred.
func (r *URef) AddSrc(src SrcItem) bool {
	for i, s := range r.Srcs {
		if s.EqualKey(src) {
			if s.Equal(src) {
				return false
			}
			// Update mutable attributes.
			r.Srcs[i] = src
			return true
		}
	}
	r.Srcs = append(r.Srcs, src)
	return true
}

// AddDst adds a DstItem to the ref. If a matching DstItem exists, it's mutable
// attributes will be updated. The method returns true if any changes to the
// URef or existing DstItem occurred.
func (r *URef) AddDst(dst DstItem) bool {
	for i, d := range r.Dsts {
		if d.EqualKey(dst) {
			if d.Equal(dst) {
				return false
			}
			// Update mutable attributes.
			r.Dsts[i] = dst
			return true
		}
	}
	r.Dsts = append(r.Dsts, dst)
	return true
}

// GetSrc returns the source item matching srcID. It return false if no source
// item was found.
func (r *URef) GetSrc(srcID SrcID) (SrcItem, bool) {
	for _, item := range r.Srcs {
		if item.SrcID == srcID {
			return item, true
		}
	}
	return SrcItem{}, false
}

// GetDst returns the destination item matching dstID. It return false if no
// destination item was found.
func (r *URef) GetDst(dstID DstID) (DstItem, bool) {
	for _, item := range r.Dsts {
		if item.DstID == dstID {
			return item, true
		}
	}
	return DstItem{}, false
}
