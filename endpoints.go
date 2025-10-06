package goxtream

const (
	AuthenticationUrl   = "%s/player_api.php?username=%s&password=%s"
	LiveCategoriesUrl   = AuthenticationUrl + "&action=get_live_categories"
	VodCategoriesUrl    = AuthenticationUrl + "&action=get_vod_categories"
	SeriesCategoriesUrl = AuthenticationUrl + "&action=get_series_categories"
)
