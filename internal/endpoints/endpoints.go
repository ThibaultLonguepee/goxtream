package endpoints

const (
	AuthenticationUrl = "%v/player_api.php?username=%v&password=%v"

	LiveCategoriesUrl      string = AuthenticationUrl + "&action=get_live_categories"
	LiveStreamsUrl         string = AuthenticationUrl + "&action=get_live_streams"
	CategoryLiveStreamsUrl string = AuthenticationUrl + "&action=get_live_streams&category_id=%v"
	LiveStreamEpgUrl       string = AuthenticationUrl + "&action=get_short_epg&stream_id=%v"
	LiveStreamUri          string = "%v/live/%v/%v/%v.%v"

	VodCategoriesUrl      string = AuthenticationUrl + "&action=get_vod_categories"
	VodStreamsUrl         string = AuthenticationUrl + "&action=get_vod_streams"
	CategoryVodStreamsUrl string = AuthenticationUrl + "&action=get_vod_streams&category_id=%v"
	VodDetailsUrl         string = AuthenticationUrl + "&action=get_vod_info&vod_id=%v"
	VodUri                string = "%v/movie/%v/%v/%v.%v"

	ShowCategoriesUrl string = AuthenticationUrl + "&action=get_series_categories"
	ShowsUrl          string = AuthenticationUrl + "&action=get_series"
	CategoryShowsUrl  string = AuthenticationUrl + "&action=get_series&category_id=%v"
	ShowDetailsUrl    string = AuthenticationUrl + "&action=get_series_info&series_id=%v"
	EpisodeUri        string = "%v/series/%v/%v/%v.%v"
)
