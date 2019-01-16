==========================================
! Welcome to your personal photo archive !
==========================================

     -----------------------------
     DO NOT EDIT OR MOVE ANY FILES
     -----------------------------

The files here are maintained by Recentralized.

    https://www.recentralized.org

Their structure and format is open source and well documented.

    https://github.com/recentralized/structure


How to view photos and metadata
===============================

Recentralized provides free tools to browse and manage this set of files, but
we also strive to make them accessible by hand and with common unix tools.

    Visit www.recentralized.org for options.


If you'd like to view them by hand, here's a quick overview:

Photos are stored under `media`, and organized by year, day, and named after a
fingerprint of their content plus, their format:

    media/<year>/<day>/<fingerprint>.<format>

For example:

    media/2018/2018-01-07/f383c624d283bf1c81e4e89c60cf3bfc00c6bf57.jpg

---

Metadata is stored in JSON files under `meta`. To find the metadata for a
photo, use the photo's fingerprint. The first four characters are used as
directory names, with the remaining making up the file name:

    meta/<1:2>/<3:4/<4:>.json

For example:

    meta/f3/83/c624d283bf1c81e4e89c60cf3bfc00c6bf57.json

A metadata file contains information inherent to the photo-such as embedded
EXIF data-as well as information gathered from the place it was found. You can
use the tool `jq` to explore it [1].

---

The file `index.json` contains a list of files stored here, and where each file
came from. You can use the tool `jq` to explore it [1].


-------------------------------------------------------------------------------

[1] jq https://stedolan.github.io/jq/
