package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AbdKaan/pokedexcli/internal/pokecache"
)

func (c *Client) DetailLocation(pokeCache *pokecache.Cache, location string) (Area, error) {
	var details Area
	url := baseURL + "/location-area/" + location

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return details, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return details, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return details, err
	}

	// Add data to pokecache
	pokeCache.Add(location, data)

	if err := json.Unmarshal(data, &details); err != nil {
		return details, err
	}

	return details, nil
}
