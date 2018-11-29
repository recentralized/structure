package uri

import (
	"encoding/json"
	"log"
	"net/url"
	"testing"
)

func TestURIJSON(t *testing.T) {
	newURL := func(str string) *url.URL {
		u, err := url.Parse(str)
		if err != nil {
			t.Fatalf("failed to parse URL: %s", err)
		}
		return u
	}
	tests := []struct {
		desc string
		uri  URI
		json string
	}{
		{
			// Zero value.
			desc: "zero value",
			uri:  URI{},
			json: `""`,
		},
		{
			// invalid rawStr
			desc: "rawstr",
			uri:  URI{rawStr: "%"},
			json: `"%"`,
		},
		{
			// URI with error.
			desc: "uri",
			uri:  URI{url: newURL("https://www.example.com")},
			json: `"https://www.example.com"`,
		},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.uri)
		if err != nil {
			t.Errorf("%q URI error encoding JSON: %s", tt.desc, err)
			continue
		}
		if string(data) != tt.json {
			t.Errorf("%q JSON got, %s, want %s", tt.desc, data, tt.json)
			continue
		}
		log.Printf("%q URI JSON %s", tt.desc, data)
		got := URI{}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Errorf("%q URI error decoding JSON: %s", tt.desc, err)
			continue
		}
		if !got.Equal(tt.uri) {
			t.Errorf("%q URI JSON\ngot  %#v\nwant %#v", tt.desc, got, tt.uri)
		}
	}
}
