package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas:")
	for _, areas := range resp.Results {
		fmt.Printf(" - %s\n", areas.Name)
	}
	cfg.nextLocationAddress = resp.Next
	cfg.previousLocationAddress = resp.Previous
	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.previousLocationAddress == nil {
		return errors.New("you are on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas:")
	for _, areas := range resp.Results {
		fmt.Printf(" - %s\n", areas.Name)
	}
	cfg.nextLocationAddress = resp.Next
	cfg.previousLocationAddress = resp.Previous
	return nil
}
