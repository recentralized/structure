package cid

import (
	"bytes"
	"encoding/json"
	"log"
	"reflect"
	"testing"
)

func TestCIDJSON(t *testing.T) {
	newCID := func(fmt Format, input string) ContentID {
		cid, err := NewInFormat(bytes.NewBufferString(input), fmt)
		if err != nil {
			t.Fatalf("failed to cid: %s", err)
		}
		return cid
	}
	tests := []struct {
		desc string
		cid  ContentID
		json string
	}{
		{
			desc: "zero value",
			cid:  ContentID{},
			json: `""`,
		},
		{
			desc: "test literal",
			cid:  NewLiteral("abc"),
			json: `"abc"`,
		},
		{
			desc: "hash",
			cid:  newCID(Hash, "testing 123"),
			json: `"b8dfb080bc33fb564249e34252bf143d88fc018f"`,
		},
		{
			desc: "cidv0",
			cid:  newCID(CidV0, "testing 123"),
			json: `"Qmc6SoJUtjspmudTyBHk71prbGnd7ajhS6uxCLsy8NtxEL"`,
		},
		{
			desc: "cidv1",
			cid:  newCID(CidV1, "testing 123"),
			json: `"zb2rhkQ5HMh8b8qj6V1xH42nvDKMYW7q54SLsi2W1mYtes8S4"`,
		},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.cid)
		if err != nil {
			t.Errorf("%q ContentID error encoding JSON: %s", tt.desc, err)
			continue
		}
		if string(data) != tt.json {
			t.Errorf("%q JSON got, %s, want %s", tt.desc, data, tt.json)
			continue
		}
		log.Printf("%q ContentID JSON %s", tt.desc, data)
		got := ContentID{}
		err = json.Unmarshal(data, &got)
		if err != nil {
			t.Errorf("%q ContentID error decoding JSON: %s", tt.desc, err)
			continue
		}
		if !reflect.DeepEqual(tt.cid, got) {
			t.Errorf("%q ContentID JSON\ngot  %#v\nwant %#v", tt.desc, got, tt.cid)
		}
	}
}
