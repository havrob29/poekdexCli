package main

import (
	"fmt"
	"os"
)

func commandExit(*config) error {
	fmt.Println("Exiting the Pokedex.")
	os.Exit(0)
	return nil
}
