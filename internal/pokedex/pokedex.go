package pokedex

import (
	"fmt"

	"github.com/lookasc/pokedexcli/internal/api"
)

type Pokedex struct {
	pokemons map[string]api.GetPokemonDetailsResponseBody
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		pokemons: make(map[string]api.GetPokemonDetailsResponseBody),
	}
}

func (pokedex *Pokedex) Add(pokemonData *api.GetPokemonDetailsResponseBody) {
	if pokemonData == nil {
		return
	}
	pokedex.pokemons[pokemonData.Name] = *pokemonData
}

func (pokedex *Pokedex) Get(pokemonName string) (api.GetPokemonDetailsResponseBody, bool) {
	data, ok := pokedex.pokemons[pokemonName]
	return data, ok
}

func (pokedex *Pokedex) List() {
	if len(pokedex.pokemons) != 0 {
		fmt.Printf("Your Pokedex:\n")
		for pokemonName, _ := range pokedex.pokemons {
			fmt.Printf(" - %s\n", pokemonName)
		}
	} else {
		fmt.Printf("Your Pokedex is empty\n")
	}
}
