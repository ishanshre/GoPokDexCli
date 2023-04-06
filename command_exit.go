package main

import (
	"fmt"
	"os"
)

func callbackExit() error {
	fmt.Println("---------- Exiting ----------")
	os.Exit(0)
	return nil
}
