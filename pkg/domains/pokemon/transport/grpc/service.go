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
	converter := PokemonGrpcConverter{}

	out := converter.FromGrpcFilters(in)

	pokemonList, err := s.repository.FindByFilters(ctx, out)
	if err != nil {
		return nil, err
	}

	pokemons.Pokemons = converter.FromModels(pokemonList)

	return &pokemons, nil
}
