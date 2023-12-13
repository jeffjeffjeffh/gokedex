package main

import (
	"fmt"
	"os"

	"github.com/jeffjeffjeff/pokedex2/internal/pokeapi"
)

type command struct {
	description string
	run         func() error
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
			run: pokeapi.Map,
		},
	}

	return cmds
}

func help() error {
	for name, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}

	return nil
}

func exit() error {
	os.Exit(0)
	return nil
}