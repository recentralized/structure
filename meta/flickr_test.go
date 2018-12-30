package meta

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFlickrJSON(t *testing.T) {
	tests := []struct {
		desc   string
		flickr FlickrMedia
		json   string
	}{
		{
			desc:   "zero value",
			flickr: FlickrMedia{},
			json:   `{"id":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			data, err := json.Marshal(tt.flickr)
			if err != nil {
				t.Fatalf("Failed to marshal: %s", err)
			}
			if string(data) != tt.json {
				t.Errorf("JSON\ngot  %s\nwant %s", data, tt.json)
			}
			got := FlickrMedia{}
			err = json.Unmarshal(data, &got)
			if err != nil {
				t.Fatalf("Failed to unmarshal: %s", err)
			}
			if !reflect.DeepEqual(tt.flickr, got) {
				t.Errorf("Roundtrip\ngot  %#v\nwant %#v", got, tt.flickr)
			}
		})
	}
}
