package content

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestMetaJSON(t *testing.T) {
	tests := []struct {
		desc  string
		meta  Meta
		setup func(*Meta)
		json  string
	}{
		{
			desc: "zero value",
			meta: Meta{},
			json: `{"content_type":"","size":0}`,
		},
		{
			desc: "basic fields",
			meta: Meta{
				ContentType: JPG,
				Size:        100,
				Inherent: MetaContent{
					Created: time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC),
					Image: MetaImage{
						Width:  100,
						Height: 60,
					},
					Exif: Exif{
						"CreateData": ExifValue{
							ID:  "0x9004",
							Val: "2013:07:17 19:59:58",
						},
					},
				},
			},
			json: `{"content_type":"jpg","size":100,"inherent":{"created":"0001-02-03T04:05:06.000000007Z","image":{"width":100,"height":60},"exif":{"CreateData":{"id":"0x9004","val":"2013:07:17 19:59:58"}}}}`,
		},
		{
			desc: "src-specific fields",
			meta: Meta{
				Srcs: SrcSpecific{
					//Flickr:
				},
			},
			json: `{"content_type":"","size":0}`,
		},
	}
	for _, tt := range tests {
		if tt.setup != nil {
			tt.setup(&tt.meta)
		}
		data, err := json.Marshal(tt.meta)
		if err != nil {
			t.Fatalf("%q Meta failed to marshal: %s", tt.desc, err)
		}
		if string(data) != tt.json {
			t.Errorf("%q Meta JSON\ngot  %s\nwant %s", tt.desc, data, tt.json)
		}
		log.Printf("%q Meta JSON %s", tt.desc, data)
		got := Meta{}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Fatalf("%q Meta failed to unmarshal: %s", tt.desc, err)
		}
		if !reflect.DeepEqual(tt.meta, got) {
			t.Errorf("%q Meta Roundtrip\ngot  %#v\nwant %#v", tt.desc, got, tt.meta)
		}
	}
}
