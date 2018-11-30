package structure

import (
	"fmt"

	"github.com/recentralized/structure/cid"
)

// Ref is a distinct piece of content (represented by Hash), where it was found
// (SrcItem) and where it was stored (DstItem). Refs are generally created
// while searching a Src and copying to a Dst.
type Ref struct {
	Hash cid.ContentID
	Src  SrcItem
	Dst  DstItem
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
type URef struct {
	Hash cid.ContentID
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

// AddSrc implements the logic to add a new instance of a source item. Source
// items are unique for each Source and Data URI. Storage implementations that
// write directly to a database should probably implement this logic directly
// in the db.
func (r *URef) AddSrc(src SrcItem) bool {
	for i, s := range r.Srcs {
		if s.EqualKey(src) {
			if s.Equal(src) {
				return false
			}
			r.Srcs[i] = src
			return true
		}
	}
	r.Srcs = append(r.Srcs, src)
	return true
}

// AddDst implements the logic to add a new instance of a destinatino item.
// Destination items are unique per Destination. Storage implementations that
// write directly to a database should probably implement this logic directly
// in the db.
func (r *URef) AddDst(dst DstItem) bool {
	for _, d := range r.Dsts {
		if d.EqualKey(dst) {
			return false
		}
	}
	r.Dsts = append(r.Dsts, dst)
	return true
}

// GetSrc returns the Src matching srcID. If no match is found, the boolean is
// false.
func (r *URef) GetSrc(srcID SrcID) (SrcItem, bool) {
	for _, item := range r.Srcs {
		if item.SrcID == srcID {
			return item, true
		}
	}
	return SrcItem{}, false
}

// GetDst returns the Dst matching dstID. If no match is found, the boolean is
// false.
func (r *URef) GetDst(dstID DstID) (DstItem, bool) {
	for _, item := range r.Dsts {
		if item.DstID == dstID {
			return item, true
		}
	}
	return DstItem{}, false
}
