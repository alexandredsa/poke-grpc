package container

import (
	"context"
	"os"
	"time"

	"github.com/alexandredsa/poke-grpc/internal/components/mongo"
	pokeRepository "github.com/alexandredsa/poke-grpc/pkg/domains/pokemon/repository"
)

type Components struct {
	MongoClient mongo.Client
}

type Providers struct {
	PokemonRepository pokeRepository.PokemonRepository
}

type Dependencies struct {
	Components Components
	Providers  Providers
}

func (d Dependencies) Setup(ctx context.Context) error {

	credentials := mongo.Credentials{
		ConnectionTimeout: 10 * time.Second,
		DBName:            os.Getenv("MONGO_DB_NAME"),
		URI:               os.Getenv("MONGO_URI"),
	}

	client, err := mongo.NewClient(ctx, credentials)
	if err != nil {
		return err
	}

	pokemonRepository := pokeRepository.NewPokemonMongoRepository(client)

	d.Components = Components{
		MongoClient: *client,
	}

	d.Providers = Providers{
		PokemonRepository: pokemonRepository,
	}

	return nil
}
