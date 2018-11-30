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

### Index

The index is a catalog of content that's been stored. Where it was found (the
source), and where it was put (the destination). An index can store any number
of sources and any number of destinations.

```json
{
  "srcs": [
    {
      "src_id": "e8400c72-f7d0-53f9-98ca-ee23238231fe",
      "src_uri": "file:///tmp/src/"
    }
  ],
  "dsts": [
    {
      "dst_id": "b9fa3564-95cb-57b3-b2cd-32123cc2032b",
      "index_uri": "file:///tmp/dst/",
      "data_uri": "file:///tmp/dst/data/",
      "meta_uri": "file:///tmp/dst/meta/"
    }
  ]
}
```

### Meta

Meta is a document describing the content. Each piece of content may have none,
some, or all possible information available.

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

---

Copyright (c) 2018 Ryan Carver / www.recentralized.org

License: MIT
