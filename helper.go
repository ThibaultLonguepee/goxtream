package goxtream

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func tryGetParsed[T any](endpoint string, params ...any) (*T, error) {
	url := fmt.Sprintf(endpoint, params...)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var content T
	err = json.Unmarshal(body, &content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}
