package dota

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DotaMatchClient struct {
	DotaClient
}

// NewMatchClient returns a new intitialised Dota Match Client using Dota Client.
// The Steam credentials are passed onto the Dota Match Client.
//
//	Example:
//		matches := dota.NewMatchClient(client)
func NewMatchClient(client *DotaClient) *DotaMatchClient {
	return &DotaMatchClient{
		DotaClient: *client,
	}
}

// GetLeaugeListing Returns list of all leagues supported in-game via DotaTV.
func (m *DotaMatchClient) GetLeagueListing() ([]League, error) {
	res, err := http.Get(fmt.Sprintf("%s/%s/GetLeagueListing/v1/?key=%s", baseURL, match, m.key))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response struct {
		Result struct {
			Leagues []League `json:"leagues"`
		} `json:"result"`
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Result.Leagues, nil
}
