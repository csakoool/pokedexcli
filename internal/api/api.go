package api

import (
	"github.com/lookasc/pokedexcli/internal/cache"
)

const (
	apiBaseUrl      = "https://pokeapi.co/api/v2/"
	locationAreaUrl = apiBaseUrl + "location-area/"
	pokemonUrl      = apiBaseUrl + "pokemon/"
)

func NewApi(cache *cache.Cache) *Api {
	api := Api{
		config: &apiConfig{
			locationArea: locationAreaConfig{
				baseUrl: locationAreaUrl,
				pagination: paginationLinks{
					next: locationAreaUrl,
				},
			},
		},
		cache: cache,
	}
	return &api
}

func (api *Api) updateLocationAreaPaginationUrls(rb *getLocationAreasResponseBody) {
	if rb.Next == nil {
		api.config.locationArea.pagination.next = ""
	} else {
		api.config.locationArea.pagination.next = *rb.Next
	}

	if rb.Previous == nil {
		api.config.locationArea.pagination.previous = ""
	} else {
		api.config.locationArea.pagination.previous = *rb.Previous
	}
}

func (api *Api) getLocationAreas(url string) ([]locationArea, error) {
	parsedData, err := getData(api, url, getLocationAreasResponseBody{})
	if err != nil {
		return nil, err
	}
	api.updateLocationAreaPaginationUrls(&parsedData)
	return parsedData.Results, nil
}

func (api *Api) GetNextLocationAreas() ([]locationArea, error) {
	locationAreas, err := api.getLocationAreas(api.config.locationArea.pagination.next)
	return locationAreas, err
}

func (api *Api) GetPreviousLocationAreas() ([]locationArea, error) {
	locationAreas, err := api.getLocationAreas(api.config.locationArea.pagination.previous)
	return locationAreas, err
}

func (api *Api) GetPokemonsEncountersInLocationArea(locationArea string) ([]pokemonEncounter, error) {
	locationAreaUrl := api.config.locationArea.baseUrl + locationArea
	parsedData, err := getData(api, locationAreaUrl, getPokemonsEncountersInLocationAreaResponseBody{})
	if err != nil {
		return nil, err
	}
	return parsedData.Encounters, nil
}

func (api *Api) GetPokemonDetails(pokemonName string) (*GetPokemonDetailsResponseBody, error) {
	pokemonDetailsUrl := pokemonUrl + pokemonName
	parsedData, err := getData(api, pokemonDetailsUrl, GetPokemonDetailsResponseBody{})
	if err != nil {
		return nil, err
	}
	return &parsedData, nil
}
