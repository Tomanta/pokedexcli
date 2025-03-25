package pokeAPI

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationDetails(idOrName string) (RespDetailLocations, error) {
	url := baseURL + "/location-area/" + idOrName + "/"

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespDetailLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespDetailLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDetailLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailLocations{}, err
	}

	locationsResp := RespDetailLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespDetailLocations{}, err
	}

	return locationsResp, nil
}