package goxtream

const (
	AuthenticationUrl = "%v/player_api.php?username=%v&password=%v"

	LiveCategoriesUrl   string = AuthenticationUrl + "&action=get_live_categories"
	VodCategoriesUrl    string = AuthenticationUrl + "&action=get_vod_categories"
	SeriesCategoriesUrl string = AuthenticationUrl + "&action=get_series_categories"

	LiveStreamsUrl         string = AuthenticationUrl + "&action=get_live_streams"
	VodStreamsUrl          string = AuthenticationUrl + "&action=get_vod_streams"
	CategoryLiveStreamsUrl string = AuthenticationUrl + "&action=get_live_streams&category_id=%v"
	CategoryVodStreamsUrl  string = AuthenticationUrl + "&action=get_vod_streams&category_id=%v"
)
