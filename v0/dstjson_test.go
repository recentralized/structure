package structure

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
		dst  Dst
		json string
	}{

		// Zero value.
		{
			dst:  Dst{},
			json: `{"dst_id":"","index_uri":"","data_uri":"","meta_uri":""}`,
		},
		// Basic fields.
		{
			dst: Dst{
				DstID:    DstID("abc"),
				IndexURI: uri.TrustedNew("http://example.com/"),
				DataURI:  uri.TrustedNew("http://example.com/data"),
				MetaURI:  uri.TrustedNew("http://example.com/meta"),
			},
			json: `{"dst_id":"abc","index_uri":"http://example.com/","data_uri":"http://example.com/data","meta_uri":"http://example.com/meta"}`,
		},
	}
	for i, test := range tests {
		data, err := json.Marshal(test.dst)
		if err != nil {
			t.Fatalf("%d Dst failed to marshal: %s", i, err)
		}
		if test.json != "" {
			if string(data) != test.json {
				t.Errorf("%d Dst JSON\ngot  %s\nwant %s", i, data, test.json)
			}
		}
		log.Printf("%d Dst JSON %s", i, data)
		got := Dst{}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Fatalf("%d Dst failed to unmarshal: %s", i, err)
		}
		if !reflect.DeepEqual(test.dst, got) {
			t.Errorf("%d Dst Roundtrip\ngot  %#v\nwant %#v", i, got, test.dst)
		}
	}
}

func TestDstItemJSON(t *testing.T) {
	tests := []struct {
		item DstItem
		json string
	}{

		// Zero value.
		{
			item: DstItem{},
			json: `{"dst_id":"","data_uri":"","meta_uri":"","stored_at":null}`,
		},
		// Basic fields.
		{
			item: DstItem{
				DstID:    DstID("abc"),
				DataURI:  uri.TrustedNew("http://example.com/data/abc.jpg"),
				MetaURI:  uri.TrustedNew("http://example.com/meta/abc.json"),
				StoredAt: time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
			},
			json: `{"dst_id":"abc","data_uri":"http://example.com/data/abc.jpg","meta_uri":"http://example.com/meta/abc.json","stored_at":"0001-02-03T04:05:06.000000007Z"}`,
		},
	}
	for i, test := range tests {
		data, err := json.Marshal(test.item)
		if err != nil {
			t.Fatalf("%d DstItem failed to marshal: %s", i, err)
		}
		if test.json != "" {
			if string(data) != test.json {
				t.Errorf("%d DstItem JSON\ngot  %s\nwant %s", i, data, test.json)
			}
		}
		log.Printf("%d DstItem JSON %s", i, data)
		got := DstItem{}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Fatalf("%d DstItem failed to unmarshal: %s", i, err)
		}
		if !reflect.DeepEqual(test.item, got) {
			t.Errorf("%d DstItem Roundtrip\ngot  %#v\nwant %#v", i, got, test.item)
		}
	}
}
