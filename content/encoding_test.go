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
			json: `{"version":"","content_type":"","size":0}`,
		},
		{
			desc: "basic fields",
			meta: Meta{
				Version:     "v1",
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
				Sidecar: MetaContent{
					Created: time.Date(2, 2, 3, 4, 5, 6, 7, time.UTC),
				},
			},
			json: `{"version":"v1","content_type":"jpg","size":100,"inherent":{"created":"0001-02-03T04:05:06.000000007Z","image":{"width":100,"height":60},"exif":{"CreateData":{"id":"0x9004","val":"2013:07:17 19:59:58"}}},"sidecar":{"created":"0002-02-03T04:05:06.000000007Z"}}`,
		},
		{
			desc: "src-specific fields",
			meta: Meta{
				Version: "v1",
				Srcs:    SrcSpecific{
					//Flickr:
				},
			},
			json: `{"version":"v1","content_type":"","size":0}`,
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

func TestExifValueJSON(t *testing.T) {
	tests := []struct {
		desc  string
		value ExifValue
		json  string
	}{
		{
			desc:  "zero value",
			value: ExifValue{},
			json:  `{"id":"","val":null}`,
		},
		{
			desc: "string",
			value: ExifValue{
				ID:  "0x0111",
				Val: "ok",
			},
			json: `{"id":"0x0111","val":"ok"}`,
		},
		{
			desc: "date as string",
			value: ExifValue{
				ID:  "0x0112",
				Val: "2014:02:09 18:11:38",
			},
			json: `{"id":"0x0112","val":"2014:02:09 18:11:38"}`,
		},
		{
			desc: "int as float",
			value: ExifValue{
				ID:  "0x010d",
				Val: float64(500),
			},
			json: `{"id":"0x010d","val":500}`,
		},
		{
			desc: "float",
			value: ExifValue{
				ID:  "float",
				Val: 1.12345,
			},
			json: `{"id":"float","val":1.12345}`,
		},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.value)
		if err != nil {
			t.Fatalf("%q Value failed to marshal: %s", tt.desc, err)
		}
		if string(data) != tt.json {
			t.Errorf("%q Value JSON\ngot  %s\nwant %s", tt.desc, data, tt.json)
		}
		log.Printf("%q Value JSON %s", tt.desc, data)
		got := ExifValue{}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Fatalf("%q Value failed to unmarshal: %s", tt.desc, err)
		}
		if !reflect.DeepEqual(tt.value, got) {
			t.Errorf("%q Value Roundtrip\ngot  %#v\nwant %#v", tt.desc, got, tt.value)
		}
	}
}
