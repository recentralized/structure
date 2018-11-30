package cid

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	if defaultFormat != hash {
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
		fmt      format
		input    string
		wantCID  string
		wantHash string
	}{
		{
			desc:     "legacy hash format",
			fmt:      hash,
			input:    "testing 123",
			wantCID:  "b8dfb080bc33fb564249e34252bf143d88fc018f",
			wantHash: "b8dfb080bc33fb564249e34252bf143d88fc018f",
		},
		{
			desc:     "cidv0",
			fmt:      cidV0,
			input:    "testing 123",
			wantCID:  "Qmc6SoJUtjspmudTyBHk71prbGnd7ajhS6uxCLsy8NtxEL",
			wantHash: "Qmc6SoJUtjspmudTyBHk71prbGnd7ajhS6uxCLsy8NtxEL",
		},
		{
			desc:     "cidv1",
			fmt:      cidV1,
			input:    "testing 123",
			wantCID:  "zb2rhkQ5HMh8b8qj6V1xH42nvDKMYW7q54SLsi2W1mYtes8S4",
			wantHash: "Qmc6SoJUtjspmudTyBHk71prbGnd7ajhS6uxCLsy8NtxEL",
		},
	}
	for _, tt := range tests {
		cid, err := newInFormat(bytes.NewBufferString(tt.input), tt.fmt)
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
