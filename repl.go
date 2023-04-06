package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		textInp := scanner.Text()
		cleaned := cleanInput(textInp)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		avaliableCommands := getCommands()
		command, ok := avaliableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		command.callback()
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommands struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Displays the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokdex cli",
			callback:    callbackExit,
		},
	}
}
