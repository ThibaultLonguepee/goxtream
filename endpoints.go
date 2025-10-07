package goxtream

const (
	authenticationUrl = "%v/player_api.php?username=%v&password=%v"

	liveCategoriesUrl      string = authenticationUrl + "&action=get_live_categories"
	liveStreamsUrl         string = authenticationUrl + "&action=get_live_streams"
	categoryLiveStreamsUrl string = authenticationUrl + "&action=get_live_streams&category_id=%v"

	vodCategoriesUrl      string = authenticationUrl + "&action=get_vod_categories"
	vodStreamsUrl         string = authenticationUrl + "&action=get_vod_streams"
	categoryVodStreamsUrl string = authenticationUrl + "&action=get_vod_streams&category_id=%v"
	vodDetailsUrl         string = authenticationUrl + "&action=get_vod_info&vod_id=%v"

	showCategoriesUrl string = authenticationUrl + "&action=get_series_categories"
	showsUrl          string = authenticationUrl + "&action=get_series"
	categoryShowsUrl  string = authenticationUrl + "&action=get_series&category_id=%v"
	showDetailsUrl    string = authenticationUrl + "&action=get_series_info&series_id=%v"
)
