package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Credentials interface {
	Authenticated() bool
	Url() string
	Username() string
	Password() string
}

func Get[Model, Dto any](Parser func(Dto) (Model, error), endpoint string, params ...any) (Model, error) {
	var model Model
	var content Dto

	url := fmt.Sprintf(endpoint, params...)
	resp, err := http.Get(url)
	if err != nil {
		return model, fmt.Errorf("could not get data at endpoint '%v': %v", url, err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model, err
	}

	err = json.Unmarshal(body, &content)
	if err != nil {
		return model, err
	}

	model, err = Parser(content)
	if err != nil {
		return model, err
	}

	return model, nil
}

func GetWithCredentials[Model, Dto any](Parser func(Dto) (Model, error), endpoint string, creds Credentials, params ...any) (Model, error) {
	if !creds.Authenticated() {
		var empty Model
		return empty, fmt.Errorf("could not perform request: source is unauthentified")
	}
	combined := append([]any{creds.Url(), creds.Username(), creds.Password()}, params...)
	return Get(Parser, endpoint, combined...)
}
