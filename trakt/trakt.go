package trakt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Trakt struct {
	Url         string
	ApiKey      string
	AccessToken string
}

func (t Trakt) Get(path string) (resp *http.Response, err error) {
	url := fmt.Sprintf("%s%s", t.Url, path)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("trakt-api-version", "2")
	req.Header.Add("trakt-api-key", t.ApiKey)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.AccessToken))
	return client.Do(req)
}
