package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
)

func (c *Client) ListLocations (pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	
	if pageURL != nil {
		url = *pageURL
	}

	//check if it exists in cache
	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationArea{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationArea{}, err
		}

		return locationsResp, nil
	}

	// if it doesn't exist in cache
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	// use httpClient to make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationsResp := LocationArea{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationArea{}, err
	}

	// add to cache
	c.cache.Add(url, dat)
	return locationsResp, nil
}