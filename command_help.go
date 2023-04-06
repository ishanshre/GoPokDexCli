package main

import "fmt"

func callbackHelp() error {
	fmt.Println("++++++++++++++++++ Welcome to PokDex ++++++++++++++++++")
	fmt.Println("++++++++++++++++++ Avaliable Commands ++++++++++++++++++")
	fmt.Println("----------------------------------------------------------")
	avaliableCommands := getCommands()
	for _, cmd := range avaliableCommands {
		fmt.Printf("+ %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
