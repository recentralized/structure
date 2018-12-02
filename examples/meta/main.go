package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/recentralized/structure/content"
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

func buildMeta() (*content.Meta, error) {

	meta := content.NewMeta()
	meta.ContentType = content.JPG
	meta.Size = 1024

	meta.Inherent = content.MetaContent{
		Created: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		Image: content.MetaImage{
			Width:  3000,
			Height: 5000,
		},
		Exif: content.Exif{
			"ExposureTime": content.ExifValue{ID: "ShutterSpeed", Val: "1/60"},
		},
	}
	meta.Sidecar = content.MetaContent{
		Exif: content.Exif{
			"FNumber": content.ExifValue{ID: "0x829d", Val: 1.8},
		},
	}

	return meta, nil
}
