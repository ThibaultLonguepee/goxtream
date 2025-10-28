package mappers

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/thibaultlonguepee/goxtream/internal/dtos"
	"github.com/thibaultlonguepee/goxtream/models"
)

func ParseEpg(src dtos.EpgResult) (*models.Epg, error) {
	id, err := strconv.Atoi(src.Id)
	if err != nil {
		return nil, fmt.Errorf("wrong format for epg id: %v", src.Id)
	}

	title, err := base64.StdEncoding.DecodeString(src.Title)
	if err != nil {
		return nil, fmt.Errorf("could not decode epg title: %v", src.Title)
	}

	description, err := base64.StdEncoding.DecodeString(src.Description)
	if err != nil {
		return nil, fmt.Errorf("could not decode epg description: %v", src.Description)
	}

	start, err := strconv.Atoi(src.StartTimestamp)
	if err != nil {
		return nil, fmt.Errorf("wrong format for epg start timestamp: %v", src.StartTimestamp)
	}

	end, err := strconv.Atoi(src.StopTimestamp)
	if err != nil {
		return nil, fmt.Errorf("wrong format for epg end timestamp: %v", src.StopTimestamp)
	}

	return &models.Epg{
		Id:          id,
		Title:       string(title),
		Description: string(description),
		Start:       time.Unix(int64(start), 0),
		End:         time.Unix(int64(end), 0),
		Language:    src.Lang,
	}, nil
}

func ParseEpgs(src dtos.EpgListingResult) ([]*models.Epg, error) {
	epgs := make([]*models.Epg, 0)
	for _, e := range src.EpgListings {
		epg, err := ParseEpg(e)
		if err != nil {
			return nil, err
		}
		epgs = append(epgs, epg)
	}
	return epgs, nil
}
