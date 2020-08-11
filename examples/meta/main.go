package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/recentralized/structure/data"
	"github.com/recentralized/structure/index"
	"github.com/recentralized/structure/meta"
)

func main() {
	meta, err := buildMeta()
	if err != nil {
		fmt.Printf("Failed to build meta: %s", err)
		os.Exit(1)
	}

	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		fmt.Printf("Failed to create json: %s", err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}

func buildMeta() (*meta.Meta, error) {

	srcID := index.SrcID("e8400c72-f7d0-53f9-98ca-ee23238231fe")

	doc := meta.New()
	doc.Type = data.JPG
	doc.Size = 1024

	doc.Inherent = meta.Content{
		Created: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		Image: meta.Image{
			Width:  3000,
			Height: 5000,
		},
		Exif: meta.Exif{
			"ExposureTime": meta.ExifValue{ID: "ShutterSpeed", Val: "1/60"},
		},
	}
	doc.Src = map[index.SrcID]meta.SrcSpecific{
		srcID: {
			Sidecar: &meta.Content{
				Exif: meta.Exif{
					"FNumber": meta.ExifValue{ID: "0x829d", Val: 1.8},
				},
			},
		},
	}

	return doc, nil
}
