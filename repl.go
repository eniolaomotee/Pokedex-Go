package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"github.com/eniolaomotee/Pokedex-Go/internal/pokeapi"
)



type config struct {
	Next *string
	Previous *string
	pokeapiClient *pokeapi.Client
	cmd string
	args []string
}



type cliCommand struct{
	name string
	description string
	callback func(*config, []string) error
}



func StartRepl(cfg *config){
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanned := scanner.Scan()
		if !scanned{
			return 
		}
		word := scanner.Text()
		text := CleanInput(word)
		if len(text) == 0 {
			continue
		}

		cmd := text[0]
		cfg.cmd = cmd
		if len(text) > 1{
			cfg.args = text[1:]
		}else{
			cfg.args = []string{}
		}
		

		if command, ok := getCommands()[cmd]; ok{
			if err := command.callback(cfg, cfg.args); err != nil{
				fmt.Println("error", err)
			}
			continue
		}

		fmt.Println("Unknown command")
		continue
		
	}

}



func CleanInput(text string) []string{
	cleanText := strings.ToLower(text)
	words := strings.Fields(cleanText)
	return words
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
		"explore":{
			name: "explore",
			description: "explore the location area of a Pokemon",
			callback: commandExplore,
		},
		"catch":{
			name :"catch",
			description: "catch a Pokemon in a location area",
			callback: commandCatch,
		},
	}
}
