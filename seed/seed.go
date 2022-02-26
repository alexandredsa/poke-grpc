package main

import (
	"fmt"
	"log"

	"github.com/alexandredsa/poke-grpc/seed/reader"
)

func main() {
	csvReader := reader.NewPokemonCSVReader()
	pokemonList, err := csvReader.Load("./pokemon.csv")
	if err != nil {
		log.Fatal(err)
	}

	for _, pokemon := range pokemonList {
		fmt.Println(pokemon.Name)
	}
}
