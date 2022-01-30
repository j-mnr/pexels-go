package pexels

import "net/http"

// client.go test cases
var newTT = []struct {
	expectErr bool
	options   Options
}{
	{
		true,
		Options{}, // No API key
	},
	{
		false,
		Options{
			APIKey:     "pexels-api-key",
			UserAgent:  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.162 Safari/537.36",
			HTTPClient: &http.Client{},
		},
	},
	{
		false,
		Options{
			APIKey:    "pexels-api-key",
			UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.162 Safari/537.36",
		},
	},
}

var rateLimitHeadersTT = []struct {
	statusCode      int
	options         *Options
	photoID         uint64
	respBody        string
	headerLimit     string
	headerRemaining string
	headerReset     string
}{
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		2014422,
		`{"id":2014422,"width":3024,"height": 3024,"url":"https://www.pexels.com/photo/brown-rocks-during-golden-hour-2014422/"}`,
		"20000",
		"18000",
		"1625092515",
	},
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		2014422,
		`{"id":2014422,"width":3024,"height": 3024,"url":"https://www.pexels.com/photo/brown-rocks-during-golden-hour-2014422/"}`,
		"",
		"",
		"",
	},
}

var setRequestHeadersTT = []struct {
	endpoint  string
	APIKey    string
	UserAgent string
}{
	{"/photos/", "testAPIKey", "testUserAgent"},
	{"/videos/", "testAPIKey", ""},
	{"/collections/", "testAPIKey", "otherUserAgent"},
}

// video.go test cases
var getVideoTT = []struct {
	statusCode int
	opts       *Options
	ID         uint64
	respBody   string
}{
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		2499611,
		`{"id":2499611,"width":1080,"height":1920,"url":"https://www.pexels.com/video/2499611/","image":"https://images.pexels.com/videos/2499611/free-video-2499611.jpg?fit=crop&w=1200&h=630&auto=compress&cs=tinysrgb","duration":22,"user":{"id":680589,"name":"JoeyFarina","url":"https://www.pexels.com/@joey"},"video_files":[{"id":125004,"quality":"hd","file_type":"video/mp4","width":1080,"height":1920,"link":"https://player.vimeo.com/external/342571552.hd.mp4?s=6aa6f164de3812abadff3dde86d19f7a074a8a66&profile_id=175&oauth2_token_id=57447761"},{"id":125005,"quality":"sd","file_type":"video/mp4","width":540,"height":960,"link":"https://player.vimeo.com/external/342571552.sd.mp4?s=e0df43853c25598dfd0ec4d3f413bce1e002deef&profile_id=165&oauth2_token_id=57447761"},{"id":125006,"quality":"sd","file_type":"video/mp4","width":240,"height":426,"link":"https://player.vimeo.com/external/342571552.sd.mp4?s=e0df43853c25598dfd0ec4d3f413bce1e002deef&profile_id=139&oauth2_token_id=57447761"},{"id":125007,"quality":"hd","file_type":"video/mp4","width":720,"height":1280,"link":"https://player.vimeo.com/external/342571552.hd.mp4?s=6aa6f164de3812abadff3dde86d19f7a074a8a66&profile_id=174&oauth2_token_id=57447761"},{"id":125008,"quality":"sd","file_type":"video/mp4","width":360,"height":640,"link":"https://player.vimeo.com/external/342571552.sd.mp4?s=e0df43853c25598dfd0ec4d3f413bce1e002deef&profile_id=164&oauth2_token_id=57447761"},{"id":125009,"quality":"hls","file_type":"video/mp4","width":null,"height":null,"link":"https://player.vimeo.com/external/342571552.m3u8?s=53433233e4176eead03ddd6fea04d9fb2bce6637&oauth2_token_id=57447761"}],"video_pictures":[{"id":308178,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-0.jpg","nr":0},{"id":308179,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-1.jpg","nr":1},{"id":308180,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-2.jpg","nr":2},{"id":308181,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-3.jpg","nr":3},{"id":308182,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-4.jpg","nr":4},{"id":308183,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-5.jpg","nr":5},{"id":308184,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-6.jpg","nr":6},{"id":308185,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-7.jpg","nr":7},{"id":308186,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-8.jpg","nr":8},{"id":308187,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-9.jpg","nr":9},{"id":308188,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-10.jpg","nr":10},{"id":308189,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-11.jpg","nr":11},{"id":308190,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-12.jpg","nr":12},{"id":308191,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-13.jpg","nr":13},{"id":308192,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-14.jpg","nr":14}]}`,
	},
	{
		http.StatusNotFound,
		&Options{APIKey: "testAPIKey"},
		0,
		`{"status":404,"error":"Not Found"}`,
	},
}

