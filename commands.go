package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/lookasc/pokedexcli/internal/api"
	"github.com/lookasc/pokedexcli/internal/pokedex"
)

type cliCommand struct {
	name        string
	description string
	callback    func(api *api.Api, pokedex *pokedex.Pokedex, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of all the Pokémon in a given area: eg. 'explore lake-of-rage-area'",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch the Pokémon with a given name: eg. 'catch pikachu'",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Shows Pokémon's details if already caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints caught Pokémons",
			callback:    commandPokedex,
		},
	}
}

func executeCommand(userInput string, api *api.Api, pokedex *pokedex.Pokedex) error {
	parsedCommand := strings.Split(userInput, " ")
	commandName := parsedCommand[0]

	command, hasCommand := getCommands()[commandName]
	if !hasCommand {
		fmt.Println("Unknown command:", userInput)
		return nil
	}

	return command.callback(api, pokedex, parsedCommand[1:]...)
}

func commandHelp(api *api.Api, pokedex *pokedex.Pokedex, args ...string) error {
	commands := getCommands()

	fmt.Print("\nWelcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Print("\n")
	return nil
}

func commandExit(api *api.Api, pokedex *pokedex.Pokedex, args ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(api *api.Api, pokedex *pokedex.Pokedex, args ...string) error {
	locationAreas, err := api.GetNextLocationAreas()
	if err != nil {
		if err.Error() == "Url is missing" {
			return errors.New("No more items to map")
		}
		return err
	}

	fmt.Print("\n")
	for _, location := range locationAreas {
		fmt.Printf("%s\n", location.Name)
	}
	fmt.Print("\n")

	return nil
}

func commandMapb(api *api.Api, pokedex *pokedex.Pokedex, args ...string) error {
	locationAreas, err := api.GetPreviousLocationAreas()
	if err != nil {
		if err.Error() == "Url is missing" {
			return errors.New("No more items to map")
		}
		return err
	}

	fmt.Print("\n")
	for _, location := range locationAreas {
		fmt.Printf("%s\n", location.Name)
	}
	fmt.Print("\n")

	return nil
}

func commandExplore(api *api.Api, pokedex *pokedex.Pokedex, args ...string) error {
	if len(args) == 0 {
		return errors.New(
			"Missing location name in 'explore' command. " +
				"Try 'explore canalave-city-area' or use 'map' command first.",
		)
	}
	locationArea := args[0]
	fmt.Printf("\nExploring %s...\n\n", locationArea)

	encounters, err := api.GetPokemonsEncountersInLocationArea(locationArea)
	if err != nil {
		return err
	}

	fmt.Print("Found Pokemons: \n")
	for _, encounter := range encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	fmt.Print("\n")

	return nil
}

func commandCatch(api *api.Api, pokedex *pokedex.Pokedex, args ...string) error {
	if len(args) == 0 {
		return errors.New("Missing Pokémon's name in 'catch' command.")
	}
	pokemonName := args[0]
	fmt.Printf("\nThrowing a Pokeball at %s...\n", pokemonName)

	details, err := api.GetPokemonDetails(pokemonName)
	if err != nil {
		return err
	}

	chance := rand.Intn(255) - details.BaseExperience
	if chance > 0 {
		fmt.Printf("%s was caught!\n", details.Name)
		pokedex.Add(details)
	} else {
		fmt.Printf("%s escaped!\n", details.Name)
	}

	return nil
}

func commandInspect(api *api.Api, pokedex *pokedex.Pokedex, args ...string) error {
	if len(args) == 0 {
		return errors.New("Missing Pokémon's name in 'inspect' command.")
	}
	pokemonName := args[0]
	details, hasPokemon := pokedex.Get(pokemonName)

	if hasPokemon {
		fmt.Printf("\nName: %s\n", details.Name)
		fmt.Printf("Height: %v\n", details.Height)
		fmt.Printf("Weight: %v\n", details.Weight)
		if len(details.Stats) != 0 {
			fmt.Printf("Stats:\n")
			for _, stat := range details.Stats {
				fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
			}
		}
		if len(details.Types) != 0 {
			fmt.Printf("Types:\n")
			for _, t := range details.Types {
				fmt.Printf("  -%s\n", t.Type.Name)
			}
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}

func commandPokedex(api *api.Api, pokedex *pokedex.Pokedex, args ...string) error {
	pokedex.List()
	return nil
}
