package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)


func (c *Client) GetLocation(locationName string) (LocationAreaName, error){
	url := baseURL + "location-area/" + locationName 

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationAreaName{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationAreaName{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaName{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaName{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaName{}, err
	}

	locationResp := LocationAreaName{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationAreaName{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}


func (c *Client) GetPokemon(pokemonName string) (CatchPokeMon, error){
	url := baseURL + "pokemon/" + pokemonName 

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := CatchPokeMon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return CatchPokeMon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CatchPokeMon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CatchPokeMon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return CatchPokeMon{}, err
	}

	pokemonResp := CatchPokeMon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return CatchPokeMon{}, err
	}

	c.cache.Add(url, dat)

	return pokemonResp, nil
}