var getPopularVideosTT = []struct {
	statusCode int
	opts       *Options
	params     *PopularVideoParams
	respBody   string
}{
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		&PopularVideoParams{
			MinWidth:    4096,
			MinHeight:   2160,
			MinDuration: 10,
			MaxDuration: 60,
			Page:        2,
			PerPage:     2,
		},
		`{"videos":[{"id":2953633,"width":4096,"height":2160,"url":"https://www.pexels.com/video/people-having-rest-in-a-ledge-2953633/","image":"https://images.pexels.com/videos/2953633/free-video-2953633.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":27,"user":{"id":1179532,"name":"KellyLacy","url":"https://www.pexels.com/@kelly-lacy-1179532"},"video_files":[{"id":187214,"quality":"sd","file_type":"video/mp4","width":640,"height":338,"link":"https://player.vimeo.com/external/360602969.sd.mp4?s=a77f0cd29f80c90a26ade4c5c8cdbf374287b28a\u0026profile_id=164\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":444963,"picture":"https://images.pexels.com/videos/2953633/pictures/preview-0.jpg","nr":0}]},{"id":1448735,"width":4096,"height":2160,"url":"https://www.pexels.com/video/video-of-forest-1448735/","image":"https://images.pexels.com/videos/1448735/free-video-1448735.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":32,"user":{"id":574687,"name":"RuvimMiksanskiy","url":"https://www.pexels.com/@digitech"},"video_files":[{"id":58649,"quality":"sd","file_type":"video/mp4","width":640,"height":338,"link":"https://player.vimeo.com/external/291648067.sd.mp4?s=7f9ee1f8ec1e5376027e4a6d1d05d5738b2fbb29\u0026profile_id=164\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":133236,"picture":"https://images.pexels.com/videos/1448735/pictures/preview-0.jpg","nr":0}]}],"total_results":33838,"page":2,"per_page":2,"prev_page":"https://api.pexels.com/v1/videos/popular/?max_duration=60\u0026min_duration=10\u0026min_height=2160\u0026min_width=4096\u0026page=1\u0026per_page=2","next_page":"https://api.pexels.com/v1/videos/popular/?max_duration=60\u0026min_duration=10\u0026min_height=2160\u0026min_width=4096\u0026page=3\u0026per_page=2"}`,
	},
}

