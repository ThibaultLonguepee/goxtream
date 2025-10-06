package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func TryGetContent[Dto any](endpoint string, params ...any) (Dto, error) {
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

func TryGetParsed[Model, Dto any](parser func(Dto) (Model, error), endpoint string, params ...any) (Model, error) {
	var model Model

	content, err := TryGetContent[Dto](endpoint, params...)
	if err != nil {
		return model, err
	}

	model, err = parser(content)
	if err != nil {
		return model, err
	}

	return model, nil
}
