package main

import (
	"errors"
	"fmt"
)

func cmdMap(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = locationsResp.Next
	cfg.previousLocation = locationsResp.Next

	for _, result := range locationsResp.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func cmdMapb(cfg *config) error {
	if cfg.previousLocation == nil {
		return errors.New("cant go back")
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = locationsResp.Next
	cfg.previousLocation = locationsResp.Next

	for _, result := range locationsResp.Results {
		fmt.Println(result.Name)
	}
	return nil
}
