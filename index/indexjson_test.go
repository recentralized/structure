package index

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"

	"github.com/recentralized/structure/cid"
)

func TestIndexJSON(t *testing.T) {
	tests := []struct {
		desc string
		idx  Index
		json string
	}{

		{
			desc: "zero value",
			idx:  Index{},
			json: `{}`,
		},
		{
			desc: "srcs + dsts",
			idx: Index{
				Srcs: []Src{
					{SrcID: SrcID("a")},
				},
				Dsts: []Dst{
					{DstID: DstID("abc")},
				},
			},
			json: `{"srcs":[{"src_id":"a","src_uri":""}],"dsts":[{"dst_id":"abc","index_uri":"","data_uri":"","meta_uri":""}]}`,
		},
		{
			desc: "refs",
			idx: Index{
				Refs: []*URef{
					{
						Hash: cid.NewLiteral("xyz"),
						Srcs: []SrcItem{
							{SrcID: SrcID("a")},
						},
						Dsts: []DstItem{
							{DstID: DstID("abc")},
						},
					},
				},
			},
			json: `{"refs":[{"hash":"xyz","srcs":[{"src_id":"a","data_uri":"","meta_uri":"","modified_at":null}],"dsts":[{"dst_id":"abc","data_uri":"","meta_uri":"","stored_at":null}]}]}`,
		},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.idx)
		if err != nil {
			t.Fatalf("%q Index failed to marshal: %s", tt.desc, err)
		}
		log.Printf("%q Src JSON %s", tt.desc, data)
		if string(data) != tt.json {
			t.Errorf("%q Index JSON\ngot  %s\nwant %s", tt.desc, data, tt.json)
		}
		got := Index{}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Fatalf("%q Index failed to unmarshal: %s", tt.desc, err)
		}
		if !reflect.DeepEqual(tt.idx, got) {
			t.Errorf("%q Index Roundtrip\ngot  %#v\nwant %#v", tt.desc, got, tt.idx)
		}
	}
}
