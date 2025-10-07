package goxtream

import (
	"github.com/thibaultlonguepee/goxtream/internal"
	"github.com/thibaultlonguepee/goxtream/internal/endpoints"
	"github.com/thibaultlonguepee/goxtream/internal/mappers"
	"github.com/thibaultlonguepee/goxtream/models"
)

type Source struct {
	Id          int
	Name        string
	Credentials *models.Credentials
}

func NewSource(id int, name, url, username, password string) *Source {
	return &Source{
		Id:   id,
		Name: name,
		Credentials: &models.Credentials{
			Url:      url,
			Username: username,
			Password: password,
		},
	}
}

func (src *Source) Authenticate() (*models.Account, error) {
	return internal.TryGetParsed(mappers.ParseAccount, endpoints.AuthenticationUrl, src.Credentials)
}

func (src *Source) GetLiveCategories() ([]*models.Category, error) {
	return internal.TryGetParsed(mappers.ParseCategories, endpoints.LiveCategoriesUrl, src.Credentials)
}

func (src *Source) GetLiveStreams() ([]*models.LiveStream, error) {
	return internal.TryGetParsed(mappers.ParseLiveStreams, endpoints.LiveStreamsUrl, src.Credentials)
}

func (src *Source) GetCategoryLiveStreams(category int) ([]*models.LiveStream, error) {
	return internal.TryGetParsed(mappers.ParseLiveStreams, endpoints.CategoryLiveStreamsUrl, src.Credentials, category)
}

// Vods

func (src *Source) GetVodCategories() ([]*models.Category, error) {
	return internal.TryGetParsed(mappers.ParseCategories, endpoints.VodCategoriesUrl, src.Credentials)
}

func (src *Source) GetVods() ([]*models.Vod, error) {
	return internal.TryGetParsed(mappers.ParseVods, endpoints.VodStreamsUrl, src.Credentials)
}

func (src *Source) GetCategoryVods(category int) ([]*models.Vod, error) {
	return internal.TryGetParsed(mappers.ParseVods, endpoints.CategoryVodStreamsUrl, src.Credentials, category)
}

func (src *Source) GetVodDetails(vod int) (*models.VodDetails, error) {
	return internal.TryGetParsed(mappers.ParseVodDetails, endpoints.VodDetailsUrl, src.Credentials, vod)
}

// Shows

func (src *Source) GetShowCategories() ([]*models.Category, error) {
	return internal.TryGetParsed(mappers.ParseCategories, endpoints.ShowCategoriesUrl, src.Credentials)
}

func (src *Source) GetShows() ([]*models.Show, error) {
	return internal.TryGetParsed(mappers.ParseShows, endpoints.ShowsUrl, src.Credentials)
}

func (src *Source) GetCategoryShows(category int) ([]*models.Show, error) {
	return internal.TryGetParsed(mappers.ParseShows, endpoints.CategoryShowsUrl, src.Credentials, category)
}

func (src *Source) GetShowDetails(show int) (*models.ShowDetails, error) {
	return internal.TryGetParsed(mappers.ParseShowDetails, endpoints.ShowDetailsUrl, src.Credentials, show)
}
