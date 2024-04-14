package pokeAPI

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	if val, ok := c.cache.Get(url); ok {
		RespLoc := RespLocations{}
		err := json.Unmarshal(val, &RespLoc)
		if err != nil {
			return RespLocations{}, err
		}
		return RespLoc, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocations{}, err
	}
	RespLoc := RespLocations{}
	err = json.Unmarshal(body, &RespLoc)
	if err != nil {
		return RespLocations{}, err
	}
	c.cache.Add(url, body)
	return RespLoc, nil
}
