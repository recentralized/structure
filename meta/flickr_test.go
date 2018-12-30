package meta

import (
	"testing"
	"time"
)

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
		{
			desc: "all data",
			data: &FlickrMediaComment{
				ID:       "101630-3022774962-72157608889750784",
				UserID:   "36521980389@N01",
				Username: "Mike Monteiro",
				Text:     "I love chess club reuntions!",
				Date:     datePtr(2008, 11, 11, 19, 23, 38, 0, time.UTC),
				URL:      "https://www.flickr.com/photos/fss/3022774962/#comment72157608889750784",
			},
			json: `{
			  "id": "101630-3022774962-72157608889750784",
			  "user_id": "36521980389@N01",
			  "username": "Mike Monteiro",
			  "text": "I love chess club reuntions!",
			  "date": "2008-11-11T19:23:38Z",
			  "url": "https://www.flickr.com/photos/fss/3022774962/#comment72157608889750784"
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaComment{})
		})
	}
}