var searchVideosTT = []struct {
	statusCode int
	opts       *Options
	params     *VideoSearchParams
	respBody   string
}{
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		&VideoSearchParams{Query: ""},
		``,
	},
	{
		http.StatusForbidden,
		&Options{APIKey: "invalid-APIKey"},
		&VideoSearchParams{Query: "City Lights"},
		`{"error": "Access to this API has been disallowed"}`,
	},
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		&VideoSearchParams{Query: "naasdfasdfasdfasdftasldkjfasdlkjfure"},
		`{"page":1,"per_page":1,"videos":[],"total_results":0,"url":"https://api-server.pexels.com/search/videos/naasdfasdfasdfasdftasldkjfasdlkjfure/"}18`,
	},
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		&VideoSearchParams{
			Query: "dogs",
			General: General{
				Locale:      "ukua",
				Orientation: "landscape",
				Size:        "medium",
			},
			Page:    4,
			PerPage: 3,
		},
		`{"videos":[{"id":3077158,"width":3840,"height":2160,"url":"https://www.pexels.com/uk-ua/video/3077158/","image":"https://images.pexels.com/videos/3077158/free-video-3077158.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":84,"user":{"id":102775,"name":"MagdaEhlers","url":"https://www.pexels.com/uk-ua/@magda-ehlers-pexels"},"video_files":[{"id":354864,"quality":"sd","file_type":"video/mp4","width":960,"height":540,"link":"https://player.vimeo.com/external/366179532.sd.mp4?s=fcfddee8ce5c72bcd9c6f3097ec48d530a0b8600\u0026profile_id=165\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":497782,"picture":"https://images.pexels.com/videos/3077158/pictures/preview-0.jpg","nr":0}]},{"id":3191251,"width":4096,"height":2160,"url":"https://www.pexels.com/uk-ua/video/3191251/","image":"https://images.pexels.com/videos/3191251/free-video-3191251.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":13,"user":{"id":1583460,"name":"Pressmaster","url":"https://www.pexels.com/uk-ua/@pressmaster"},"video_files":[{"id":253935,"quality":"sd","file_type":"video/mp4","width":960,"height":506,"link":"https://player.vimeo.com/external/371562597.sd.mp4?s=2d39d964e8af2dae0b0a0b303c04ea300c9be7fc\u0026profile_id=165\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":572772,"picture":"https://images.pexels.com/videos/3191251/pictures/preview-0.jpg","nr":0}]},{"id":2795691,"width":3840,"height":2160,"url":"https://www.pexels.com/uk-ua/video/2795691/","image":"https://images.pexels.com/videos/2795691/free-video-2795691.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":6,"user":{"id":1007758,"name":"EvgeniaKirpichnikova","url":"https://www.pexels.com/uk-ua/@evgenia-kirpichnikova-1007758"},"video_files":[{"id":157832,"quality":"hd","file_type":"video/mp4","width":1280,"height":720,"link":"https://player.vimeo.com/external/353554745.hd.mp4?s=e7aa9bc30aaa1618d0f1f076bec2ce45db61300d\u0026profile_id=174\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":385549,"picture":"https://images.pexels.com/videos/2795691/pictures/preview-0.jpg","nr":0}]}],"total_results":2879,"page":4,"per_page":3,"prev_page":"https://api.pexels.com/v1/videos/search/?locale=uk-UA\u0026orientation=landscape\u0026page=3\u0026per_page=3\u0026query=dogs\u0026size=medium","next_page":"https://api.pexels.com/v1/videos/search/?locale=uk-UA\u0026orientation=landscape\u0026page=5\u0026per_page=3\u0026query=dogs\u0026size=medium"}`,
	},
}

// photo.go test cases
var getPhotoTT = []struct {
	statusCode int
	opts       *Options
	ID         uint64
	respBody   string
}{
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		2014422,
		`{"id":2014422,"width":3024,"height":3024,"url":"https://www.pexels.com/photo/brown-rocks-during-golden-hour-2014422/","photographer":"Joey Farina","photographer_url":"https://www.pexels.com/@joey","photographer_id":680589,"avg_color":"#978E82","src":{"original":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg","large2x":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=2\u0026h=650\u0026w=940","large":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=650\u0026w=940","medium":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=350","small":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=130","portrait":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=1200\u0026w=800","landscape":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=627\u0026w=1200","tiny":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=1\u0026fit=crop\u0026h=200\u0026w=280"},"liked":false}`,
	},
	{
		http.StatusNotFound,
		&Options{APIKey: "testAPIKey"},
		0,
		`{"status":404,"error":"Not Found"}`,
	},
}

