package main

import (
	"fmt"
)

func commandHelp(*config) error {
	fmt.Println("Here are your poke-commands: \n")
	for _, cmd := range getCmds() {
		fmt.Printf("%v: %v\n\n", cmd.name, cmd.description)
	}
	return nil
}
