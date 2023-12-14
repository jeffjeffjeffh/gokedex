package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"internal/pokeapi"
)


type Config struct{
	pokeClient pokeapi.Client
	next *string
	prev *string
}

func main() {
	var BASE_URL string = "https://pokeapi.co/api/v2/location-area"
	
	cfg := Config{
		pokeClient: pokeapi.NewClient(time.Minute, time.Second, time.Second / 2),
		next: &BASE_URL,
		prev: nil,
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