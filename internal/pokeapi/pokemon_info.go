package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name

	val, ok := c.cache.Get(url)
	if ok {
		pokemonInfo := RespPokemon{}
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			return RespPokemon{}, err
		}
		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	pokemonInfo := RespPokemon{}
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemonInfo, nil
}
