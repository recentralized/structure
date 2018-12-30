package meta

import (
	"encoding/json"
	"reflect"
	"testing"
)

func assertJSONRoundtrip(t *testing.T, input interface{}, jsonString string, output interface{}) {
	data, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Failed to marshal: %s", err)
	}
	if string(data) != jsonString {
		t.Errorf("JSON\ngot  %s\nwant %s", data, jsonString)
	}
	got := output
	err = json.Unmarshal(data, &got)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s", err)
	}
	if !reflect.DeepEqual(input, got) {
		t.Errorf("Roundtrip\ngot  %#v\nwant %#v", got, input)
	}
}

func TestFlickrJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMedia
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMedia{},
			json: `{"id":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMedia{})
		})
	}
}

func TestFlickrMediaFaveJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaFave
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaFave{},
			json: `{"user_id":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaFave{})
		})
	}
}

func TestFlickrMediaGeoJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaGeo
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaGeo{},
			json: `{"latitude":0,"longitude":0}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaGeo{})
		})
	}
}

func TestFlickrPlaceJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrPlace
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrPlace{},
			json: `{"woe_id":"","latitude":0,"longitude":0}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrPlace{})
		})
	}
}

func TestFlickrMediaPersonJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaPerson
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaPerson{},
			json: `{"user_id":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaPerson{})
		})
	}
}

func TestFlickrMediaNoteJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaNote
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaNote{},
			json: `{"id":"","text":"","coords":{"x":0,"y":0,"w":0,"h":0}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaNote{})
		})
	}
}

func TestFlickrMediaInSetJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaInSet
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaInSet{},
			json: `{"id":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaInSet{})
		})
	}
}

func TestFlickrMediaInPoolJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaInPool
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaInPool{},
			json: `{"id":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaInPool{})
		})
	}
}

func TestFlickrMediaCommentJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaComment
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaComment{},
			json: `{"id":"","user_id":"","text":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaComment{})
		})
	}
}
