package goldapi

import (
	"io/ioutil"
	"net/http"
)

type API struct {
	Key string
	Host string
	Client *http.Client
}

func NewAPIClient(apiKey string) *API {
	return &API{
		Key:    apiKey,
		Host:   "https://www.goldapi.io/api/",
		Client: &http.Client{},
	}
}

func (a *API) get(path string) (int, []byte, error) {
	req, err := http.NewRequest("GET", a.Host+path, nil)
	if err != nil {
		return -1, nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-access-token", a.Key)

	resp, err := a.Client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, byt, err
}
