package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/recentralized/structure/cid"
	"github.com/recentralized/structure/index"
	"github.com/recentralized/structure/uri"
)

func main() {
	index, err := buildIndex()
	if err != nil {
		fmt.Printf("Failed to build index: %s", err)
		os.Exit(1)
	}

	err = addRefs(index)
	if err != nil {
		fmt.Printf("Failed to add refs: %s", err)
		os.Exit(1)
	}

	data, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		fmt.Printf("Failed to create json: %s", err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}

func buildIndex() (*index.Index, error) {
	srcPath, err := uri.NewDirPath("/tmp/src")
	if err != nil {
		return nil, fmt.Errorf("Could not create src path: %s", err)
	}

	dstPath, err := uri.NewDirPath("/tmp/dst")
	if err != nil {
		return nil, fmt.Errorf("Could not create dst path: %s", err)
	}

	src := index.NewSrc(srcPath.URI)
	dst := index.NewDstWithStandardPaths(dstPath.URI)

	idx := &index.Index{
		Srcs: []index.Src{src},
		Dsts: []index.Dst{dst},
	}

	return idx, nil
}

func addRefs(idx *index.Index) error {
	src := idx.Srcs[0]
	dst := idx.Dsts[0]

	data := []byte("fictional image data")
	cid, err := cid.New(bytes.NewReader(data))
	if err != nil {
		return err
	}

	srcItem, err := buildSrcItem(src)
	if err != nil {
		return err
	}

	dstItem, err := buildDstItem(cid, dst)

	idx.AddRef(index.Ref{
		Hash: cid,
		Src:  srcItem,
		Dst:  dstItem,
	})

	return nil
}

func buildSrcItem(src index.Src) (index.SrcItem, error) {
	var item index.SrcItem

	dataPath := uri.TrustedNew("fictional/image.jpg")
	dataURI, err := src.SrcURI.ResolveReference(dataPath)
	if err != nil {
		return item, err
	}

	metaPath := uri.TrustedNew("fictional/image.xmp")
	metaURI, err := src.SrcURI.ResolveReference(metaPath)
	if err != nil {
		return item, err
	}

	item = index.SrcItem{
		SrcID:      src.SrcID,
		DataURI:    dataURI,
		MetaURI:    metaURI,
		ModifiedAt: time.Date(2018, 11, 12, 0, 0, 0, 0, time.UTC),
	}

	return item, nil
}

func buildDstItem(cid cid.ContentID, dst index.Dst) (index.DstItem, error) {
	var item index.DstItem

	dataPath := uri.TrustedNew(fmt.Sprintf("%s.jpg", cid.String()))
	dataURI, err := dst.DataURI.ResolveReference(dataPath)
	if err != nil {
		return item, err
	}

	metaPath := uri.TrustedNew(fmt.Sprintf("%s.json", cid.String()))
	metaURI, err := dst.MetaURI.ResolveReference(metaPath)
	if err != nil {
		return item, err
	}

	item = index.DstItem{
		DstID:    dst.DstID,
		DataURI:  dataURI,
		MetaURI:  metaURI,
		StoredAt: time.Date(2018, 11, 13, 0, 0, 0, 0, time.UTC),
	}
	return item, nil
}
