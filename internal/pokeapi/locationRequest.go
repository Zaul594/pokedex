package pokeapi

//pokeapi is used to set up the api for the pokedex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

// GetMap gets the location from https://pokeapi.co/api/v2
func (c *Client) GetMap(URL *string) (LocationResp, error) {
	endpoint := "/location-area"
	url := baseURL + endpoint
	if URL != nil {
		url = *URL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationResp{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationResp{}, err
		}

		return locationResp, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResp{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationResp{}, err
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return LocationResp{}, fmt.Errorf("bad status code %v", response.Status)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationResp{}, err
	}

	locationResp := LocationResp{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationResp{}, err
	}

	c.cache.Add(url, data)
	return locationResp, nil
}
