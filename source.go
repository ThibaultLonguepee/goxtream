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
	response, err := tryGetParsed[dtos.AuthResult](AuthenticationFormat, src.url, src.username, src.password)
	if err != nil {
		return nil, err
	}

	auth, err := parseAuthentication(response)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
