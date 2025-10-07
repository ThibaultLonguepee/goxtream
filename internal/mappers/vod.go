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

func ParseVod(src dtos.VodStreamResult) (*models.Vod, error) {
	added, err := strconv.Atoi(src.Added)
	if err != nil {
		return nil, errors.New("wrong format for Vod stream addition date")
	}

	category, err := strconv.Atoi(src.CategoryId)
	if err != nil {
		return nil, errors.New("wrong format for Vod stream category id")
	}

	return &models.Vod{
		Id:           src.StreamId,
		Number:       src.Num,
		Name:         src.Name,
		Thumbnail:    src.StreamIcon,
		Rating:       float64(src.Rating),
		AdditionDate: time.Unix(int64(added), 0),
		IsAdult:      src.IsAdult != "0",
		Extension:    src.ContainerExtension,
		Category:     category,
	}, nil
}

func ParseVods(srcs []dtos.VodStreamResult) ([]*models.Vod, error) {
	streams := make([]*models.Vod, 0)
	for _, s := range srcs {
		stream, err := ParseVod(s)
		if err != nil {
			return nil, err
		}
		streams = append(streams, stream)
	}
	return streams, nil
}

func ParseVodDetails(src dtos.VodDetailsResult) (*models.VodDetails, error) {
	releaseDate, _ := time.Parse(time.DateOnly, src.Info.ReleaseDate)

	added, err := strconv.Atoi(src.MovieData.Added)
	if err != nil {
		return nil, fmt.Errorf("wrong format for vod details addition date: %v", src.MovieData.Added)
	}

	category, err := strconv.Atoi(src.MovieData.CategoryId)
	if err != nil {
		return nil, fmt.Errorf("wrong format for vod details category id: %v", src.MovieData.CategoryId)
	}

	return &models.VodDetails{
		Id:           src.MovieData.StreamId,
		Name:         src.Info.Name,
		OriginalName: src.Info.OName,
		Plot:         src.Info.Plot,
		Genre:        src.Info.Genre,
		Cast:         strings.Split(src.Info.Cast, ", "),
		Directors:    strings.Split(src.Info.Director, ", "),
		Duration:     src.Info.DurationSecs,
		RealeaseDate: releaseDate,
		Country:      src.Info.Country,
		Age:          src.Info.Age,
		MpaaRating:   src.Info.RatingMpaa,
		Rating:       float64(src.Info.Rating),
		RatingCount:  src.Info.RatingCountKinopoisk,
		AdditionDate: time.Unix(int64(added), 0),
		Cover:        src.Info.CoverBig,
		Backdrops:    src.Info.BackdropPath,
		Bitrate:      src.Info.Bitrate,
		Extension:    src.MovieData.ContainerExtension,
		Category:     category,
	}, nil
}
