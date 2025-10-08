package models

import (
	"fmt"
	"time"

	"github.com/thibaultlonguepee/goxtream/internal/endpoints"
)

type Show struct {
	Id               int
	Number           int
	Name             string
	Cover            string
	Plot             string
	Cast             []string
	Directors        []string
	Genre            string
	ReleaseDate      time.Time
	LastModification time.Time
	Rating           float64
	Backdrops        []string
	Trailer          string
	RunTime          int
	Category         int
}

type ShowDetails struct {
	Name             string
	Cover            string
	Plot             string
	Cast             []string
	Directors        []string
	Genre            string
	ReleaseDate      time.Time
	LastModification time.Time
	Rating           float64
	Backdrops        []string
	Trailer          string
	RunTime          int
	Category         int
	Seasons          []*Season
}

func (show *ShowDetails) GetSeason(number int) *Season {
	for _, season := range show.Seasons {
		if season.Number == number {
			return season
		}
	}
	return nil
}

type Season struct {
	Id       int
	Number   int
	Name     string
	Overview string
	AirDate  time.Time
	Rating   float64
	Cover    string
	Episodes []*Episode
}

func (season *Season) GetEpisode(number int) *Episode {
	for _, episode := range season.Episodes {
		if episode.Number == number {
			return episode
		}
	}
	return nil
}

type Episode struct {
	Id           int
	Number       int
	Season       int
	Title        string
	Plot         string
	RealeaseDate time.Time
	Rating       float64
	Duration     int
	AdditionDate time.Time
	Extension    string
}

func (vod *Episode) GetStreamingLink(account *Account) string {
	return fmt.Sprintf(endpoints.EpisodeUri, account.Server.GetUrl(), account.User.Name, account.User.Password, vod.Id, vod.Extension)
}
