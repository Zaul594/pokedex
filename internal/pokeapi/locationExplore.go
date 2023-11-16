package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationExplore(location string) (Location, error) {

	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}

		return locationResp, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return Location{}, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)

	return locationResp, nil
}
