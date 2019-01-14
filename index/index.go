package index

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/recentralized/structure/data"
)

// Version identifies the version of the index structure. A new version will be
// introduced for backward-incompatible changes.
const Version = versionV1

const (
	// versionv0 is the original implementation.
	versionV0 = ""

	// versionv1 is the first `structure` implementation.
	versionV1 = "v1"
)

// ErrWrongVersion means that the parsed index is not at the current version,
// so its data may be incorrectly interpreted.
var ErrWrongVersion = errors.New("index is not at a compatible version")

// Index is the set of sources, destinations, and refs.
//
// The methods it provides are convenience, suitable for small in-memory
// implementations. Since index could be implemented in any number of ways,
// such as a relational dataase, the methods here serve as documentation of the
// algorithms to add and retrieve data.
type Index struct {
	Version string  `json:"version"`
	Srcs    []Src   `json:"srcs,omitempty"`
	Dsts    []Dst   `json:"dsts,omitempty"`
	Refs    []*URef `json:"refs,omitempty"`
}

// New initializes a new Index at the current version.
func New() *Index {
	return &Index{Version: Version}
}

// ParseJSON loads an Index from JSON. If the loaded data cannot be
// transparently upgraded to the current version then ErrWrongVersion is
// returned.
func ParseJSON(r io.Reader) (*Index, error) {
	idx := &Index{}
	err := json.NewDecoder(r).Decode(idx)
	if err != nil {
		return nil, err
	}
	switch idx.Version {
	case versionV1:
	case versionV0:
		idx.Version = versionV1
	default:
		return nil, ErrWrongVersion
	}
	return idx, nil
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
		if u.Hash.Equal(ref.Hash) {
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
func (i *Index) GetRef(hash data.Hash) (*URef, bool) {
	for _, uref := range i.Refs {
		if uref.Hash.Equal(hash) {
			return uref, true
		}
	}
	return nil, false
}
