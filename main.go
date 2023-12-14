package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		pokeClient: pokeapi.NewClient(time.Minute, time.Minute * 5, time.Minute * 5),
		next: &BASE_URL,
		prev: nil,
	}

	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	var userCmd string
	
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := strings.Fields(scanner.Text())

		command, ok := commands[input[0]]
		if !ok {
			fmt.Printf("no such command %s\n", userCmd)
			continue
		}

		if len(input) == 1 {
			err := command.run(&cfg, nil)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := command.run(&cfg, &input[1])
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
