package cid

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		desc string
		cid  string
	}{
		{
			desc: "hash format",
			cid:  "b8dfb080bc33fb564249e34252bf143d88fc018f",
		},
		{
			desc: "cidv0",
			cid:  "Qmc6SoJUtjspmudTyBHk71prbGnd7ajhS6uxCLsy8NtxEL",
		},
		{
			desc: "cidv1",
			cid:  "zb2rhkQ5HMh8b8qj6V1xH42nvDKMYW7q54SLsi2W1mYtes8S4",
		},
	}
	for _, tt := range tests {
		cid, err := Parse(tt.cid)
		if err != nil {
			t.Fatalf("%q failed: %s", tt.desc, err)
		}
		if got, want := cid.String(), tt.cid; got != want {
			t.Errorf("%q String() got %s want %s", tt.desc, got, want)
		}
	}
}

func TestNew(t *testing.T) {
	if defaultFormat != Hash {
		t.Fatalf("default format changed")
	}
	cid, err := New(bytes.NewBufferString("testing 12"))
	if err != nil {
		t.Fatalf("failed to new: %s", err)
	}
	if cid.cid != nil {
		t.Fatalf("cid must be nil")
	}
	if cid.hash == nil {
		t.Fatalf("hash must not be nil")
	}
}

func TestNewInFormat(t *testing.T) {
	tests := []struct {
		desc     string
		fmt      Format
		input    string
		wantCID  string
		wantHash string
	}{
		{
			desc:     "hash format",
			fmt:      Hash,
			input:    "testing 123",
			wantCID:  "b8dfb080bc33fb564249e34252bf143d88fc018f",
			wantHash: "b8dfb080bc33fb564249e34252bf143d88fc018f",
		},
		{
			desc:     "cidv0",
			fmt:      CidV0,
			input:    "testing 123",
			wantCID:  "Qmc6SoJUtjspmudTyBHk71prbGnd7ajhS6uxCLsy8NtxEL",
			wantHash: "Qmc6SoJUtjspmudTyBHk71prbGnd7ajhS6uxCLsy8NtxEL",
		},
		{
			desc:     "cidv1",
			fmt:      CidV1,
			input:    "testing 123",
			wantCID:  "zb2rhkQ5HMh8b8qj6V1xH42nvDKMYW7q54SLsi2W1mYtes8S4",
			wantHash: "Qmc6SoJUtjspmudTyBHk71prbGnd7ajhS6uxCLsy8NtxEL",
		},
	}
	for _, tt := range tests {
		cid, err := NewInFormat(bytes.NewBufferString(tt.input), tt.fmt)
		if err != nil {
			t.Fatalf("%q failed: %s", tt.desc, err)
		}
		if got, want := cid.String(), tt.wantCID; got != want {
			t.Errorf("%q String() got %s want %s", tt.desc, got, want)
		}
		if got, want := cid.Hash(), tt.wantHash; got != want {
			t.Errorf("%q Hash() got %s want %s", tt.desc, got, want)
		}
	}
}

func TestEquality(t *testing.T) {
	build := func(str string, fmt Format) ContentID {
		cid, err := NewInFormat(bytes.NewBufferString(str), fmt)
		if err != nil {
			t.Fatalf("creating cid: %s", err)
		}
		return cid
	}
	tests := []struct {
		desc          string
		a             ContentID
		b             ContentID
		wantEqual     bool
		wantEqualHash bool
	}{
		{
			desc:          "equal hash",
			a:             build("a", Hash),
			b:             build("a", Hash),
			wantEqual:     true,
			wantEqualHash: true,
		},
		{
			desc:          "unequal hash",
			a:             build("a", Hash),
			b:             build("b", Hash),
			wantEqual:     false,
			wantEqualHash: false,
		},
		{
			desc:          "equal cidV0",
			a:             build("a", CidV0),
			b:             build("a", CidV0),
			wantEqual:     true,
			wantEqualHash: true,
		},
		{
			desc:          "unequal cidV0",
			a:             build("a", CidV0),
			b:             build("b", CidV0),
			wantEqual:     false,
			wantEqualHash: false,
		},
		{
			desc:          "equal cidV1",
			a:             build("a", CidV1),
			b:             build("a", CidV1),
			wantEqual:     true,
			wantEqualHash: true,
		},
		{
			desc:          "unequal cidV1",
			a:             build("a", CidV1),
			b:             build("b", CidV1),
			wantEqual:     false,
			wantEqualHash: false,
		},
		{
			desc:          "equal cidv0 and cidV1",
			a:             build("a", CidV0),
			b:             build("a", CidV1),
			wantEqual:     false,
			wantEqualHash: true,
		},
		{
			desc:          "unequal cidv0 and cidV1",
			a:             build("a", CidV0),
			b:             build("b", CidV1),
			wantEqual:     false,
			wantEqualHash: false,
		},
	}
	for _, tt := range tests {
		got := tt.a.Equal(tt.b)
		if got, want := got, tt.wantEqual; !reflect.DeepEqual(got, want) {
			t.Errorf("%q Equal()\ngot %t want %t", tt.desc, got, want)
		}
		got = tt.a.EqualHash(tt.b)
		if got, want := got, tt.wantEqualHash; !reflect.DeepEqual(got, want) {
			t.Errorf("%q EqualHash()\ngot %t want %t", tt.desc, got, want)
		}
	}
}
