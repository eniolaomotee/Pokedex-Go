package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


type LocationArea struct{
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct{
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}


func commandMap(cfg *config) error{


	url := "https://pokeapi.co/api/v2/location-area"

	// Use if there is a next page(cfg.Next) or previous page(cfg.Previous)
	if cfg.Next != nil{
		url = *cfg.Next
	}

	cached, ok := cfg.cache.Get(url)
	if ok{
		var data LocationArea
		err := json.Unmarshal(cached, &data)
		if err != nil{
			log.Fatalf("Error unmarshaling cached JSON : %v", err)

		}

		cfg.Next = data.Next
		cfg.Previous = data.Previous

		for _, location := range data.Results{
			fmt.Println(location.Name)
		}
	}
	

	

	// Make the HTTP GET request
	res, err := http.Get(url)
	if err != nil{
		log.Fatalf("Error fetching locations: %v", err)
	}

	body,err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299{
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)

	}
	if err != nil{
		log.Fatal(err)
	}

	// Unmarshal the JSON response into a struct
	var data LocationArea
	err = json.Unmarshal(body, &data)
	if err != nil{
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// cache the response data
	cfg.cache.Add(url, body)

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, location := range data.Results{
		fmt.Println(location.Name)
	}
	return nil
}


func commandMapB(cfg *config) error{

	if cached, ok := cfg.cache.Get(*cfg.Previous); ok{
		var data LocationArea
		err := json.Unmarshal(cached, &data)
		if err != nil{
			log.Fatalf("Error unmarshaling cached JSON : %v", err)

		}

		cfg.Next = data.Next
		cfg.Previous = data.Previous

		for _, location := range data.Results{
			fmt.Println(location.Name)
		}
		return nil
	}


	if cfg.Previous == nil{
		fmt.Println("No previous page available.")
		return  nil
	}

	url := *cfg.Previous

	// Make the HTTP GET request
	res, err := http.Get(url)
	if err != nil{
		log.Fatalf("error fectching locations : %v", err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299{
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil{
		log.Fatalf("Error reading response body: %v", err)
	}

	cfg.cache.Add(url, body)

	var data LocationArea
	err = json.Unmarshal(body, &data)
	if err != nil{
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, location := range data.Results{
		fmt.Println(location.Name)
	}

	return  nil
	
}