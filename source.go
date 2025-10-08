package goxtream

import (
	"fmt"

	"github.com/thibaultlonguepee/goxtream/internal"
	"github.com/thibaultlonguepee/goxtream/internal/endpoints"
	"github.com/thibaultlonguepee/goxtream/internal/mappers"
	"github.com/thibaultlonguepee/goxtream/models"
)

type Source struct {
	Id            int
	Name          string
	Account       *models.Account
	authenticated bool
}

func NewSource(id int, name string) *Source {
	return &Source{
		Id:      id,
		Name:    name,
		Account: nil,
	}
}

func (src *Source) Authenticated() bool { return src.authenticated }
func (src *Source) Username() string    { return src.Account.User.Name }
func (src *Source) Password() string    { return src.Account.User.Password }
func (src *Source) Url() string {
	return fmt.Sprintf("%v://%v:%v", src.Account.Server.Protocol, src.Account.Server.Url, src.Account.Server.Port)
}

func (src *Source) Authenticate(url, username, password string) (*models.Account, error) {
	account, err := internal.Get(mappers.ParseAccount, endpoints.AuthenticationUrl, url, username, password)
	if err != nil {
		return nil, fmt.Errorf("could not authenticate user %v: %v", username, err)
	}
	src.Account = account
	src.authenticated = true
	return account, nil
}

func (src *Source) GetLiveCategories() ([]*models.Category, error) {
	return internal.GetWithCredentials(mappers.ParseCategories, endpoints.LiveCategoriesUrl, src)
}

func (src *Source) GetLiveStreams() ([]*models.LiveStream, error) {
	return internal.GetWithCredentials(mappers.ParseLiveStreams, endpoints.LiveStreamsUrl, src)
}

func (src *Source) GetCategoryLiveStreams(category int) ([]*models.LiveStream, error) {
	return internal.GetWithCredentials(mappers.ParseLiveStreams, endpoints.CategoryLiveStreamsUrl, src, category)
}

func (src *Source) GetLiveStreamUrls(stream int) []string {
	links := make([]string, 0)
	for _, format := range src.Account.User.Formats {
		link := fmt.Sprintf(endpoints.LiveStreamUri, src.Account.Server.GetUrl(), src.Username(), src.Password(), stream, format)
		links = append(links, link)
	}
	return links
}

// Vods

func (src *Source) GetVodCategories() ([]*models.Category, error) {
	return internal.GetWithCredentials(mappers.ParseCategories, endpoints.VodCategoriesUrl, src)
}

func (src *Source) GetVods() ([]*models.Vod, error) {
	return internal.GetWithCredentials(mappers.ParseVods, endpoints.VodStreamsUrl, src)
}

func (src *Source) GetCategoryVods(category int) ([]*models.Vod, error) {
	return internal.GetWithCredentials(mappers.ParseVods, endpoints.CategoryVodStreamsUrl, src, category)
}

func (src *Source) GetVodDetails(vod int) (*models.VodDetails, error) {
	return internal.GetWithCredentials(mappers.ParseVodDetails, endpoints.VodDetailsUrl, src, vod)
}

func (src *Source) GetVodUrl(vod int) (string, error) {
	details, err := src.GetVodDetails(vod)
	if err != nil {
		return "", err
	}
	return details.GetStreamingLink(src.Account), nil
}

// Shows

func (src *Source) GetShowCategories() ([]*models.Category, error) {
	return internal.GetWithCredentials(mappers.ParseCategories, endpoints.ShowCategoriesUrl, src)
}

func (src *Source) GetShows() ([]*models.Show, error) {
	return internal.GetWithCredentials(mappers.ParseShows, endpoints.ShowsUrl, src)
}

func (src *Source) GetCategoryShows(category int) ([]*models.Show, error) {
	return internal.GetWithCredentials(mappers.ParseShows, endpoints.CategoryShowsUrl, src, category)
}

func (src *Source) GetShowDetails(show int) (*models.ShowDetails, error) {
	return internal.GetWithCredentials(mappers.ParseShowDetails, endpoints.ShowDetailsUrl, src, show)
}

func (src *Source) GetEpisodeUrl(show, season, episode int) (string, error) {
	details, err := src.GetShowDetails(show)
	if err != nil {
		return "", err
	}

	s := details.GetSeason(season)
	if s == nil {
		return "", fmt.Errorf("show '%v' doesn't have a season %v", details.Name, season)
	}

	e := s.GetEpisode(episode)
	if e == nil {
		return "", fmt.Errorf("show's '%v' season %v doesn't have an %v", details.Name, season, episode)
	}

	return e.GetStreamingLink(src.Account), nil
}
