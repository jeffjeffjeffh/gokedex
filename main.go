package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"internal/pokeapi"
)

type Config struct{
	pokeclient pokeapi.Client
	nextUrl *string
	prevUrl *string
}

func main() {
	cfg := Config{
		pokeclient: pokeapi.NewClient(time.Minute),
		nextUrl: nil,
		prevUrl: nil,
	}

	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	var userCmd string
	
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		userCmd = scanner.Text()

		command, ok := commands[userCmd]
		if !ok {
			fmt.Printf("no such command %s\n", userCmd)
			continue
		}

		err := command.run(&cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}