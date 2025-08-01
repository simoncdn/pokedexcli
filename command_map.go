package main

import "fmt"

func commandMapF(cfg *config, args ...string) error {
	listLocationRes, err := cfg.pokeapiClient.ListLocation(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = listLocationRes.Next
	cfg.prevLocationsURL = listLocationRes.Previous

	for _, location := range listLocationRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(cfg *config, args ...string) error { 
	url := cfg.prevLocationsURL
	listLocationRes, err := cfg.pokeapiClient.ListLocation(url)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = listLocationRes.Next
	cfg.prevLocationsURL = listLocationRes.Previous

	for _, location := range listLocationRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}
