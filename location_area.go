package main

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Initial endpoint
var locationArea LocationArea = LocationArea{
	Next: "https://pokeapi.co/api/v2/location-area",
}