var searchPhotosTT = []struct {
	statusCode int
	opts       *Options
	params     *PhotoSearchParams
	respBody   string
}{
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		&PhotoSearchParams{Query: ""},
		``,
	},
	{
		http.StatusForbidden,
		&Options{APIKey: "invalid-APIKey"},
		&PhotoSearchParams{Query: "City Lights"},
		`{"error": "Access to this API has been disallowed"}`,
	},
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		&PhotoSearchParams{Query: "random-garbagepaosdjfpoijqwpoeijo!@#$%:;(*)"},
		`{"page":1,"per_page":1,"photos":[],"total_results":0,"url":"https://api-server.pexels.com/search/videos/random-garbagepaosdjfpoijqwpoeijo!@#$%:;(*)/"}`,
	},
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		&PhotoSearchParams{Query: "nature",
			General: General{
				Locale:      "uk-ua",
				Orientation: "landscape",
				Size:        "medium",
			},
			Color:   "red",
			Page:    1,
			PerPage: 1,
		},
		`{"page":1,"per_page":1,"photos":[{"id":130988,"width":6000,"height":3376,"url":"https://www.pexels.com/photo/red-flowers-130988/","photographer":"Mike","photographer_url":"https://www.pexels.com/@mikebirdy","photographer_id":20649,"avg_color":"#82251B","src":{"original":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg","large2x":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=2\u0026h=650\u0026w=940","large":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=650\u0026w=940","medium":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=350","small":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=130","portrait":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=1200\u0026w=800","landscape":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=627\u0026w=1200","tiny":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=1\u0026fit=crop\u0026h=200\u0026w=280"},"liked":false}],"total_results":8000,"next_page":"https://api.pexels.com/v1/search/?color=red\u0026page=2\u0026per_page=1\u0026query=nature\u0026size=medium"}`,
	},
}

var getCuratedPhotosTT = []struct {
	statusCode int
	opts       *Options
	params     *CuratedPhotosParams
	respBody   string
}{
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		&CuratedPhotosParams{Page: 2, PerPage: 5},
		`{"page":2,"per_page":5,"photos":[{"id":8536867,"width":4291,"height":5364,"url":"https://www.pexels.com/photo/silhouette-of-man-raising-his-hands-8536867/","photographer":"Trarete","photographer_url":"https://www.pexels.com/@trarete-73723862","photographer_id":73723862,"avg_color":"#405580","src":{"original":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg","large2x":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=2\u0026h=650\u0026w=940","large":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=650\u0026w=940","medium":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=350","small":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=130","portrait":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=1200\u0026w=800","landscape":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=627\u0026w=1200","tiny":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=1\u0026fit=crop\u0026h=200\u0026w=280"},"liked":false}],"total_results":8000,"next_page":"https://api.pexels.com/v1/curated/?page=3\u0026per_page=5","prev_page":"https://api.pexels.com/v1/curated/?page=1\u0026per_page=5"}`,
	},
}

