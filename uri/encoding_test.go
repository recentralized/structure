package uri

import (
	"encoding/json"
	"log"
	"net/url"
	"reflect"
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
		if !reflect.DeepEqual(got, tt.uri) {
			t.Errorf("%q URI JSON\ngot  %#v\nwant %#v", tt.desc, got, tt.uri)
		}
	}
}

func TestPathJSON(t *testing.T) {
	tests := []struct {
		desc string
		path Path
		json string
	}{
		{
			desc: "zero value",
			path: Path{},
			json: `""`,
		},
		{
			desc: "file path",
			path: Path{
				RawPath: "/tmp/file",
			},
			json: `"file:///tmp/file"`,
		},
		{
			desc: "dir path",
			path: Path{
				RawPath: "/tmp/file",
				IsDir:   true,
			},
			json: `"file:///tmp/file/"`,
		},
		{
			desc: "badly encoded path",
			path: Path{
				RawPath: "/tmp/file%2with%20space",
			},
			json: `"file:///tmp/file%2with%20space"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			data, err := json.Marshal(tt.path)
			if err != nil {
				t.Fatalf("error encoding JSON: %s", err)
			}
			if string(data) != tt.json {
				t.Fatalf("JSON got, %s, want %s", data, tt.json)
			}
			got := Path{}
			err = json.Unmarshal(data, &got)
			if err != nil {
				t.Fatalf("error decoding JSON: %s", err)
			}
			if !reflect.DeepEqual(got, tt.path) {
				t.Errorf("JSON\ngot  %#v\nwant %#v", got, tt.path)
			}
		})
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

func TestPathDatabase(t *testing.T) {
	pathIn := Path{
		RawPath: "/tmp/file",
	}
	val, err := pathIn.Value()
	if err != nil {
		t.Fatalf("Value() failed: %s", err)
	}
	_, ok := val.([]byte)
	if !ok {
		t.Fatalf("Value() did not return bytes")
	}
	pathOut := &Path{}
	if err := pathOut.Scan(val); err != nil {
		t.Fatalf("Scan() failed: %s", err)
	}
	if pathIn != *pathOut {
		t.Fatalf("Round-trip failed: got %s want %s", pathOut, pathIn)
	}
}
