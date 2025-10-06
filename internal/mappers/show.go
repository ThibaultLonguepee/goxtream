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
