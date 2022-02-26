package grpc

import (
	"context"
	"log"

	"github.com/alexandredsa/poke-grpc/pkg/domains/pokemon/repository"
	grpc "google.golang.org/grpc"
)

type PokemonService struct {
	UnimplementedPokemonServiceServer

	repository repository.PokemonRepository
}

func NewPokemonService(repository repository.PokemonRepository) PokemonService {
	return PokemonService{repository: repository}
}

func (service PokemonService) Register(grpcServer *grpc.Server) {
	RegisterPokemonServiceServer(grpcServer, &service)
}

func (s PokemonService) List(ctx context.Context, in *PokemonFilters) (*Pokemons, error) {
	log.Printf("Received: %v", in)

	pokemons := Pokemons{}
	pokemons.Pokemons = make([]*Pokemon, 0)
	pokemons.Pokemons = append(pokemons.Pokemons, &Pokemon{
		Name: "Bulbassoro",
		Id:   "92c5cf40-ac46-4782-b6ba-00c40e402602",
	})

	return &pokemons, nil
}
