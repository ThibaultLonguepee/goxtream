package goxtream

import (
	"github.com/thibaultlonguepee/goxtream/internal"
	"github.com/thibaultlonguepee/goxtream/internal/mappers"
	"github.com/thibaultlonguepee/goxtream/models"
)

type Source struct {
	url      string
	username string
	password string
}

func NewSource(url, username, password string) *Source {
	return &Source{
		url:      url,
		username: username,
		password: password,
	}
}

func (src *Source) Authenticate() (*models.Authentication, error) {
	return internal.TryGetParsed(mappers.ParseAuthentication, authenticationUrl, src.url, src.username, src.password)
}

func (src *Source) GetLiveCategories() ([]*models.Category, error) {
	return internal.TryGetParsed(mappers.ParseCategories, liveCategoriesUrl, src.url, src.username, src.password)
}

func (src *Source) GetVodCategories() ([]*models.Category, error) {
	return internal.TryGetParsed(mappers.ParseCategories, vodCategoriesUrl, src.url, src.username, src.password)
}

func (src *Source) GetSeriesCategories() ([]*models.Category, error) {
	return internal.TryGetParsed(mappers.ParseCategories, showCategoriesUrl, src.url, src.username, src.password)
}

func (src *Source) GetLiveStreams() ([]*models.LiveStream, error) {
	return internal.TryGetParsed(mappers.ParseLiveStreams, liveStreamsUrl, src.url, src.username, src.password)
}

func (src *Source) GetCategoryLiveStreams(category int) ([]*models.LiveStream, error) {
	return internal.TryGetParsed(mappers.ParseLiveStreams, categoryLiveStreamsUrl, src.url, src.username, src.password, category)
}

func (src *Source) GetVodStreams() ([]*models.Vod, error) {
	return internal.TryGetParsed(mappers.ParseVods, vodStreamsUrl, src.url, src.username, src.password)
}

func (src *Source) GetCategoryVodStreams(category int) ([]*models.Vod, error) {
	return internal.TryGetParsed(mappers.ParseVods, categoryVodStreamsUrl, src.url, src.username, src.password, category)
}

func (src *Source) GetVodDetails(vod int) (*models.VodDetails, error) {
	return internal.TryGetParsed(mappers.ParseVodDetails, vodDetailsUrl, src.url, src.username, src.password, vod)
}

func (src *Source) GetShows() ([]*models.Show, error) {
	return internal.TryGetParsed(mappers.ParseShows, showsUrl, src.url, src.username, src.password)
}

func (src *Source) GetCategoryShows(category int) ([]*models.Show, error) {
	return internal.TryGetParsed(mappers.ParseShows, categoryShowsUrl, src.url, src.username, src.password, category)
}

func (src *Source) GetShowDetails(show int) (*models.ShowDetails, error) {
	return internal.TryGetParsed(mappers.ParseShowDetails, showDetailsUrl, src.url, src.username, src.password, show)
}
