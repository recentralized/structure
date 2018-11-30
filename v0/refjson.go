package structure

import (
	"encoding/json"

	"github.com/recentralized/structure/cid"
)

type uRefJSON struct {
	Hash cid.ContentID `json:"hash"`
	Srcs []SrcItem     `json:"srcs"`
	Dsts []DstItem     `json:"dsts"`
}

// MarshalJSON implements json.Marshaler.
func (r URef) MarshalJSON() ([]byte, error) {
	var rj uRefJSON
	rj.Hash = r.Hash
	if len(r.Srcs) == 0 {
		rj.Srcs = make([]SrcItem, 0)
	} else {
		rj.Srcs = r.Srcs
	}
	if len(r.Dsts) == 0 {
		rj.Dsts = make([]DstItem, 0)
	} else {
		rj.Dsts = r.Dsts
	}
	return json.Marshal(rj)
}

// UnmarshalJSON implements json.Unmarshaler.
func (r *URef) UnmarshalJSON(data []byte) error {
	var rj uRefJSON
	if err := json.Unmarshal(data, &rj); err != nil {
		return err
	}
	r.Hash = rj.Hash
	if len(rj.Srcs) > 0 {
		r.Srcs = rj.Srcs
	}
	if len(rj.Dsts) > 0 {
		r.Dsts = rj.Dsts
	}
	return nil
}
