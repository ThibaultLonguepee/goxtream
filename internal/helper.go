package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thibaultlonguepee/goxtream/models"
)

func tryGetContent[Dto any](endpoint string, params ...any) (Dto, error) {
	var content Dto

	url := fmt.Sprintf(endpoint, params...)
	resp, err := http.Get(url)
	if err != nil {
		return content, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return content, err
	}

	err = json.Unmarshal(body, &content)
	if err != nil {
		return content, err
	}

	return content, nil
}

func TryGetParsed[Model, Dto any](Parser func(Dto) (Model, error), endpoint string, creds *models.Credentials, params ...any) (Model, error) {
	var model Model

	combined := append([]any{creds.Url, creds.Username, creds.Password}, params...)
	content, err := tryGetContent[Dto](endpoint, combined...)
	if err != nil {
		return model, err
	}

	model, err = Parser(content)
	if err != nil {
		return model, err
	}

	return model, nil
}
