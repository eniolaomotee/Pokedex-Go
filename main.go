package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	internal "github.com/eniolaomotee/Pokedex-Go/internal"
	"github.com/eniolaomotee/Pokedex-Go/internal/pokecache"
)

type config struct {
	Next *string
	Previous *string
	cache *pokecache.Cache

}

type cliCommand struct{
	name string
	description string
	callback func(*config) error
}



func main(){

	cfg := &config{}

	cache := pokecache.NewCache(5 * time.Second)
	cfg.cache = cache

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanned := scanner.Scan()
		if !scanned{
			return 
		}
		word := scanner.Text()
		text := internal.CleanInput(word)
		if len(text) == 0 {
			continue
		}

		cmd := text[0]

		if command, ok := getCommands()[cmd]; ok{
			if err := command.callback(cfg); err != nil{
				fmt.Println("error", err)
			}
			continue
		}

		fmt.Println("Unknown command")
		continue
		
	}

}


func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit":{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map":{
			name: "map",
			description: "Location areas of Pokemon map",
			callback: commandMap,
		},
		"mapb":{
			name :"mapb",
			description: "Previous page of location areas of Pokemon map",
			callback: commandMapB,
		},
	}
}

