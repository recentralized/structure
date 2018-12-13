package data

import (
	"encoding/json"
	"log"
	"testing"
)

func TestHashJSON(t *testing.T) {
	hashIn := LiteralHash("abc")
	data, err := json.Marshal(hashIn)
	if err != nil {
		t.Fatalf("Marshal() failed: %s", err)
	}
	log.Printf("DATA %v", data)
	hashOut := &Hash{}
	if err := json.Unmarshal(data, hashOut); err != nil {
		t.Fatalf("Unmarshal() failed: %s", err)
	}
	if !hashIn.Equal(*hashOut) {
		t.Fatalf("Round-trip failed: got %s want %s", hashOut, hashIn)
	}
}
func TestHashDatabase(t *testing.T) {
	hashIn := LiteralHash("abc")
	val, err := hashIn.Value()
	if err != nil {
		t.Fatalf("Value() failed: %s", err)
	}
	_, ok := val.([]byte)
	if !ok {
		t.Fatalf("Value() did not return bytes")
	}
	hashOut := &Hash{}
	if err := hashOut.Scan(val); err != nil {
		t.Fatalf("Scan() failed: %s", err)
	}
	if !hashIn.Equal(*hashOut) {
		t.Fatalf("Round-trip failed: got %s want %s", hashOut, hashIn)
	}
}
