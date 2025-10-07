package dtos

type VodStreamResult struct {
	Num                int    `json:"num"`
	Name               string `json:"name"`
	StreamType         string `json:"stream_type"`
	StreamId           int    `json:"stream_id"`
	StreamIcon         string `json:"stream_icon"`
	Rating             Rating `json:"rating"`
	Rating5Based       Rating `json:"rating_5based"`
	Added              string `json:"added"`
	IsAdult            string `json:"is_adult"`
	CategoryId         string `json:"category_id"`
	ContainerExtension string `json:"container_extension"`
	CustomSid          string `json:"custom_sid"`
	DirectSource       string `json:"direct_source"`
}

type VodDetailsResult struct {
	Info      VodDetailsInfo `json:"info"`
	MovieData VodMoveData    `json:"movie_data"`
}

type VodDetailsInfo struct {
	MovieImage     string   `json:"movie_image"`
	TmbdId         string   `json:"tmbd_id"`
	Backdrop       string   `json:"backdrop"`
	YoutubeTrailer string   `json:"youtube_trailer"`
	Genre          string   `json:"genre"`
	Plot           string   `json:"plot"`
	Cast           string   `json:"cast"`
	Rating         Rating   `json:"rating"`
	Director       string   `json:"director"`
	ReleaseDate    string   `json:"releaseDate"`
	BackdropPath   []string `json:"backdrop_path"`
	DurationSecs   int      `json:"duration_secs"`
	Duration       string   `json:"duration"`
	// Video			string `json:"video"`
	// Audio             string `json:"audio"`
	Bitrate              int    `json:"bitrate"`
	KinopoiskUrl         string `json:"kinopoisk_url"`
	Name                 string `json:"name"`
	OName                string `json:"o_name"`
	CoverBig             string `json:"cover_big"`
	EpisodeRunTime       int    `json:"episode_run_time"`
	Actors               string `json:"actors"`
	Description          string `json:"description"`
	Age                  string `json:"age"`
	RatingMpaa           string `json:"rating_mpaa"`
	RatingCountKinopoisk int    `json:"rating_count_kinopoisk"`
	Country              string `json:"country"`
}

type VodMoveData struct {
	StreamId           int    `json:"stream_id"`
	Name               string `json:"name"`
	Added              string `json:"added"`
	CategoryId         string `json:"category_id"`
	ContainerExtension string `json:"container_extension"`
	CustomSid          string `json:"custom_sid"`
	DirectSource       string `json:"direct_source"`
}
