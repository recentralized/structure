# Structure

Official structures and specs of the storage format. Reference implementation in Go (golang).

## Concepts

Structure supports a photo-centric storage system. Its primary goal is to consume content from many places (called _sources_), and store them in any number of places (called _destinations_). It keeps track of the sources, destinations, and each piece of content in the Index. Along the way, it extracts metadata from the content and stores it in Meta.

### Terminology

* Source (short: Src) - A places that's searched for content.
* Destination (short: Dst) - A place where content is stored.
* URI - The way that all content is referenced. Can specifically be any URL or URN.
* Hash - The way all content is identified, a unique fingerprint of the content.
* Index - A database of sources, destinations: where content was found, and where it was stored.
* Meta - Information extracted from a piece of content at source.

### Example

As an example, we'll use a filesystem-based approach. However, the specifics of data location, file naming, and storage format are all flexible. That said, for simplicity let's say we have a source directory at `/tmp/src` like so:

```
/tmp/src/
  fictional/
    image.jpg
    image.xmp
```

And we want to store that content at `/tmp/dst`. The storage format looks like this:
    
```
/tmp/dst/
  index.json
  data/
    f42a59131aaf2e5c475f8a35126b034549c05bd5.jpg
  meta/
    f42a59131aaf2e5c475f8a35126b034549c05bd5.json
```

The image has been stored at `<content-hash>.<format>` and its associated metadata at `<content-hash>.json`. 

### Index

The index is a catalog of content that's been stored. Where it was found (the
source), and where it was put (the destination). An index can store any number
of sources and any number of destinations. You can use the index to find content and retrieve it from the destination, or even put it back where it was originally found on the source.

The file `/tmp/dst/index.json` contains:

```json
{
  "srcs": [
    {
      "src_id":  "e8400c72-f7d0-53f9-98ca-ee23238231fe",
      "src_uri": "file:///tmp/src/"
    }
  ],
  "dsts": [
    {
      "dst_id":    "b9fa3564-95cb-57b3-b2cd-32123cc2032b",
      "index_uri": "file:///tmp/dst/",
      "data_uri":  "file:///tmp/dst/data/",
      "meta_uri":  "file:///tmp/dst/meta/"
    }
  ],
  "refs": [
    {
      "hash": "f42a59131aaf2e5c475f8a35126b034549c05bd5",
      "srcs": [
        {
          "src_id":      "e8400c72-f7d0-53f9-98ca-ee23238231fe",
          "data_uri":    "file:///tmp/src/fictional/image.jpg",
          "meta_uri":    "file:///tmp/src/fictional/image.xmp",
          "modified_at": "2018-11-12T00:00:00Z"
        }
      ],
      "dsts": [
        {
          "dst_id":    "b9fa3564-95cb-57b3-b2cd-32123cc2032b",
          "data_uri":  "file:///tmp/dst/data/f42a59131aaf2e5c475f8a35126b034549c05bd5.jpg",
          "meta_uri":  "file:///tmp/dst/meta/f42a59131aaf2e5c475f8a35126b034549c05bd5.json",
          "stored_at": "2018-11-13T00:00:00Z"
        }
      ]
    }
  ]
}
```
See [how to generate this output](examples/index/main.go).


### Meta

Meta is a document describing the content. Each piece of content may have none,
some, or all possible information available. You can use the meta document to find out information about a piece of content.

The file `/tmp/dst/meta/f42a59131aaf2e5c475f8a35126b034549c05bd5.json` contains:

```json
{
  "content_type": "jpg",
  "size": 1024,
  "inherent": {
    "created": "2018-01-01T00:00:00Z",
    "image": {
      "width": 3000,
      "height": 5000
    },
    "exif": {
      "ExposureTime": {
        "id": "ShutterSpeed",
        "val": "1/60"
      }
    }
  },
  "sidecar": {
    "exif": {
      "FNumber": {
        "id": "0x829d",
        "val": 1.8
      }
    }
  }
}
```
See [how to generate this output](examples/meta/main.go).

---

Copyright (c) 2018 Ryan Carver / www.recentralized.org

License: MIT