// collection.go test cases
var getCollectionTT = []struct {
	statusCode int
	opts       *Options
	params     *CollectionMediaParams
	respBody   string
}{
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		&CollectionMediaParams{
			ID:      "g3aedup",
			Type:    "videos",
			Page:    1,
			PerPage: 3,
		},
		`{"id":"g3aedup","Videos":[{"id":5409152,"width":1080,"height":1920,"url":"https://www.pexels.com/video/learning-spanish-simple-minimalism-5409152/","image":"https://images.pexels.com/videos/5409152/foreign-language-home-study-home-work-learning-5409152.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":15,"user":{"id":2946790,"name":"OlyaKobruseva","url":"https://www.pexels.com/@olyakobruseva"},"video_files":[{"id":1373504,"quality":"sd","file_type":"video/mp4","width":360,"height":640,"link":"https://player.vimeo.com/external/460194477.sd.mp4?s=96856e4d6f24eaa7b075ddf9d1bbbe21dc26b7da\u0026profile_id=164\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":2786191,"picture":"https://images.pexels.com/videos/5409152/pictures/preview-0.jpg","nr":0}],"type":"Video"},{"id":2759451,"width":1920,"height":1080,"url":"https://www.pexels.com/video/an-old-video-footage-of-people-enjoying-their-day-in-a-beach-2759451/","image":"https://images.pexels.com/videos/2759451/free-video-2759451.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":9,"user":{"id":1092752,"name":"LifeOnSuper8","url":"https://www.pexels.com/@life-on-super-8-1092752"},"video_files":[{"id":151386,"quality":"hd","file_type":"video/mp4","width":1280,"height":720,"link":"https://player.vimeo.com/external/352010608.hd.mp4?s=16e946fe241f74b03c68b7c73fc44d5e78ec0b10\u0026profile_id=174\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":370599,"picture":"https://images.pexels.com/videos/2759451/pictures/preview-0.jpg","nr":0}],"type":"Video"},{"id":2034302,"width":3840,"height":2160,"url":"https://www.pexels.com/video/an-old-phonograph-2034302/","image":"https://images.pexels.com/videos/2034302/free-video-2034302.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":19,"user":{"id":1054894,"name":"JakubGorajek","url":"https://www.pexels.com/@jakub-gorajek-1054894"},"video_files":[{"id":84782,"quality":"hd","file_type":"video/mp4","width":3840,"height":2160,"link":"https://player.vimeo.com/external/325039010.hd.mp4?s=0b149cb826b5811e593007d992715ebe0649290c\u0026profile_id=172\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":208651,"picture":"https://images.pexels.com/videos/2034302/pictures/preview-0.jpg","nr":0}],"type":"Video"}],"Photos":null,"total_results":3,"page":1,"per_page":15,"prev_page":"","next_page":""}`,
	},
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		&CollectionMediaParams{
			ID:      "c10vohh",
			Type:    "photos",
			Page:    1,
			PerPage: 3,
		},
		`{"page":1,"per_page":15,"media":[{"type":"Photo","id":5391353,"width":3456,"height":3456,"url":"https://www.pexels.com/photo/close-up-photography-of-green-plant-5391353/","photographer":"Antonio Prado","photographer_url":"https://www.pexels.com/@antonio-prado-1050855","photographer_id":1050855,"avg_color":"#373F38","src":{},"liked":false},{"type":"Photo","id":7120688,"width":3238,"height":4046,"url":"https://www.pexels.com/photo/woman-in-white-and-red-floral-shirt-standing-beside-white-flowers-7120688/","photographer":"TRAVELBLOG","photographer_url":"https://www.pexels.com/@travelblog-26954066","photographer_id":26954066,"avg_color":"#959486","src":{},"liked":false},{"type":"Photo","id":8043220,"width":2688,"height":4032,"url":"https://www.pexels.com/photo/woman-in-black-spaghetti-strap-top-8043220/","photographer":"Eman Genatilan","photographer_url":"https://www.pexels.com/@eman-genatilan-3459781","photographer_id":3459781,"avg_color":"#707360","src":{},"liked":false}],"total_results":3,"id":"c10vohh"}`,
	},
	{
		http.StatusNotFound,
		&Options{APIKey: "testAPIKey"},
		&CollectionMediaParams{
			ID: "",
		},
		`{"status":404,"error":"Not Found"}`,
	},
}

var getCollectionsTT = []struct {
	statusCode int
	opts       *Options
	respBody   string
}{
	{
		http.StatusOK,
		&Options{APIKey: "testAPIKey"},
		`{"page":1,"per_page":1,"collections":[{"id":"c10vohh","title":"Test API Collection","description":null,"private":false,"media_count":3,"photos_count":3,"videos_count":0}],"total_results":1}`,
	},
	{
		http.StatusForbidden,
		&Options{APIKey: "invalid-APIKey"},
		`{"error": "Access to this API has been disallowed"}`,
	},
}
