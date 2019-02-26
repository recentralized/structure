package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/recentralized/structure/data"
	"github.com/recentralized/structure/dst"
	"github.com/recentralized/structure/index"
	"github.com/recentralized/structure/meta"
	"github.com/recentralized/structure/uri"
)

func main() {
	index, err := buildIndex()
	if err != nil {
		fmt.Printf("Failed to build index: %s", err)
		os.Exit(1)
	}

	layout := dst.NewFilesystemLayout()

	err = addRefs(layout, index)
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
	srcPath, err := uri.ParseDir("/tmp/src")
	if err != nil {
		return nil, fmt.Errorf("Could not create src path: %s", err)
	}

	dstPath, err := uri.ParseDir("/tmp/dst")
	if err != nil {
		return nil, fmt.Errorf("Could not create dst path: %s", err)
	}

	src := index.NewSrc(srcPath.URI)
	dst := index.NewDstAllAt(dstPath.URI)

	idx := index.New()
	idx.Srcs = []index.Src{src}
	idx.Dsts = []index.Dst{dst}

	return idx, nil
}

func addRefs(layout dst.Layout, idx *index.Index) error {
	src := idx.Srcs[0]
	dst := idx.Dsts[0]

	// Find something on the source
	srcItem, meta, err := buildSrcItem(src)
	if err != nil {
		return err
	}

	// Get data from the source.
	data := []byte(strings.Repeat("fictional image data", 100))

	// Calculate fingerprint of data.
	hash, err := layout.NewHash(bytes.NewReader(data))
	if err != nil {
		return err
	}

	// Decide how it will be stored on the destination.
	dstItem, err := buildDstItem(layout, dst, hash, meta)
	if err != nil {
		return err
	}

	// Calculate size of data to be stored.
	c, err := io.Copy(ioutil.Discard, bytes.NewReader(data))
	if err != nil {
		return err
	}
	dstItem.DataSize = c

	// Calculate size of meta to be stored.
	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return err
	}
	dstItem.MetaSize = int64(len(metaJSON))

	// Here you would store data and metadata on the destination.

	// Add a ref to the index.
	idx.AddRef(index.Ref{
		Hash: hash,
		Src:  srcItem,
		Dst:  dstItem,
	})

	return nil
}

func buildSrcItem(src index.Src) (index.SrcItem, *meta.Meta, error) {
	var item index.SrcItem

	dataPath := uri.TrustedNew("fictional/image.jpg")
	dataURI, err := src.SrcURI.ResolveReference(dataPath)
	if err != nil {
		return item, nil, err
	}

	metaPath := uri.TrustedNew("fictional/image.xmp")
	metaURI, err := src.SrcURI.ResolveReference(metaPath)
	if err != nil {
		return item, nil, err
	}

	item = index.SrcItem{
		SrcID:      src.SrcID,
		DataURI:    dataURI,
		MetaURI:    metaURI,
		ModifiedAt: time.Date(2018, 11, 12, 0, 0, 0, 0, time.UTC),
	}

	doc := meta.New()
	doc.Type = data.JPG
	doc.Inherent = meta.Content{
		Created: time.Date(2018, 11, 10, 0, 0, 0, 0, time.UTC),
	}

	return item, doc, nil
}

func buildDstItem(layout dst.Layout, dst index.Dst, hash data.Hash, meta *meta.Meta) (index.DstItem, error) {
	var item index.DstItem

	dataURI := layout.DataURI(hash, meta)
	metaURI := layout.MetaURI(hash, meta)

	item = index.DstItem{
		DstID:     dst.DstID,
		DataURI:   dataURI,
		MetaURI:   metaURI,
		DataType:  data.Stored{Type: meta.Type},
		DataSize:  0, // updated with actual data size
		MetaSize:  0, // updated with actual meta size
		StoredAt:  time.Date(2018, 11, 13, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2018, 11, 14, 0, 0, 0, 0, time.UTC),
	}
	return item, nil
}
