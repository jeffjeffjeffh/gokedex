package main

import (
	"fmt"
	"os"
)

type command struct {
	description string
	run         func(*Config) error
}

func getCommands() map[string]command {
	cmds := map[string]command{
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
	}

	return cmds
}

func help(cfg *Config) error {
	for name, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}

	return nil
}

func exit(cfg *Config) error {
	fmt.Println("Bye!")
	os.Exit(0)
	return nil
}

func getMap(cfg *Config) error {
	resp, err := cfg.pokeclient.ListLocations(cfg.nextUrl)
	if err != nil {
		return err
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}