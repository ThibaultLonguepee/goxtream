package dtos

type SeriesResult struct {
	Num            int    `json:"num"`
	Name           string `json:"name"`
	SeriesId       int    `json:"series_id"`
	Cover          string `json:"cover"`
	Plot           string `json:"plot"`
	Cast           string `json:"cast"`
	Director       string `json:"director"`
	Genre          string `json:"genre"`
	ReleaseDate    string `json:"releaseDate"`
	LastModified   string `json:"last_modified"`
	Rating         Rating `json:"rating"`
	Rating5Based   Rating `json:"rating_5based"`
	BackdropPath   Paths  `json:"backdrop_path"`
	YoutubeTrailer string `json:"youtube_trailer"`
	EpisodeRunTime string `json:"episode_run_time"`
	CategoryId     string `json:"category_id"`
}

type SeriesDetailsResult struct {
	Seasons  []SeriesSeason          `json:"seasons"`
	Info     SeriesInfo              `json:"info"`
	Episodes map[int][]SeriesEpisode `json:"episodes"`
}

type SeriesSeason struct {
	AirDate      string `json:"air_date"`
	EpisodeCount int    `json:"episode_count"`
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	SeasonNumber int    `json:"season_number"`
	VoteAverage  Rating `json:"vote_average"`
	Cover        string `json:"cover"`
	CoverBig     string `json:"cover_big"`
}

type SeriesInfo struct {
	Name           string `json:"name"`
	Cover          string `json:"cover"`
	Plot           string `json:"plot"`
	Cast           string `json:"cast"`
	Director       string `json:"director"`
	Genre          string `json:"genre"`
	ReleaseDate    string `json:"releaseDate"`
	LastModified   string `json:"last_modified"`
	Rating         Rating `json:"rating"`
	Rating5Based   Rating `json:"rating_5based"`
	BackdropPath   Paths  `json:"backdrop_path"`
	YoutubeTrailer string `json:"youtube_trailer"`
	EpisodeRunTime string `json:"episode_run_time"`
	CategoryId     string `json:"category_id"`
}

type SeriesEpisode struct {
	Id                 string      `json:"id"`
	EpisodeNum         int         `json:"episode_num"`
	Title              string      `json:"title"`
	ContainerExtension string      `json:"container_extension"`
	Info               EpisodeInfo `json:"info"`
	CustomSid          string      `json:"custom_sid"`
	Added              string      `json:"added"`
	Season             int         `json:"season"`
	DirectSource       string      `json:"direct_source"`
}

type EpisodeInfo struct {
	MovieImage   string `json:"movie_image"`
	Plot         string `json:"plot"`
	ReleaseDate  string `json:"releaseDate"`
	Rating       Rating `json:"rating"`
	Name         string `json:"name"`
	DurationSecs int    `json:"duration_secs"`
	Duration     string `json:"duration"`
	// TODO: Video        EpisodeVideoInfo `json:"video"`
	// TODO: Audio        EpisodeAudioInfo `json:"audio"`
	Bitrate int `json:"bitrate"`
}
