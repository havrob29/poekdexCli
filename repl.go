package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/havrob29/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocation     *string
	previousLocation *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Welcome to your pokedex, Ash!\n")
	for {
		fmt.Print("Pokedex >")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
			}
		}
		input := scanner.Text()
		if cmd, ok := getCmds()[input]; ok {
			if err := cmd.callback(cfg); err != nil {
				fmt.Println("Error executing command:", err)
			}
			continue
		} else {
			fmt.Println("Command not found. Use 'help'")
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCmds() map[string]cliCommand {
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
			description: "display 20 location areas",
			callback:    cmdMap,
		},
		"mapb": {
			name:        "mapb",
			description: "display previous 20 locations",
			callback:    cmdMapb,
		},
	}
}
