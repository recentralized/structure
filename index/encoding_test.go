package index

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/recentralized/structure/data"
	"github.com/recentralized/structure/uri"
)

func TestSrcJSON(t *testing.T) {
	tests := []struct {
		desc string
		src  Src
		json string
	}{

		{
			desc: "zero value",
			src:  Src{},
			json: `{"src_id":"","src_uri":""}`,
		},
		{
			desc: "basic fields",
			src: Src{
				SrcID:  SrcID("a"),
				SrcURI: uri.TrustedNew("http://example.com"),
			},
			json: `{"src_id":"a","src_uri":"http://example.com"}`,
		},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.src)
		if err != nil {
			t.Fatalf("%q Src failed to marshal: %s", tt.desc, err)
		}
		log.Printf("%q Src JSON %s", tt.desc, data)
		if string(data) != tt.json {
			t.Errorf("%q Src JSON\ngot  %s\nwant %s", tt.desc, data, tt.json)
		}
		got := Src{}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Fatalf("%q Src failed to unmarshal: %s", tt.desc, err)
		}
		if !reflect.DeepEqual(tt.src, got) {
			t.Errorf("%q Src Roundtrip\ngot  %#v\nwant %#v", tt.desc, got, tt.src)
		}
	}
}

func TestSrcItemJSON(t *testing.T) {
	tests := []struct {
		desc string
		item SrcItem
		json string
	}{

		{
			desc: "zero value",
			item: SrcItem{},
			json: `{"src_id":"","data_uri":"","meta_uri":"","modified_at":null}`,
		},
		{
			desc: "basic fields",
			item: SrcItem{
				SrcID:      SrcID("a"),
				DataURI:    uri.TrustedNew("http://example.com/data.jpg"),
				MetaURI:    uri.TrustedNew("http://example.com/meta.json"),
				ModifiedAt: time.Date(2015, 2, 3, 4, 5, 6, 7, time.UTC),
			},
			json: `{"src_id":"a","data_uri":"http://example.com/data.jpg","meta_uri":"http://example.com/meta.json","modified_at":"2015-02-03T04:05:06.000000007Z"}`,
		},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.item)
		if err != nil {
			t.Fatalf("%q SrcItem failed to marshal: %s", tt.desc, err)
		}
		log.Printf("%q SrcItem JSON %s", tt.desc, data)
		if string(data) != tt.json {
			t.Errorf("%q SrcItem JSON\ngot  %s\nwant %s", tt.desc, data, tt.json)
		}
		got := SrcItem{}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Fatalf("%q SrcItem failed to unmarshal: %s", tt.desc, err)
		}
		if !reflect.DeepEqual(tt.item, got) {
			t.Errorf("%q SrcItem Roundtrip\ngot  %#v\nwant %#v", tt.desc, got, tt.item)
		}
	}
}
func TestDstJSON(t *testing.T) {
	tests := []struct {
		desc string
		dst  Dst
		json string
	}{
		{
			desc: "zero value",
			dst:  Dst{},
			json: `{"dst_id":"","index_uri":"","data_uri":"","meta_uri":""}`,
		},
		{
			desc: "basic fields",
			dst: Dst{
				DstID:    DstID("abc"),
				IndexURI: uri.TrustedNew("http://example.com/"),
				DataURI:  uri.TrustedNew("http://example.com/data"),
				MetaURI:  uri.TrustedNew("http://example.com/meta"),
			},
			json: `{"dst_id":"abc","index_uri":"http://example.com/","data_uri":"http://example.com/data","meta_uri":"http://example.com/meta"}`,
		},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.dst)
		if err != nil {
			t.Fatalf("%q Dst failed to marshal: %s", tt.desc, err)
		}
		if string(data) != tt.json {
			t.Errorf("%q Dst JSON\ngot  %s\nwant %s", tt.desc, data, tt.json)
		}
		log.Printf("%q Dst JSON %s", tt.desc, data)
		got := Dst{}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Fatalf("%q Dst failed to unmarshal: %s", tt.desc, err)
		}
		if !reflect.DeepEqual(tt.dst, got) {
			t.Errorf("%q Dst Roundtrip\ngot  %#v\nwant %#v", tt.desc, got, tt.dst)
		}
	}
}

