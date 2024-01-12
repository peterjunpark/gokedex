package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2"

func (c *Client) ListLocations(pageUrl *string) (LocationsRes, error) {
	url := baseUrl + "/location-area?offset=0&limit=20"

	if pageUrl != nil {
		url = *pageUrl
	}

	// check cache
	data, ok := c.cache.Get(url)
	if ok {
		// cache hit
		locations := LocationsRes{}
		err := json.Unmarshal(data, &locations)
		if err != nil {
			return LocationsRes{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsRes{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsRes{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationsRes{}, fmt.Errorf("something went wrong %d", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationsRes{}, err
	}

	locations := LocationsRes{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationsRes{}, err
	}

	c.cache.Add(url, data)

	return locations, nil
}

func (c *Client) GetLocation(locationName string) (LocationRes, error) {
	url := baseUrl + "/location-area/" + locationName

	// check cache
	data, ok := c.cache.Get(url)
	if ok {
		// cache hit
		location := LocationRes{}
		err := json.Unmarshal(data, &location)
		if err != nil {
			return LocationRes{}, err
		}

		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationRes{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationRes{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationRes{}, fmt.Errorf("something went wrong %d", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationRes{}, err
	}

	location := LocationRes{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return LocationRes{}, err
	}

	c.cache.Add(url, data)

	return location, nil
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + pokemonName

	// check cache
	data, ok := c.cache.Get(url)
	if ok {
		// cache hit
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("something went wrong %d", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return pokemon, nil
}
