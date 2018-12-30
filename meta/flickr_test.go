package meta

import (
	"testing"
	"time"
)

func TestFlickrJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMedia
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMedia{},
			json: `{"id":""}`,
		},
		{
			desc: "all data",
			// <?xml version="1.0" encoding="utf-8" ?>
			// <rsp stat="ok">
			//   <photo id="12688635254" secret="b428e566c1" server="3714" farm="4" dateuploaded="1393040381" isfavorite="0" license="4" safety_level="0" rotation="0" originalsecret="b00b8e351e" originalformat="jpg" views="399" media="photo">
			//     <owner nsid="47882233@N00" username="rcarver" realname="Ryan Carver" location="San Francisco, USA" iconserver="23" iconfarm="1" path_alias="fss" />
			//     <title>Mama &amp; Weston</title>
			//     <description>At &lt;a href=&quot;http://www.cafestjorge.com&quot; rel=&quot;noreferrer nofollow&quot;&gt;Cafe St. Jorge&lt;/a&gt;, a cute little Portuguese café in San Francisco.</description>
			//     <visibility ispublic="1" isfriend="0" isfamily="0" />
			//     <dates posted="1393040381" taken="2014-02-17 07:03:32" takengranularity="0" takenunknown="0" lastupdate="1516086752" />
			//     <permissions permcomment="3" permaddmeta="2" />
			//     <editability cancomment="1" canaddmeta="1" />
			//     <publiceditability cancomment="1" canaddmeta="0" />
			//     <usage candownload="1" canblog="1" canprint="1" canshare="1" />
			//     <comments>1</comments>
			//     <notes>
			//       <note id="72157676638204768" photo_id="12688635254" author="47882233@N00" authorname="rcarver" authorrealname="Ryan Carver" authorispro="1" authorisdeleted="0" x="172" y="40" w="215" h="168" pro_badge="legacy">these two</note>
			//       <note id="72157703126973021" photo_id="12688635254" author="47882233@N00" authorname="rcarver" authorrealname="Ryan Carver" authorispro="1" authorisdeleted="0" x="260" y="298" w="118" h="74" pro_badge="legacy">hasselblad + light</note>
			//     </notes>
			//     <people haspeople="1" />
			//     <tags>
			//       <tag id="101630-12688635254-83965415" author="47882233@N00" authorname="rcarver" raw="westoncarver" machine_tag="0">westoncarver</tag>
			//       <tag id="101630-12688635254-11746" author="47882233@N00" authorname="rcarver" raw="hasselblad" machine_tag="0">hasselblad</tag>
			//     </tags>
			//     <urls>
			//       <url type="photopage">https://www.flickr.com/photos/fss/12688635254/</url>
			//     </urls>
			//   </photo>
			// </rsp>
			data: &FlickrMedia{
				ID:            "12688635254",
				UserID:        "47882233@N00",
				Username:      "rcarver",
				Title:         "Mama & Weston",
				Description:   "At Cafe St. Jorge, a cute little Portuguese café in San Francisco.",
				TakenAt:       datePtr(2014, 2, 17, 7, 3, 32, 0, time.UTC),
				PostedAt:      datePtr(2014, 2, 22, 3, 39, 41, 0, time.UTC),
				LastUpdateAt:  datePtr(2018, 1, 16, 7, 12, 32, 0, time.UTC),
				URL:           "https://www.flickr.com/photos/fss/12688635254/",
				Visibility:    FlickrPublic,
				IsPublic:      true,
				IsFriendsOnly: false,
				IsFamilyOnly:  false,
				License:       "All Rights Reserved",
				LicenseURL:    "",
				Views:         100,
				// an empty element for each array.
				Faves:    []FlickrMediaFave{{}},
				Tags:     []FlickrMediaTag{{}},
				People:   []FlickrMediaPerson{{}},
				Notes:    []FlickrMediaNote{{}},
				Sets:     []FlickrMediaInSet{{}},
				Pools:    []FlickrMediaInPool{{}},
				Comments: []FlickrMediaComment{{}},
			},
			json: `{
			  "id": "12688635254",
			  "user_id": "47882233@N00",
			  "username": "rcarver",
			  "title": "Mama & Weston",
			  "description": "At Cafe St. Jorge, a cute little Portuguese café in San Francisco.",
			  "taken_at": "2014-02-17T07:03:32Z",
			  "posted_at": "2014-02-22T03:39:41Z",
			  "last_update_at": "2018-01-16T07:12:32Z",
			  "url": "https://www.flickr.com/photos/fss/12688635254/",
			  "visibility": "public",
			  "is_public": true,
			  "license": "All Rights Reserved",
			  "views": 100,
			  "faves": [
			    {
			      "user_id": ""
			    }
			  ],
			  "tags": [
			    {
			      "id": ""
			    }
			  ],
			  "people": [
			    {
			      "user_id": ""
			    }
			  ],
			  "notes": [
			    {
			      "id": "",
			      "text": "",
			      "coords": { "h": 0, "w": 0, "x": 0, "y": 0 }
			    }
			  ],
			  "sets": [
			    {
			      "id": ""
			    }
			  ],
			  "pools": [
			    {
			      "id": ""
			    }
			  ],
			  "comments": [
			    {
			      "id": "",
			      "user_id": "",
			      "text": ""
			    }
			  ]
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMedia{})
		})
	}
}

func TestFlickrMediaFaveJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaFave
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaFave{},
			json: `{"user_id":""}`,
		},
		{
			desc: "all data",
			// <person nsid="51035806117@N01" username="Rob Dumas" realname="Robert Dumas" favedate="1368902430" iconserver="2836" iconfarm="3" contact="0" friend="0" family="0" />
			data: &FlickrMediaFave{
				UserID:   "51035806117@N01",
				Username: "Rob Dumas",
				Date:     datePtr(2013, 5, 18, 18, 40, 30, 0, time.UTC),
			},
			json: `{
			  "user_id": "51035806117@N01",
			  "username": "Rob Dumas",
			  "date": "2013-05-18T18:40:30Z"
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaFave{})
		})
	}
}

func TestFlickrMediaGeoJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaGeo
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaGeo{},
			json: `{"latitude":0,"longitude":0}`,
		},
		{
			desc: "all data",
			// <photo id="10918764525">
			//   <location latitude="37.772546" longitude="-122.460329" accuracy="16" context="2" place_id="Oox27QtTUb9cqOQvnQ" woeid="23512032">
			//     <neighbourhood place_id="Oox27QtTUb9cqOQvnQ" woeid="23512032">Inner Richmond</neighbourhood>
			//     <locality place_id="7.MJR8tTVrIO1EgB" woeid="2487956">San Francisco</locality>
			//     <county place_id=".7sOmlRQUL9nK.kMzA" woeid="12587707">San Francisco</county>
			//     <region place_id="NsbUWfBTUb4mbyVu" woeid="2347563">California</region>
			//     <country place_id="nz.gsghTUb4c2WAecA" woeid="23424977">United States</country>
			//   </location>
			// </photo>
			data: &FlickrMediaGeo{
				WoeID:     "23512032",
				Latitude:  37.772546,
				Longitude: -122.460329,
				Accuracy:  16,
				Context:   FlickrGeoContextInside,
				Places: []FlickrPlace{
					{
						WoeID:     "23512032",
						Name:      "Inner Richmond",
						Type:      FlickrPlaceNeighborhood,
						Latitude:  37.779,
						Longitude: -122.468,
					},
				},
			},
			json: `{
			  "woe_id": "23512032",
			  "latitude": 37.772546,
			  "longitude": -122.460329,
			  "accuracy": 16,
			  "context": "inside",
			  "places": [
			    {
			      "latitude": 37.779,
			      "longitude": -122.468,
			      "name": "Inner Richmond",
			      "type": "neighborhood",
			      "woe_id": "23512032"
			    }
			  ]
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaGeo{})
		})
	}
}

func TestFlickrPlaceJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrPlace
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrPlace{},
			json: `{"woe_id":"","latitude":0,"longitude":0}`,
		},
		{
			desc: "all data",
			// <neighbourhood place_id="Oox27QtTUb9cqOQvnQ" woeid="23512032" latitude="37.779" longitude="-122.468" place_url="/United+States/California/San+Francisco/Inner+Richmond">Inner Richmond, San Francisco, CA, US, United States</neighbourhood>
			// <neighbourhood place_id="Oox27QtTUb9cqOQvnQ" woeid="23512032">Inner Richmond</neighbourhood>
			data: &FlickrPlace{
				WoeID:     "23512032",
				Name:      "Inner Richmond",
				Type:      "neighborhood",
				Latitude:  37.779,
				Longitude: -122.468,
			},
			json: `{
			  "woe_id": "23512032",
			  "name": "Inner Richmond",
			  "type": "neighborhood",
			  "latitude": 37.779,
			  "longitude": -122.468
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrPlace{})
		})
	}
}

func TestFlickrMediaPersonJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaPerson
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaPerson{},
			json: `{"user_id":""}`,
		},
		{
			desc: "all data",
			// <person nsid="37996612733@N01" username="Erika Hall" iconserver="1" iconfarm="1" realname="Erika Hall" path_alias="mulegirl" added_by="47882233@N00" is_deleted="0" user_contact="1" user_friend="1" user_family="0" />
			data: &FlickrMediaPerson{
				UserID:        "37996612733@N01",
				Username:      "Erika Hall",
				AddedByUserID: "47882233@N00",
			},
			json: `{
			  "user_id": "37996612733@N01",
			  "username": "Erika Hall",
			  "added_by_user_id": "47882233@N00"
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaPerson{})
		})
	}
}

func TestFlickrMediaNoteJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaNote
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaNote{},
			json: `{"id":"","text":"","coords":{"x":0,"y":0,"w":0,"h":0}}`,
		},
		{
			desc: "all data",
			// <note id="72157676638204768" photo_id="12688635254" author="47882233@N00" authorname="rcarver" authorrealname="Ryan Carver" authorispro="1" authorisdeleted="0" x="172" y="40" w="215" h="168" pro_badge="legacy">these two</note>
			data: &FlickrMediaNote{
				ID:       "72157676638204768",
				UserID:   "47882233@N00",
				Username: "rcarver",
				Text:     "these two",
				Coords:   NormalizedCoords{X: 172, Y: 40, W: 215, H: 168},
			},
			json: `{
			  "id": "72157676638204768",
			  "user_id": "47882233@N00",
			  "username": "rcarver",
			  "text": "these two",
			  "coords": {
			    "x": 172,
			    "y": 40,
			    "w": 215,
			    "h": 168
			  }
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaNote{})
		})
	}
}

func TestFlickrMediaInSetJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaInSet
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaInSet{},
			json: `{"id":""}`,
		},
		{
			desc: "all data",
			// <set title="029 500cm / Portra 400 NC" id="72157623276212363" primary="4346559546" secret="cf62fd3a50" server="4008" farm="5" view_count="20" comment_count="0" count_photo="9" count_video="0" />
			data: &FlickrMediaInSet{
				ID:       "72157623276212363",
				Name:     "029 500cm / Portra 400 NC",
				Position: 10,
			},
			json: `{
			  "id": "72157623276212363",
			  "name": "029 500cm / Portra 400 NC",
			  "position": 10
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaInSet{})
		})
	}
}

func TestFlickrMediaInPoolJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaInPool
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaInPool{},
			json: `{"id":""}`,
		},
		{
			desc: "all data",
			// <pool title="I Shoot Film" url="/groups/ishootfilm/pool/" id="67377471@N00" iconserver="7386" iconfarm="8" members="110509" pool_count="2873763" />
			data: &FlickrMediaInPool{
				ID:   "67377471@N00",
				Name: "I Shoot Film",
				URL:  "/groups/ishootfilm/pool/",
			},
			json: `{
			  "id": "67377471@N00",
			  "name": "I Shoot Film",
			  "url": "/groups/ishootfilm/pool/"
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaInPool{})
		})
	}
}

func TestFlickrMediaCommentJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *FlickrMediaComment
		json string
	}{
		{
			desc: "zero value",
			data: &FlickrMediaComment{},
			json: `{"id":"","user_id":"","text":""}`,
		},
		{
			desc: "all data",
			// <comment id="101630-3022774962-72157608889750784" author="36521980389@N01" author_is_deleted="0" authorname="Mike Monteiro" iconserver="1" iconfarm="1" datecreate="1226431418" permalink="https://www.flickr.com/photos/fss/3022774962/#comment72157608889750784" path_alias="dorkmaster" realname="Mike Monteiro">I love chess club reunions!</comment>
			data: &FlickrMediaComment{
				ID:       "101630-3022774962-72157608889750784",
				UserID:   "36521980389@N01",
				Username: "Mike Monteiro",
				Text:     "I love chess club reuntions!",
				Date:     datePtr(2008, 11, 11, 19, 23, 38, 0, time.UTC),
				URL:      "https://www.flickr.com/photos/fss/3022774962/#comment72157608889750784",
			},
			json: `{
			  "id": "101630-3022774962-72157608889750784",
			  "user_id": "36521980389@N01",
			  "username": "Mike Monteiro",
			  "text": "I love chess club reuntions!",
			  "date": "2008-11-11T19:23:38Z",
			  "url": "https://www.flickr.com/photos/fss/3022774962/#comment72157608889750784"
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &FlickrMediaComment{})
		})
	}
}
