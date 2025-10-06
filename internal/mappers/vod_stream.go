package mappers

import (
	"errors"
	"strconv"
	"time"

	"github.com/thibaultlonguepee/goxtream/internal/dtos"
	"github.com/thibaultlonguepee/goxtream/models"
)

func ParseVodStream(src dtos.VodStreamResult) (*models.VodStream, error) {
	added, err := strconv.Atoi(src.Added)
	if err != nil {
		return nil, errors.New("wrong format for Vod stream addition date")
	}

	category, err := strconv.Atoi(src.CategoryId)
	if err != nil {
		return nil, errors.New("wrong format for Vod stream category id")
	}

	return &models.VodStream{
		Id:           src.StreamId,
		Number:       src.Num,
		Name:         src.Name,
		Icon:         src.StreamIcon,
		Rating:       float64(src.Rating),
		AdditionDate: time.Unix(int64(added), 0),
		IsAdult:      src.IsAdult != "0",
		Extension:    src.ContainerExtension,
		CategoryId:   category,
	}, nil
}

func ParseVodStreams(srcs []dtos.VodStreamResult) ([]*models.VodStream, error) {
	streams := make([]*models.VodStream, 0)
	for _, s := range srcs {
		stream, err := ParseVodStream(s)
		if err != nil {
			return nil, err
		}
		streams = append(streams, stream)
	}
	return streams, nil
}
