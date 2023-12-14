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
		"catch": {
			description: "enter the name of a pokemon to try and catch",
			run: catch,
		},
		"inspect": {
			description: "enter the name of a pokemon in your pokedex to see its info",
			run: inspect,
		},
		"pokedex": {
			description: "list all the caught pokemon in your pokedex",
			run: pokedex,
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
		return errors.New("please specify an area to explore")
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

func catch(cfg *Config, name *string) error {
	if name == nil {
		return errors.New("please specify a pokemon to try to catch")
	}

	if _, ok := cfg.pokedex.get(*name); ok {
		return errors.New("you already have this pokemon")
	}
	
	res, err := cfg.pokeClient.Catch(*name)
	if err != nil {
		return err
	}

	fmt.Printf("you rudely threw a pokeball at a %s...\n", *name)

	catchSucceeded := rollToCatch(res.BaseExperience)
	if !catchSucceeded {
		fmt.Printf("aw, crap! the %s got away!\n", *name)
		return nil
	}

	fmt.Printf("hey, man, sick! you caught that %s!\n", *name)

	cfg.pokedex.add(*name, newPokemon(res))

	return nil
}

func inspect(cfg *Config, name *string) error {
	if name == nil {
		return errors.New("please enter the name of a pokemon you wish to inspect")
	}

	pokemon, ok := cfg.pokedex.get(*name)
	if !ok {
		return fmt.Errorf("you have not caught a %s yet", *name)
	}

	fmt.Printf("Name: %s\n", pokemon.name)
	fmt.Printf("Height: %d\n", pokemon.height)
	fmt.Printf("Weight: %d\n", pokemon.weight)

	fmt.Printf("Stats:\n")
	for name, val := range pokemon.stats {
		fmt.Printf("\t%s: %d\n", name, val)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.types {
		fmt.Printf("\t%s\n", t)
	}

	return nil
}

func pokedex(cfg *Config, s *string) error {
	names, err := cfg.pokedex.list()
	if err != nil {
		return err
	}

	fmt.Println("your caught pokemon:")

	for _, name := range names {
		fmt.Printf("- %s\n", name)
	}

	return nil
}