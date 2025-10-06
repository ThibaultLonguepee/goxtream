package goxtream

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
	url := fmt.Sprintf(AuthenticationFormat, src.url, src.username, src.password)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var auth dtos.AuthResult
	err = json.Unmarshal(body, &auth)
	if err != nil {
		return nil, err
	}

	connection, err := parseAuthentication(auth)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
