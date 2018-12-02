package index

import (
	"testing"

	"github.com/recentralized/structure/data"
)

func TestIndexSrc(t *testing.T) {
	tests := []struct {
		desc string
		idx  *Index
		src  Src
		get  bool
		add  bool
		srcs int
	}{
		{
			desc: "empty index",
			idx:  &Index{},
			src:  Src{SrcID: SrcID("a")},
			get:  false,
			add:  true,
			srcs: 1,
		},
		{
			desc: "existing src",
			idx: &Index{
				Srcs: []Src{
					{SrcID: SrcID("a")},
				},
			},
			src:  Src{SrcID: SrcID("a")},
			get:  true,
			add:  false,
			srcs: 1,
		},
		{
			desc: "new src",
			idx: &Index{
				Srcs: []Src{
					{SrcID: SrcID("a")},
				},
			},
			src:  Src{SrcID: SrcID("b")},
			get:  false,
			add:  true,
			srcs: 2,
		},
	}
	for _, tt := range tests {
		var (
			src Src
			ok  bool
			add bool
		)

		src, ok = tt.idx.GetSrc(tt.src.SrcID)
		if got, want := ok, tt.get; got != want {
			t.Errorf("%q initial GetSrc() got %t want %t", tt.desc, got, want)
		}
		if ok {
			if got, want := src.SrcID, tt.src.SrcID; got != want {
				t.Errorf("%q GetSrc() got %s want %s", tt.desc, got, want)
			}
		}

		add = tt.idx.AddSrc(tt.src)
		if got, want := add, tt.add; got != want {
			t.Errorf("%q AddSrc() got %t want %t", tt.desc, got, want)
		}
		if got, want := len(tt.idx.Srcs), tt.srcs; got != want {
			t.Errorf("%q len(srcs) got %d want %d", tt.desc, got, want)
		}

		_, ok = tt.idx.GetSrc(tt.src.SrcID)
		if !ok {
			t.Errorf("%q GetSrc must be ok after adding", tt.desc)
		}
	}
}

func TestIndexDst(t *testing.T) {
	tests := []struct {
		desc string
		idx  *Index
		dst  Dst
		get  bool
		add  bool
		dsts int
	}{
		{
			desc: "empty index",
			idx:  &Index{},
			dst:  Dst{DstID: DstID("a")},
			get:  false,
			add:  true,
			dsts: 1,
		},
		{
			desc: "existing dst",
			idx: &Index{
				Dsts: []Dst{
					{DstID: DstID("a")},
				},
			},
			dst:  Dst{DstID: DstID("a")},
			get:  true,
			add:  false,
			dsts: 1,
		},
		{
			desc: "new dst",
			idx: &Index{
				Dsts: []Dst{
					{DstID: DstID("a")},
				},
			},
			dst:  Dst{DstID: DstID("b")},
			get:  false,
			add:  true,
			dsts: 2,
		},
	}
	for _, tt := range tests {
		var (
			dst Dst
			ok  bool
			add bool
		)

		dst, ok = tt.idx.GetDst(tt.dst.DstID)
		if got, want := ok, tt.get; got != want {
			t.Errorf("%q initial GetDst() got %t want %t", tt.desc, got, want)
		}
		if ok {
			if got, want := dst.DstID, tt.dst.DstID; got != want {
				t.Errorf("%q GetDst() got %s want %s", tt.desc, got, want)
			}
		}

		add = tt.idx.AddDst(tt.dst)
		if got, want := add, tt.add; got != want {
			t.Errorf("%q AddDst() got %t want %t", tt.desc, got, want)
		}
		if got, want := len(tt.idx.Dsts), tt.dsts; got != want {
			t.Errorf("%q len(dsts) got %d want %d", tt.desc, got, want)
		}

		_, ok = tt.idx.GetDst(tt.dst.DstID)
		if !ok {
			t.Errorf("%q GetDst must be ok after adding", tt.desc)
		}
	}
}

