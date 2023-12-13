package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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

		err := command.run()
		if err != nil {
			fmt.Println(err)
		}
	}
}