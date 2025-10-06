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
	return internal.TryGetParsed(mappers.ParseAuthentication, AuthenticationUrl, src.url, src.username, src.password)
}

func (src *Source) GetLiveCategories() ([]*models.Category, error) {
	return internal.TryGetParsed(mappers.ParseCategories, LiveCategoriesUrl, src.url, src.username, src.password)
}

func (src *Source) GetVodCategories() ([]*models.Category, error) {
	return internal.TryGetParsed(mappers.ParseCategories, VodCategoriesUrl, src.url, src.username, src.password)
}

func (src *Source) GetSeriesCategories() ([]*models.Category, error) {
	return internal.TryGetParsed(mappers.ParseCategories, SeriesCategoriesUrl, src.url, src.username, src.password)
}

func (src *Source) GetLiveStreams() ([]*models.LiveStream, error) {
	return internal.TryGetParsed(mappers.ParseLiveStreams, LiveStreamsUrl, src.url, src.username, src.password)
}

func (src *Source) GetCategoryLiveStreams(category int) ([]*models.LiveStream, error) {
	return internal.TryGetParsed(mappers.ParseLiveStreams, CategoryLiveStreamsUrl, src.url, src.username, src.password, category)
}
