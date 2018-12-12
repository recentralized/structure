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
		uri  uriStringer
		json string
	}{
		{
			desc: "URI zero value",
			uri:  URI{},
			json: `""`,
		},
		{
			desc: "URI invald rawstr",
			uri:  URI{rawStr: "%"},
			json: `"%"`,
		},
		{
			desc: "URI",
			uri:  URI{url: newURL("https://www.example.com")},
			json: `"https://www.example.com"`,
		},
		{
			desc: "FilePath zero value",
			uri:  FilePath{},
			json: `""`,
		},
		{
			desc: "FilePath",
			uri:  FilePath{URI{url: newURL("file:///path")}},
			json: `"file:///path"`,
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

func TestURIDatabase(t *testing.T) {
	uriIn := TrustedNew("http://example.com/")
	val, err := uriIn.Value()
	if err != nil {
		t.Fatalf("Value() failed: %s", err)
	}
	_, ok := val.([]byte)
	if !ok {
		t.Fatalf("Value() did not return bytes")
	}
	uriOut := &URI{}
	if err := uriOut.Scan(val); err != nil {
		t.Fatalf("Scan() failed: %s", err)
	}
	if !uriIn.Equal(*uriOut) {
		t.Fatalf("Round-trip failed: got %s want %s", uriOut, uriIn)
	}
}
