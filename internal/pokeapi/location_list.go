package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AbdKaan/pokedexcli/internal/pokecache"
)

func (c *Client) ListLocations(pokeCache *pokecache.Cache, pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	// Add data to pokecache
	pokeCache.Add(url, data)

	locationsResp := LocationArea{}
	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return LocationArea{}, err
	}

	return locationsResp, nil
}
