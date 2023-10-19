package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetMap(URL *string) (LocationResp, error) {
	endpoint := "/location-area"
	url := baseURL + endpoint
	if URL != nil {
		url = *URL
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

	return locationResp, nil
}