func TestDstItemJSON(t *testing.T) {
	tests := []struct {
		desc string
		item DstItem
		json string
	}{

		{
			desc: "zero value",
			item: DstItem{},
			json: `{"dst_id":"","data_uri":"","meta_uri":"","size":0,"stored_at":null}`,
		},
		{
			desc: "basic fields",
			item: DstItem{
				DstID:    DstID("abc"),
				DataURI:  uri.TrustedNew("http://example.com/data/abc.jpg"),
				MetaURI:  uri.TrustedNew("http://example.com/meta/abc.json"),
				Size:     100,
				StoredAt: time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
			},
			json: `{"dst_id":"abc","data_uri":"http://example.com/data/abc.jpg","meta_uri":"http://example.com/meta/abc.json","size":100,"stored_at":"0001-02-03T04:05:06.000000007Z"}`,
		},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.item)
		if err != nil {
			t.Fatalf("%q DstItem failed to marshal: %s", tt.desc, err)
		}
		if string(data) != tt.json {
			t.Errorf("%q DstItem JSON\ngot  %s\nwant %s", tt.desc, data, tt.json)
		}
		log.Printf("%q DstItem JSON %s", tt.desc, data)
		got := DstItem{}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Fatalf("%q DstItem failed to unmarshal: %s", tt.desc, err)
		}
		if !reflect.DeepEqual(tt.item, got) {
			t.Errorf("%q DstItem Roundtrip\ngot  %#v\nwant %#v", tt.desc, got, tt.item)
		}
	}
}

func TestURefJSON(t *testing.T) {
	tests := []struct {
		desc string
		ref  URef
		json string
	}{

		{
			desc: "zero value",
			ref:  URef{},
			json: `{"hash":"","srcs":[],"dsts":[]}`,
		},
		{
			desc: "all fields",
			ref: URef{
				Hash: data.LiteralHash("xyz"),
				Srcs: []SrcItem{
					{
						SrcID:      SrcID("a"),
						DataURI:    uri.TrustedNew("http://example.com/data.jpg"),
						MetaURI:    uri.TrustedNew("http://example.com/meta.json"),
						ModifiedAt: time.Date(2015, 2, 3, 4, 5, 6, 7, time.UTC),
					},
				},
				Dsts: []DstItem{
					{
						DstID:    DstID("abc"),
						DataURI:  uri.TrustedNew("http://example.com/data/abc.jpg"),
						MetaURI:  uri.TrustedNew("http://example.com/meta/abc.json"),
						Size:     100,
						StoredAt: time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
					},
				},
			},
			json: `{"hash":"xyz","srcs":[{"src_id":"a","data_uri":"http://example.com/data.jpg","meta_uri":"http://example.com/meta.json","modified_at":"2015-02-03T04:05:06.000000007Z"}],"dsts":[{"dst_id":"abc","data_uri":"http://example.com/data/abc.jpg","meta_uri":"http://example.com/meta/abc.json","size":100,"stored_at":"0001-02-03T04:05:06.000000007Z"}]}`,
		},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.ref)
		if err != nil {
			t.Fatalf("%q URef failed to marshal: %s", tt.desc, err)
		}
		if string(data) != tt.json {
			t.Errorf("%q URef JSON\ngot  %s\nwant %s", tt.desc, data, tt.json)
		}
		log.Printf("%q URef JSON %s", tt.desc, data)
		got := URef{}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Fatalf("%q URef failed to unmarshal: %s", tt.desc, err)
		}
		if !reflect.DeepEqual(tt.ref, got) {
			t.Errorf("%q URef Roundtrip\ngot  %#v\nwant %#v", tt.desc, got, tt.ref)
		}
	}
}

func TestIndexJSON(t *testing.T) {
	tests := []struct {
		desc string
		idx  Index
		json string
	}{

		{
			desc: "zero value",
			idx:  Index{},
			json: `{"version":""}`,
		},
		{
			desc: "srcs + dsts",
			idx: Index{
				Version: "v1",
				Srcs: []Src{
					{SrcID: SrcID("a")},
				},
				Dsts: []Dst{
					{DstID: DstID("abc")},
				},
			},
			json: `{"version":"v1","srcs":[{"src_id":"a","src_uri":""}],"dsts":[{"dst_id":"abc","index_uri":"","data_uri":"","meta_uri":""}]}`,
		},
		{
			desc: "refs",
			idx: Index{
				Version: "v1",
				Refs: []*URef{
					{
						Hash: data.LiteralHash("xyz"),
						Srcs: []SrcItem{
							{SrcID: SrcID("a")},
						},
						Dsts: []DstItem{
							{DstID: DstID("abc")},
						},
					},
				},
			},
			json: `{"version":"v1","refs":[{"hash":"xyz","srcs":[{"src_id":"a","data_uri":"","meta_uri":"","modified_at":null}],"dsts":[{"dst_id":"abc","data_uri":"","meta_uri":"","size":0,"stored_at":null}]}]}`,
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
