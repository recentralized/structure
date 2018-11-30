package index

import (
	"reflect"
	"testing"

	"github.com/recentralized/structure/cid"
)

func TestAddRef(t *testing.T) {
	tests := []struct {
		desc  string
		idx   *Index
		ref   Ref
		added bool
		refs  int
	}{
		{
			desc: "empty index",
			idx:  &Index{},
			ref: Ref{
				Hash: cid.NewLiteral("a"),
				Src:  SrcItem{SrcID: SrcID("a")},
				Dst:  DstItem{DstID: DstID("a")},
			},
			added: true,
			refs:  1,
		},
		{
			desc: "new hash",
			idx: &Index{
				Refs: []*URef{
					{
						Hash: cid.NewLiteral("b"),
						Srcs: []SrcItem{SrcItem{SrcID: SrcID("a")}},
						Dsts: []DstItem{DstItem{DstID: DstID("a")}},
					},
				},
			},
			ref: Ref{
				Hash: cid.NewLiteral("a"),
				Src:  SrcItem{SrcID: SrcID("a")},
				Dst:  DstItem{DstID: DstID("a")},
			},
			refs:  2,
			added: true,
		},
		{
			desc: "matching hash with existing src, not dst",
			idx: &Index{
				Refs: []*URef{
					{
						Hash: cid.NewLiteral("a"),
						Srcs: []SrcItem{SrcItem{SrcID: SrcID("a")}},
						Dsts: []DstItem{DstItem{DstID: DstID("b")}},
					},
				},
			},
			ref: Ref{
				Hash: cid.NewLiteral("a"),
				Src:  SrcItem{SrcID: SrcID("a")},
				Dst:  DstItem{DstID: DstID("a")},
			},
			added: true,
			refs:  1,
		},
		{
			desc: "matching hash with existing dst, not src",
			idx: &Index{
				Refs: []*URef{
					{
						Hash: cid.NewLiteral("a"),
						Srcs: []SrcItem{SrcItem{SrcID: SrcID("b")}},
						Dsts: []DstItem{DstItem{DstID: DstID("a")}},
					},
				},
			},
			ref: Ref{
				Hash: cid.NewLiteral("a"),
				Src:  SrcItem{SrcID: SrcID("a")},
				Dst:  DstItem{DstID: DstID("a")},
			},
			added: true,
			refs:  1,
		},
		{
			desc: "matching hash with existing src and dst",
			idx: &Index{
				Refs: []*URef{
					{
						Hash: cid.NewLiteral("a"),
						Srcs: []SrcItem{SrcItem{SrcID: SrcID("a")}},
						Dsts: []DstItem{DstItem{DstID: DstID("a")}},
					},
				},
			},
			ref: Ref{
				Hash: cid.NewLiteral("a"),
				Src:  SrcItem{SrcID: SrcID("a")},
				Dst:  DstItem{DstID: DstID("a")},
			},
			refs:  1,
			added: false,
		},
	}
	for _, tt := range tests {
		got := tt.idx.AddRef(tt.ref)
		if got, want := got, tt.added; !reflect.DeepEqual(got, want) {
			t.Errorf("%q AddRef() got %t want %t", tt.desc, got, want)
		}
		if got, want := len(tt.idx.Refs), tt.refs; got != want {
			t.Errorf("%q len(refs) got %d want %d", tt.desc, got, want)
		}
	}
}
