package meta

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/kr/pretty"
)

// assertJSONRoundtrip validates that input marshals to jsonString, then
// unmarshals back to input.  input and output must be pointer types. JSON
// compare is done with assertJSON so formatting is normalized.
func assertJSONRoundtrip(t *testing.T, input interface{}, jsonString string, output interface{}) {
	t.Helper()
	data, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Failed to marshal: %s", err)
	}
	assertJSON(t, string(data), jsonString)
	got := output
	err = json.Unmarshal(data, &got)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s", err)
	}
	if !reflect.DeepEqual(input, got) {
		t.Errorf("Roundtrip\ngot  %#v\nwant %#v", got, input)
	}
}

// assertJSON compares two JSON strings. It does this by decoding into a
// generic struct, then using reflect.DeepEqual. If they aren't equal, they're
// reformatted as JSON and printed along with a line diff. The `sub` function
// is optional, and used to return a subtree of the parsed gotJSON data before
// comparing.
func assertJSON(t *testing.T, gotJSON string, wantJSON string) {
	t.Helper()
	var gotData map[string]interface{}
	if err := json.Unmarshal([]byte(gotJSON), &gotData); err != nil {
		t.Fatalf("decode got: %s", err)
	}
	var wantData map[string]interface{}
	if err := json.Unmarshal([]byte(wantJSON), &wantData); err != nil {
		t.Fatalf("decode want: %s", err)
	}
	if got, want := gotData, wantData; !reflect.DeepEqual(got, want) {
		gj, _ := json.Marshal(got)
		wj, _ := json.Marshal(want)
		t.Errorf("got  %s\nwant %s", gj, wj)
		pretty.Ldiff(t, got, want)
	}
}

// datePtr is time.Date by returns a *time.Time pointer type.
func datePtr(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) *time.Time {
	d := time.Date(year, month, day, hour, min, sec, nsec, loc)
	return &d
}
