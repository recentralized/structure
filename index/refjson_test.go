package index

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/recentralized/structure/cid"
	"github.com/recentralized/structure/uri"
)

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
				Hash: cid.NewLiteral("xyz"),
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
						StoredAt: time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
					},
				},
			},
			json: `{"hash":"xyz","srcs":[{"src_id":"a","data_uri":"http://example.com/data.jpg","meta_uri":"http://example.com/meta.json","modified_at":"2015-02-03T04:05:06.000000007Z"}],"dsts":[{"dst_id":"abc","data_uri":"http://example.com/data/abc.jpg","meta_uri":"http://example.com/meta/abc.json","stored_at":"0001-02-03T04:05:06.000000007Z"}]}`,
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
