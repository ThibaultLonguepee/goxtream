package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/thibaultlonguepee/goxtream/internal/dtos"
)

type LiveStream struct {
	Id           int
	Number       int
	Name         string
	Icon         string
	EpgChannel   string
	AdditionDate time.Time
	IsAdult      bool
	CategoryId   int
}

func ParseLiveStream(src dtos.LiveStreamResult) (*LiveStream, error) {
	added, err := strconv.Atoi(src.Added)
	if err != nil {
		return nil, errors.New("wrong format for live stream addition date")
	}

	category, err := strconv.Atoi(src.CategoryId)
	if err != nil {
		return nil, errors.New("wrong format for live stream category id")
	}

	return &LiveStream{
		Id:           src.StreamId,
		Number:       src.Num,
		Name:         src.Name,
		Icon:         src.StreamIcon,
		EpgChannel:   src.EpgChannelId,
		AdditionDate: time.Unix(int64(added), 0),
		IsAdult:      src.IsAdult != "0",
		CategoryId:   category,
	}, nil
}

func ParseLiveStreams(srcs []dtos.LiveStreamResult) ([]*LiveStream, error) {
	streams := make([]*LiveStream, 0)
	for _, s := range srcs {
		stream, err := ParseLiveStream(s)
		if err != nil {
			return nil, err
		}
		streams = append(streams, stream)
	}
	return streams, nil
}
