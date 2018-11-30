package index

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/recentralized/structure/uri"
)

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
			json: `{"dst_id":"","data_uri":"","meta_uri":"","stored_at":null}`,
		},
		{
			desc: "basic fields",
			item: DstItem{
				DstID:    DstID("abc"),
				DataURI:  uri.TrustedNew("http://example.com/data/abc.jpg"),
				MetaURI:  uri.TrustedNew("http://example.com/meta/abc.json"),
				StoredAt: time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
			},
			json: `{"dst_id":"abc","data_uri":"http://example.com/data/abc.jpg","meta_uri":"http://example.com/meta/abc.json","stored_at":"0001-02-03T04:05:06.000000007Z"}`,
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
