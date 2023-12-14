package main

import (
	"errors"
	"fmt"
	"os"
)

type command struct {
	description string
	run         func(*Config, *string) error
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			description: "get info on possible commands",
			run:         help,
		},
		"exit": {
			description: "exit the program",
			run: exit,
		},
		"map": {
			description: "get the next 20 locations",
			run: getMap,
		},
		"bmap": {
			description: "get the previous 20 locations",
			run: getBmap,
		},
		"explore": {
			description: "enter the name of a location to find all the pokemon there",
			run: explore,
		},
	}
}

func help(cfg *Config, s *string) error {
	for name, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}

	return nil
}

func exit(cfg *Config, s *string) error {
	fmt.Println("Bye!")
	os.Exit(0)
	return nil
}

func getMap(cfg *Config, s *string) error {
	if cfg.next == nil {
		return errors.New("already at the end of location areas")
	}

	res, err := cfg.pokeClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}

	cfg.next = res.Next
	cfg.prev = res.Previous

	return nil
}

func getBmap(cfg *Config, s *string) error {
	if cfg.prev == nil {
		return errors.New("already at the beginning of location areas")
	}

	res, err := cfg.pokeClient.ListLocations(cfg.prev)
	if err != nil {
		return err
	}

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}

	cfg.next = res.Next
	cfg.prev = res.Previous

	return nil
}

func explore(cfg *Config, area *string) error {
	if area == nil {
		return errors.New("please include an area to explore")
	}
	
	pokemonList, err := cfg.pokeClient.ListPokemon(*area)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemonList {
		fmt.Println(pokemon)
	}

	return nil
}