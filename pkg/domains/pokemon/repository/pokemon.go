package repository

import (
	"context"

	"github.com/sirupsen/logrus"

	mongoComponent "github.com/alexandredsa/poke-grpc/internal/components/mongo"
	"github.com/alexandredsa/poke-grpc/pkg/domains/pokemon/model"
	"github.com/alexandredsa/poke-grpc/pkg/domains/pokemon/transport/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	listDefaultLimit int64 = 100
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

func (r PokemonMongoRepository) FindByFilters(ctx context.Context, in dto.Filters) ([]model.Pokemon, error) {
	filters := r.buildFilters(in)
	opts := options.Find()
	opts.SetLimit(listDefaultLimit)
	cursor, err := r.collection.Find(ctx, filters, opts)
	if err != nil {
		return nil, err
	}

	var pokemons []model.Pokemon
	err = cursor.All(ctx, &pokemons)

	return pokemons, err
}

func (r PokemonMongoRepository) buildExactFilter(in dto.FilterRequest) bson.E {
	return bson.E{Key: in.Key, Value: in.Value}
}

func (r PokemonMongoRepository) buildRegexFilter(in dto.FilterRequest) bson.E {
	return bson.E{
		Key: in.Key,
		Value: bson.D{
			{Key: "$regex", Value: primitive.Regex{Pattern: in.Value, Options: "i"}},
		},
	}
}

func (r PokemonMongoRepository) buildFilters(in dto.Filters) bson.D {
	filters := bson.D{}

	for _, dtoFilter := range in.Filters {
		switch dtoFilter.Type {
		case dto.FilterExactType:
			filters = append(filters, r.buildExactFilter(dtoFilter))
		case dto.FilterRegexType:
			filters = append(filters, r.buildRegexFilter(dtoFilter))
		}
	}

	return filters
}
