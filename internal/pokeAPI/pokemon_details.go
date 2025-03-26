package pokeAPI

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonDetails(idOrName string) (PokemonDetails, error) {
	url := baseURL + "/pokemon/" + idOrName + "/"

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := PokemonDetails{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return PokemonDetails{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetails{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetails{}, err
	}

	pokemonResp := PokemonDetails{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return PokemonDetails{}, err
	}

	return pokemonResp, nil
}