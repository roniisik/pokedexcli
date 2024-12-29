package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(areaName string) (RespAreaInfo, error) {
	url := baseURL + "/location-area/" + areaName

	val, ok := c.cache.Get(url)
	if ok {
		areaInfo := RespAreaInfo{}
		if err := json.Unmarshal(val, &areaInfo); err != nil {
			return RespAreaInfo{}, err
		}
		return areaInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespAreaInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespAreaInfo{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespAreaInfo{}, err
	}

	areaInfo := RespAreaInfo{}
	if err := json.Unmarshal(data, &areaInfo); err != nil {
		return RespAreaInfo{}, err
	}

	c.cache.Add(url, data)
	return areaInfo, nil

}
