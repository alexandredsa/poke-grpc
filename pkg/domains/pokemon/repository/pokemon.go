package repository

import (
	"context"

	"github.com/sirupsen/logrus"

	mongoComponent "github.com/alexandredsa/poke-grpc/internal/components/mongo"
	"github.com/alexandredsa/poke-grpc/pkg/domains/pokemon/model"
	"github.com/alexandredsa/poke-grpc/pkg/domains/pokemon/transport/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

type PokemonRepository interface {
	InsertAll(context.Context, []*model.Pokemon) error
	FindByFilters(context.Context, dto.Filters) ([]model.Pokemon, error)
}

type PokemonMongoRepository struct {
	collection mongo.Collection
}

func NewPokemonMongoRepository(client mongoComponent.ClientI) PokemonMongoRepository {
	pokemonCol := client.GetCollection("pokemons")

	return PokemonMongoRepository{
		collection: pokemonCol,
	}
}

func (r PokemonMongoRepository) InsertAll(ctx context.Context, pokemons []*model.Pokemon) error {
	docs := make([]interface{}, 0, len(pokemons))

	logrus.
		WithFields(logrus.Fields{
			"slice_size": len(pokemons),
		}).
		Debug("InsertAll")

	for _, pokemon := range pokemons {
		docs = append(docs, pokemon)
	}

	_, err := r.collection.InsertMany(ctx, docs)

	return err
}

func (r PokemonMongoRepository) FindByFilters(ctx context.Context, filters dto.Filters) ([]model.Pokemon, error) {
	return nil, nil
}
