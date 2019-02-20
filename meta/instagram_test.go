package meta

import (
	"testing"
	"time"
)

func TestInstagramMediaJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *InstagramMedia
		json string
	}{
		{
			desc: "zero value",
			data: &InstagramMedia{},
			json: `{"id":""}`,
		},
		{
			desc: "all data",
			// {
			//   "id": "1979320569926821011_11073382793",
			//   "user": {
			//     "id": "11073382793",
			//     "full_name": "Golang Client",
			//     "profile_picture": "https://scontent-sjc3-1.cdninstagram.com/vp/504ac2fa79adb1d412b31cab19be8d36/5CDDD9F1/t51.2885-19/44884218_345707102882519_2446069589734326272_n.jpg?_nc_ht=scontent-sjc3-1.cdninstagram.com",
			//     "username": "go_ig_test_0219"
			//   },
			//   "images": {
			//     "thumbnail": {
			//       "width": 150,
			//       "height": 150,
			//       "url": "https://scontent.cdninstagram.com/vp/fd0f484647ad37dc3caf0a2cdf37ca16/5CE59582/t51.2885-15/e35/c0.135.1080.1080/s150x150/50552544_116846169429307_872782777322498633_n.jpg?_nc_ht=scontent.cdninstagram.com"
			//     },
			//     "low_resolution": {
			//       "width": 320,
			//       "height": 400,
			//       "url": "https://scontent.cdninstagram.com/vp/0eda6589295b6fa43fd5cf2731afd691/5CF9331A/t51.2885-15/e35/p320x320/50552544_116846169429307_872782777322498633_n.jpg?_nc_ht=scontent.cdninstagram.com"
			//     },
			//     "standard_resolution": {
			//       "width": 640,
			//       "height": 800,
			//       "url": "https://scontent.cdninstagram.com/vp/bd6167c8e4469e16f2f6c900a62c51b9/5CF7EFF6/t51.2885-15/sh0.08/e35/p640x640/50552544_116846169429307_872782777322498633_n.jpg?_nc_ht=scontent.cdninstagram.com"
			//     }
			//   },
			//   "created_time": "1550173420",
			//   "caption": {
			//     "id": "18002756710177046",
			//     "text": "Photo post #0219test",
			//     "created_time": "1550173420",
			//     "from": {
			//       "id": "11073382793",
			//       "full_name": "Golang Client",
			//       "profile_picture": "https://scontent-sjc3-1.cdninstagram.com/vp/504ac2fa79adb1d412b31cab19be8d36/5CDDD9F1/t51.2885-19/44884218_345707102882519_2446069589734326272_n.jpg?_nc_ht=scontent-sjc3-1.cdninstagram.com",
			//       "username": "go_ig_test_0219"
			//     }
			//   },
			//   "user_has_liked": false,
			//   "likes": {
			//     "count": 1
			//   },
			//   "tags": [
			//     "0219test"
			//   ],
			//   "filter": "Crema",
			//   "comments": {
			//     "count": 2
			//   },
			//   "type": "image",
			//   "link": "https://www.instagram.com/p/Bt39ZJLHKSTFwXShw402xx8W9loUPHTyH5BsqY0/",
			//   "location": {
			//     "latitude": 37.8029,
			//     "longitude": -122.2721,
			//     "name": "Oakland, California",
			//     "id": 213051194
			//   },
			//   "attribution": null,
			//   "users_in_photo": [
			//     {
			//       "user": {
			//         "username": "rcarver"
			//       },
			//       "position": {
			//         "x": 0.57568438,
			//         "y": 0.7938808374
			//       }
			//     }
			//   ]
			// }
			data: &InstagramMedia{
				ID:           "1979320569926821011_11073382793",
				URL:          "https://www.instagram.com/p/Bt39ZJLHKSTFwXShw402xx8W9loUPHTyH5BsqY0/",
				UserID:       "11073382793",
				Username:     "go_ig_test_0219",
				UserFullName: "Golang Client",
				Caption:      "Photo post #0219test",
				CreatedAt:    datePtr(2019, 2, 14, 7, 3, 32, 0, time.UTC),
				Filter:       "Crema",
				Tags:         []string{"0219test"},
				Location: &InstagramLocation{
					ID:        "213051194",
					Name:      "Oakland, California",
					Latitude:  37.8029,
					Longitude: -122.2721,
				},
				Likes: 1,
				TaggedUsers: []InstagramTaggedUser{
					{
						Username: "rcarver",
						X:        0.57568438,
						Y:        0.7938808374,
					},
				},
			},
			json: `{
			  "id": "1979320569926821011_11073382793",
			  "url": "https://www.instagram.com/p/Bt39ZJLHKSTFwXShw402xx8W9loUPHTyH5BsqY0/",
			  "user_id": "11073382793",
			  "username": "go_ig_test_0219",
			  "user_full_name": "Golang Client",
			  "caption": "Photo post #0219test",
			  "created_at": "2019-02-14T07:03:32Z",
			  "filter": "Crema",
			  "tags": [
			      "0219test"
			  ],
			  "location": {
			    "id": "213051194",
			    "name": "Oakland, California",
			    "latitude": 37.8029,
			    "longitude": -122.2721
			  },
			  "likes": 1,
			  "tagged_users": [
			    {
			      "username": "rcarver",
			      "x": 0.57568438,
			      "y": 0.7938808374
			    }
			  ]
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &InstagramMedia{})
		})
	}
}

func TestInstagramCommentJSON(t *testing.T) {
	tests := []struct {
		desc string
		data *InstagramComment
		json string
	}{
		{
			desc: "zero value",
			data: &InstagramComment{},
			json: `{"id":"","text":""}`,
		},
		{
			desc: "all data",
			// {
			//   "id": "18034730665003447",
			//   "from": {
			//     "username": "rcarver"
			//   },
			//   "text": "That’s about right.",
			//   "created_time": "1550177777"
			// },
			data: &InstagramComment{
				ID:       "18034730665003447",
				Username: "rcarver",
				Text:     "That’s about right.",
				Date:     datePtr(2019, 2, 14, 20, 56, 17, 0, time.UTC),
			},
			json: `{
			  "id": "18034730665003447",
			  "username": "rcarver",
			  "text": "That’s about right.",
			  "date": "2019-02-14T20:56:17Z"
			}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertJSONRoundtrip(t, tt.data, tt.json, &InstagramComment{})
		})
	}
}
