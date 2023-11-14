package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationExplore(location string) (Pokemon, error) {
	var id int
	resp, err := c.GetMap(nil)
	if err != nil {
		return Pokemon{}, err
	}
	for _, area := range resp.Results {
		if location == area.Name {
			id = area.ID
		}
	}

	url := baseURL + "/location-area" + fmt.Sprint(id)

	if val, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return Pokemon{}, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return Pokemon{}, err
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code %v", response.Status)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}

	locationResp := LocationResp{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return Pokemon{}, nil
}
