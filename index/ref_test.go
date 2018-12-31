package index

import (
	"reflect"
	"testing"
	"time"

	"github.com/recentralized/structure/uri"
)

func TestURefAddSrc(t *testing.T) {
	tests := []struct {
		desc   string
		start  *URef
		add    SrcItem
		want   *URef
		update bool
	}{
		{
			desc:  "add first",
			start: &URef{},
			add:   SrcItem{SrcID: SrcID("a"), DataURI: uri.TrustedNew("a")},
			want: &URef{
				Srcs: []SrcItem{
					{SrcID: SrcID("a"), DataURI: uri.TrustedNew("a")},
				},
			},
			update: true,
		},
		{
			desc: "add duplicate is idempotent",
			start: &URef{
				Srcs: []SrcItem{
					{
						SrcID:   SrcID("a"),
						DataURI: uri.TrustedNew("a"),
						MetaURI: uri.TrustedNew("b"),
					},
				},
			},
			add: SrcItem{
				SrcID:   SrcID("a"),
				DataURI: uri.TrustedNew("a"),
				MetaURI: uri.TrustedNew("b"),
			},
			want: &URef{
				Srcs: []SrcItem{
					{
						SrcID:   SrcID("a"),
						DataURI: uri.TrustedNew("a"),
						MetaURI: uri.TrustedNew("b"),
					},
				},
			},
			update: false,
		},
		{
			desc: "add duplicate key updates mutlable attributes",
			start: &URef{
				Srcs: []SrcItem{
					{
						SrcID:      SrcID("a"),
						DataURI:    uri.TrustedNew("a"),
						MetaURI:    uri.TrustedNew("x"),
						ModifiedAt: time.Date(2, 3, 4, 5, 6, 7, 8, time.UTC),
					},
				},
			},
			add: SrcItem{
				SrcID:      SrcID("a"),
				DataURI:    uri.TrustedNew("a"),
				MetaURI:    uri.TrustedNew("x"),
				ModifiedAt: time.Date(9, 3, 4, 5, 6, 7, 8, time.UTC),
			},
			want: &URef{
				Srcs: []SrcItem{
					{
						SrcID:      SrcID("a"),
						DataURI:    uri.TrustedNew("a"),
						MetaURI:    uri.TrustedNew("x"),
						ModifiedAt: time.Date(9, 3, 4, 5, 6, 7, 8, time.UTC),
					},
				},
			},
			update: true,
		},
		{
			desc: "add new URI in the same source",
			start: &URef{
				Srcs: []SrcItem{
					{SrcID: SrcID("a"), DataURI: uri.TrustedNew("a")},
				},
			},
			add: SrcItem{SrcID: SrcID("a"), DataURI: uri.TrustedNew("b")},
			want: &URef{
				Srcs: []SrcItem{
					{SrcID: SrcID("a"), DataURI: uri.TrustedNew("a")},
					{SrcID: SrcID("a"), DataURI: uri.TrustedNew("b")},
				},
			},
			update: true,
		},
		{
			desc: "add same URI from a different source",
			start: &URef{
				Srcs: []SrcItem{
					{SrcID: SrcID("a"), DataURI: uri.TrustedNew("a")},
				},
			},
			add: SrcItem{SrcID: SrcID("b"), DataURI: uri.TrustedNew("a")},
			want: &URef{
				Srcs: []SrcItem{
					{SrcID: SrcID("a"), DataURI: uri.TrustedNew("a")},
					{SrcID: SrcID("b"), DataURI: uri.TrustedNew("a")},
				},
			},
			update: true,
		},
	}
	for _, tt := range tests {
		update := tt.start.AddSrc(tt.add)
		if !reflect.DeepEqual(tt.start, tt.want) {
			t.Errorf("%q AddSrc() got\n%#v\nwant\n%#v", tt.desc, tt.start, tt.want)
		}
		if update != tt.update {
			t.Errorf("%q AddSrc() update got %t, want %t", tt.desc, update, tt.update)
		}

	}
}

