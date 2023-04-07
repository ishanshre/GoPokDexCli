package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *config) error {
	fmt.Println("---------- Exiting ----------")
	os.Exit(0)
	return nil
}
