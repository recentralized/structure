package structure

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"time"

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
