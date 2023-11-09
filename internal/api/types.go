package api

import "github.com/lookasc/pokedexcli/internal/cache"

type paginationLinks struct {
	next     string
	previous string
}

type locationAreaConfig struct {
	baseUrl    string
	pagination paginationLinks
}

type apiConfig struct {
	locationArea locationAreaConfig
}

type Api struct {
	config *apiConfig
	cache  *cache.Cache
}

// /location-area
// Paginated listing of location areas

type locationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type getLocationAreasResponseBody struct {
	Count    uint           `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

// /location-area/{id}
// Listing of pokemons found in specified location

type pokemon struct {
	Name string `json:"name"`
}

type pokemonEncounter struct {
	Pokemon pokemon `json:"pokemon"`
}

type getPokemonsEncountersInLocationAreaResponseBody struct {
	Encounters []pokemonEncounter `json:"pokemon_encounters"`
}

// /pokemon/{id}
// Get pokemon details

type GetPokemonDetailsResponseBody struct {
	ID             int                   `json:"id"`
	Name           string                `json:"name"`
	BaseExperience int                   `json:"base_experience"`
	Height         int                   `json:"height"`
	Weight         int                   `json:"weight"`
	Stats          []pokemonDetailsStats `json:"stats"`
	Types          []pokemonDetailsTypes `json:"types"`
}

type pokemonDetailsStat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type pokemonDetailsStats struct {
	BaseStat int                `json:"base_stat"`
	Effort   int                `json:"effort"`
	Stat     pokemonDetailsStat `json:"stat"`
}

type pokemonDetailsType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type pokemonDetailsTypes struct {
	Slot int                `json:"slot"`
	Type pokemonDetailsType `json:"type"`
}
