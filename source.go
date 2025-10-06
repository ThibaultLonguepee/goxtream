package goxtream

import "github.com/thibaultlonguepee/goxtream/internal"

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

func (src *Source) Authenticate() (*Authentication, error) {
	return internal.TryGetParsed(parseAuthentication, AuthenticationUrl, src.url, src.username, src.password)
}

func (src *Source) GetLiveCategories() ([]*Category, error) {
	return internal.TryGetParsed(parseCategories, LiveCategoriesUrl, src.url, src.username, src.password)
}

func (src *Source) GetVodCategories() ([]*Category, error) {
	return internal.TryGetParsed(parseCategories, VodCategoriesUrl, src.url, src.username, src.password)
}

func (src *Source) GetSeriesCategories() ([]*Category, error) {
	return internal.TryGetParsed(parseCategories, SeriesCategoriesUrl, src.url, src.username, src.password)
}
