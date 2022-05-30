package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func fetchAndUnmarshal(url string, v interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}
