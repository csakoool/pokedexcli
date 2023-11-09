package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/lookasc/pokedexcli/internal/api"
	"github.com/lookasc/pokedexcli/internal/cache"
	"github.com/lookasc/pokedexcli/internal/pokedex"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cache := cache.NewCache(5 * time.Second)
	api := api.NewApi(cache)
	pokedex := pokedex.NewPokedex()

	fmt.Print("pokedex > ")

	for scanner.Scan() {
		userInput := scanner.Text()
		err := executeCommand(userInput, api, pokedex)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("pokedex > ")
	}

	scannerError := scanner.Err()
	if scannerError != nil {
		fmt.Println(scanner)
	} else {
		fmt.Println("Ending without error")
	}
}
