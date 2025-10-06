package goxtream

import (
	"github.com/thibaultlonguepee/go-xtream/internal/dtos"
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

func (src *Source) Authenticate() (*Authentication, error) {
	response, err := tryGetParsed[dtos.AuthResult](AuthenticationUrl, src.url, src.username, src.password)
	if err != nil {
		return nil, err
	}

	auth, err := parseAuthentication(response)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (src *Source) getCategories(url string) ([]*Category, error) {
	response, err := tryGetParsed[[]dtos.CategoryResult](url, src.url, src.username, src.password)
	if err != nil {
		return nil, err
	}

	auth, err := parseCategories(*response)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (src *Source) GetLiveCategories() ([]*Category, error) {
	return src.getCategories(LiveCategoriesUrl)
}

func (src *Source) GetVodCategories() ([]*Category, error) {
	return src.getCategories(VodCategoriesUrl)
}

func (src *Source) GetSeriesCategories() ([]*Category, error) {
	return src.getCategories(SeriesCategoriesUrl)
}
