package grpc

import (
	"github.com/alexandredsa/poke-grpc/pkg/domains/pokemon/model"
	"github.com/alexandredsa/poke-grpc/pkg/domains/pokemon/transport/dto"
)

type PokemonGrpcConverter struct{}

func (converter PokemonGrpcConverter) FromModels(pokemonModels []model.Pokemon) []*Pokemon {
	pokemons := make([]*Pokemon, 0, len(pokemonModels))
	for _, pokemonModel := range pokemonModels {
		pokemons = append(pokemons, converter.FromModel(pokemonModel))
	}

	return pokemons
}

func (converter PokemonGrpcConverter) FromModel(pokemonModel model.Pokemon) *Pokemon {
	return &Pokemon{
		Id:             pokemonModel.PokemonID,
		Name:           pokemonModel.Name,
		Type1:          pokemonModel.Type1,
		Type2:          pokemonModel.Type2,
		Total:          pokemonModel.Total,
		Hp:             pokemonModel.Hp,
		Attack:         pokemonModel.Attack,
		Defense:        pokemonModel.Defense,
		SpecialAttack:  pokemonModel.SpecialAttack,
		SpecialDefense: pokemonModel.SpecialDefense,
		Speed:          pokemonModel.Speed,
		Generation:     pokemonModel.Generation,
		Legendary:      pokemonModel.Legendary,
	}
}

func (converter PokemonGrpcConverter) FromGrpcFilters(in *PokemonFilters) dto.Filters {
	out := dto.Filters{}

	if in == nil {
		return out
	}

	if in.Filters == nil {
		return out
	}

	if len(in.Filters) == 0 {
		return out
	}

	out.Filters = make([]dto.FilterRequest, len(in.Filters)-1)

	for _, inFilter := range in.Filters {
		out.Filters = append(out.Filters, dto.FilterRequest{
			Key:   inFilter.FilterKey,
			Type:  inFilter.FilterType,
			Value: inFilter.FilterValue,
		})
	}

	return out
}