func TestURefAddDst(t *testing.T) {
	tests := []struct {
		desc   string
		start  *URef
		add    DstItem
		want   *URef
		update bool
	}{
		{
			desc:  "add first",
			start: &URef{},
			add:   DstItem{DstID: DstID("a"), DataURI: uri.TrustedNew("a")},
			want: &URef{
				Dsts: []DstItem{
					{DstID: DstID("a"), DataURI: uri.TrustedNew("a")},
				},
			},
			update: true,
		},
		{
			desc: "add duplicate is idempotent",
			start: &URef{
				Dsts: []DstItem{
					{
						DstID:     DstID("a"),
						DataURI:   uri.TrustedNew("a"),
						MetaURI:   uri.TrustedNew("a"),
						DataSize:  100,
						MetaSize:  10,
						StoredAt:  time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
						UpdatedAt: time.Date(2, 2, 3, 4, 5, 6, 7, time.UTC),
					},
				},
			},
			add: DstItem{
				DstID:     DstID("a"),
				DataURI:   uri.TrustedNew("a"),
				MetaURI:   uri.TrustedNew("a"),
				DataSize:  100,
				MetaSize:  10,
				StoredAt:  time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
				UpdatedAt: time.Date(2, 2, 3, 4, 5, 6, 7, time.UTC),
			},
			want: &URef{
				Dsts: []DstItem{
					{
						DstID:     DstID("a"),
						DataURI:   uri.TrustedNew("a"),
						MetaURI:   uri.TrustedNew("a"),
						DataSize:  100,
						MetaSize:  10,
						StoredAt:  time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
						UpdatedAt: time.Date(2, 2, 3, 4, 5, 6, 7, time.UTC),
					},
				},
			},
			update: false,
		},
		{
			desc: "add duplicate key updates mutable attributes",
			start: &URef{
				Dsts: []DstItem{
					{
						DstID:     DstID("a"),
						DataURI:   uri.TrustedNew("a"),
						MetaURI:   uri.TrustedNew("a"),
						DataSize:  100,
						MetaSize:  10,
						StoredAt:  time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
						UpdatedAt: time.Date(2, 2, 3, 4, 5, 6, 7, time.UTC),
					},
				},
			},
			add: DstItem{
				DstID:     DstID("a"),
				DataURI:   uri.TrustedNew("a"),
				MetaURI:   uri.TrustedNew("a"),
				DataSize:  102,
				MetaSize:  12,
				StoredAt:  time.Date(9, 2, 3, 4, 5, 6, 7, time.UTC),
				UpdatedAt: time.Date(8, 2, 3, 4, 5, 6, 7, time.UTC),
			},
			want: &URef{
				Dsts: []DstItem{
					{
						DstID:     DstID("a"),
						DataURI:   uri.TrustedNew("a"),
						MetaURI:   uri.TrustedNew("a"),
						DataSize:  102,
						MetaSize:  12,
						StoredAt:  time.Date(9, 2, 3, 4, 5, 6, 7, time.UTC),
						UpdatedAt: time.Date(8, 2, 3, 4, 5, 6, 7, time.UTC),
					},
				},
			},
			update: true,
		},
		{
			desc: "add same DstID with different URIs appends a new entry",
			start: &URef{
				Dsts: []DstItem{
					{
						DstID:    DstID("a"),
						DataURI:  uri.TrustedNew("a"),
						MetaURI:  uri.TrustedNew("a"),
						StoredAt: time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
					},
				},
			},
			add: DstItem{
				DstID:    DstID("a"),
				DataURI:  uri.TrustedNew("a"),
				MetaURI:  uri.TrustedNew("b"),
				StoredAt: time.Date(9, 2, 3, 4, 5, 6, 7, time.UTC),
			},
			want: &URef{
				Dsts: []DstItem{
					{
						DstID:    DstID("a"),
						DataURI:  uri.TrustedNew("a"),
						MetaURI:  uri.TrustedNew("a"),
						StoredAt: time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
					},
					{
						DstID:    DstID("a"),
						DataURI:  uri.TrustedNew("a"),
						MetaURI:  uri.TrustedNew("b"),
						StoredAt: time.Date(9, 2, 3, 4, 5, 6, 7, time.UTC),
					},
				},
			},
			update: true,
		},
	}
	for _, tt := range tests {
		update := tt.start.AddDst(tt.add)
		if !reflect.DeepEqual(tt.start, tt.want) {
			t.Errorf("%q AddDst() got\n%#v\nwant\n%#v", tt.desc, tt.start, tt.want)
		}
		if update != tt.update {
			t.Errorf("%q AddDst() update got %t, want %t", tt.desc, update, tt.update)
		}
	}
}
