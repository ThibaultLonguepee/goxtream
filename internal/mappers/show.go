package mappers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/thibaultlonguepee/goxtream/internal/dtos"
	"github.com/thibaultlonguepee/goxtream/models"
)

func ParseShow(src dtos.SeriesResult) (*models.Show, error) {
	release, _ := time.Parse(time.DateOnly, src.ReleaseDate)

	modification, err := strconv.Atoi(src.LastModified)
	if err != nil {
		return nil, errors.New("wrong format for series modification date")
	}

	runtime, err := strconv.Atoi(src.EpisodeRunTime)
	if err != nil {
		return nil, errors.New("wrong format for series run time")
	}

	category, err := strconv.Atoi(src.CategoryId)
	if err != nil {
		return nil, errors.New("wrong format for series category id")
	}

	return &models.Show{
		Id:               src.SeriesId,
		Number:           src.Num,
		Name:             src.Name,
		Cover:            src.Cover,
		Plot:             src.Plot,
		Cast:             strings.Split(src.Cast, ", "),
		Directors:        strings.Split(src.Director, ", "),
		Genre:            src.Genre,
		ReleaseDate:      release,
		LastModification: time.Unix(int64(modification), 0),
		Rating:           float64(src.Rating),
		Backdrops:        src.BackdropPath,
		Trailer:          fmt.Sprintf("https://www.youtube.com/watch?v=%v", src.YoutubeTrailer),
		RunTime:          runtime,
		Category:         category,
	}, nil
}

func ParseShows(srcs []dtos.SeriesResult) ([]*models.Show, error) {
	shows := make([]*models.Show, 0)
	for _, s := range srcs {
		show, err := ParseShow(s)
		if err != nil {
			return nil, err
		}
		shows = append(shows, show)
	}
	return shows, nil
}

func ParseShowDetails(src dtos.SeriesDetailsResult) (*models.ShowDetails, error) {
	release, _ := time.Parse(time.DateOnly, src.Info.ReleaseDate)

	seasons, err := parseSeasons(src.Seasons, src.Episodes)
	if err != nil {
		return nil, err
	}

	modification, err := strconv.Atoi(src.Info.LastModified)
	if err != nil {
		return nil, errors.New("wrong format for series modification date")
	}

	runtime, err := strconv.Atoi(src.Info.EpisodeRunTime)
	if err != nil {
		return nil, errors.New("wrong format for series run time")
	}

	category, err := strconv.Atoi(src.Info.CategoryId)
	if err != nil {
		return nil, errors.New("wrong format for series category id")
	}

	return &models.ShowDetails{
		Name:             src.Info.Name,
		Cover:            src.Info.Cover,
		Plot:             src.Info.Plot,
		Cast:             strings.Split(src.Info.Cast, ", "),
		Directors:        strings.Split(src.Info.Director, ", "),
		Genre:            src.Info.Genre,
		ReleaseDate:      release,
		LastModification: time.Unix(int64(modification), 0),
		Rating:           float64(src.Info.Rating),
		Backdrops:        src.Info.BackdropPath,
		Trailer:          fmt.Sprintf("https://www.youtube.com/watch?v=%v", src.Info.YoutubeTrailer),
		RunTime:          runtime,
		Category:         category,
		Seasons:          seasons,
	}, nil
}

func parseSeason(src dtos.SeriesSeason, eps []dtos.SeriesEpisode) (*models.Season, error) {
	air, _ := time.Parse(time.DateOnly, src.AirDate)

	episodes, err := parseEpisodes(eps)
	if err != nil {
		return nil, err
	}

	return &models.Season{
		Id:       src.Id,
		Number:   src.SeasonNumber,
		Name:     src.Name,
		Overview: src.Overview,
		AirDate:  air,
		Rating:   float64(src.VoteAverage),
		Cover:    src.CoverBig,
		Episodes: episodes,
	}, nil
}

func parseSeasons(srcs []dtos.SeriesSeason, eps map[int][]dtos.SeriesEpisode) ([]*models.Season, error) {
	seasons := make([]*models.Season, 0)
	for _, s := range srcs {
		season, err := parseSeason(s, eps[s.SeasonNumber])
		if err != nil {
			return nil, err
		}
		seasons = append(seasons, season)
	}
	return seasons, nil
}

func parseEpisode(src dtos.SeriesEpisode) (*models.Episode, error) {
	release, _ := time.Parse(time.DateOnly, src.Info.ReleaseDate)

	id, err := strconv.Atoi(src.Id)
	if err != nil {
		return nil, fmt.Errorf("wrong format for episode id: %v", id)
	}

	added, err := strconv.Atoi(src.Added)
	if err != nil {
		return nil, fmt.Errorf("wrong format for episode addition date: %v", id)
	}

	return &models.Episode{
		Id:           id,
		Number:       src.EpisodeNum,
		Season:       src.Season,
		Title:        src.Title,
		Plot:         src.Info.Plot,
		RealeaseDate: release,
		Rating:       float64(src.Info.Rating),
		Duration:     src.Info.DurationSecs,
		AdditionDate: time.Unix(int64(added), 0),
		Extension:    src.ContainerExtension,
	}, nil
}

func parseEpisodes(srcs []dtos.SeriesEpisode) ([]*models.Episode, error) {
	episodes := make([]*models.Episode, 0)
	for _, s := range srcs {
		episode, err := parseEpisode(s)
		if err != nil {
			return nil, err
		}
		episodes = append(episodes, episode)
	}
	return episodes, nil
}