func TestIndexRef(t *testing.T) {
	tests := []struct {
		desc string
		idx  *Index
		ref  Ref
		get  bool
		add  bool
		refs int
	}{
		{
			desc: "empty index",
			idx:  &Index{},
			ref: Ref{
				Hash: data.LiteralHash("a"),
				Src:  SrcItem{SrcID: SrcID("a")},
				Dst:  DstItem{DstID: DstID("a")},
			},
			get:  false,
			add:  true,
			refs: 1,
		},
		{
			desc: "new hash",
			idx: &Index{
				Refs: []*URef{
					{
						Hash: data.LiteralHash("b"),
						Srcs: []SrcItem{{SrcID: SrcID("a")}},
						Dsts: []DstItem{{DstID: DstID("a")}},
					},
				},
			},
			ref: Ref{
				Hash: data.LiteralHash("a"),
				Src:  SrcItem{SrcID: SrcID("a")},
				Dst:  DstItem{DstID: DstID("a")},
			},
			get:  false,
			add:  true,
			refs: 2,
		},
		{
			desc: "matching hash with existing src, not dst",
			idx: &Index{
				Refs: []*URef{
					{
						Hash: data.LiteralHash("a"),
						Srcs: []SrcItem{{SrcID: SrcID("a")}},
						Dsts: []DstItem{{DstID: DstID("b")}},
					},
				},
			},
			ref: Ref{
				Hash: data.LiteralHash("a"),
				Src:  SrcItem{SrcID: SrcID("a")},
				Dst:  DstItem{DstID: DstID("a")},
			},
			get:  true,
			add:  true,
			refs: 1,
		},
		{
			desc: "matching hash with existing dst, not src",
			idx: &Index{
				Refs: []*URef{
					{
						Hash: data.LiteralHash("a"),
						Srcs: []SrcItem{{SrcID: SrcID("b")}},
						Dsts: []DstItem{{DstID: DstID("a")}},
					},
				},
			},
			ref: Ref{
				Hash: data.LiteralHash("a"),
				Src:  SrcItem{SrcID: SrcID("a")},
				Dst:  DstItem{DstID: DstID("a")},
			},
			get:  true,
			add:  true,
			refs: 1,
		},
		{
			desc: "matching hash with existing src and dst",
			idx: &Index{
				Refs: []*URef{
					{
						Hash: data.LiteralHash("a"),
						Srcs: []SrcItem{{SrcID: SrcID("a")}},
						Dsts: []DstItem{{DstID: DstID("a")}},
					},
				},
			},
			ref: Ref{
				Hash: data.LiteralHash("a"),
				Src:  SrcItem{SrcID: SrcID("a")},
				Dst:  DstItem{DstID: DstID("a")},
			},
			get:  true,
			add:  false,
			refs: 1,
		},
	}
	for _, tt := range tests {
		var (
			ref *URef
			ok  bool
			add bool
		)

		ref, ok = tt.idx.GetRef(tt.ref.Hash)
		if got, want := ok, tt.get; got != want {
			t.Errorf("%q initial GetRef() got %t want %t", tt.desc, got, want)
		}
		if ok {
			if got, want := ref.Hash, tt.ref.Hash; !got.Equal(want) {
				t.Errorf("%q GetRef() got %s want %s", tt.desc, got, want)
			}
		}

		add = tt.idx.AddRef(tt.ref)
		if got, want := add, tt.add; got != want {
			t.Errorf("%q AddRef() got %t want %t", tt.desc, got, want)
		}
		if got, want := len(tt.idx.Refs), tt.refs; got != want {
			t.Errorf("%q len(refs) got %d want %d", tt.desc, got, want)
		}

		_, ok = tt.idx.GetRef(tt.ref.Hash)
		if !ok {
			t.Errorf("%q GetRef must be ok after adding", tt.desc)
		}
	}
}
