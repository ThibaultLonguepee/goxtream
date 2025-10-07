package models

import (
	"fmt"
	"time"

	"github.com/thibaultlonguepee/goxtream/internal/endpoints"
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

func (live *LiveStream) GetStreamingLinks(account *Account) []string {
	links := make([]string, 0)
	for _, format := range account.User.Formats {
		link := fmt.Sprintf(endpoints.LiveStreamUri, account.Server.Url, account.User.Name, account.User.Password, live.Id, format)
		links = append(links, link)
	}
	return links
}
