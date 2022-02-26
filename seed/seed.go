package main

import (
	"context"
	"log"

	"github.com/alexandredsa/poke-grpc/internal/container"
	"github.com/alexandredsa/poke-grpc/seed/reader"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	csvReader := reader.NewPokemonCSVReader()
	pokemonList, err := csvReader.Load("./seed/reader/pokemon.csv")
	if err != nil {
		log.Fatal(err)
	}

	logrus.WithField("total", len(pokemonList)).Info("csv loaded")

	dep := container.Dependencies{}
	if err := dep.Setup(ctx); err != nil {
		log.Fatal(err)
	}

	pokemonRepository := dep.Providers.PokemonRepository
	if err := pokemonRepository.InsertAll(ctx, pokemonList); err != nil {
		log.Fatal(err)
	}
}
