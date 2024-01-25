package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const POKEMONAPIURI = "https://pokeapi.co/api/v2/pokemon"

type PokemonItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonItemDetail struct {
	Name string `json:"name"`
	Weight int `json:"weight"`
	BaseExperience int `json:"base_experience"`
}

type PokemonResponse struct {
	Count   int           `json:"count"`
	Results []PokemonItem `json:"results"`
}

func fetchAndPrintPokemonData(pi *PokemonItem, resChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	innerClient := &http.Client{Timeout: 10 * time.Second}

	r, err := innerClient.Get(pi.URL)
	if err != nil {
		fmt.Println("Error fetching item data:", err)
		return
	}
	defer r.Body.Close()

	finalResult := new(PokemonItemDetail)
	err = json.NewDecoder(r.Body).Decode(finalResult)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	resChan <- finalResult.Name
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}
	listOfPokemons := new(PokemonResponse)
	sourceChan := make(chan string)
	var wg sync.WaitGroup

	r, err := client.Get(POKEMONAPIURI)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(listOfPokemons)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	for _, pi := range listOfPokemons.Results {
		wg.Add(1)
		localPi := pi
		go fetchAndPrintPokemonData(&localPi, sourceChan, &wg)
	}

	go func() {
		wg.Wait()
		close(sourceChan)
	}()

	for itemResult := range sourceChan {
		fmt.Println(itemResult)
	}
}
