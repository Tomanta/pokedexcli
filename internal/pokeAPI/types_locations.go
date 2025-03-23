package pokeAPI

type RespShallowLocations struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	LocationAreas  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
}