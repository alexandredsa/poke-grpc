package reader

import (
	"os"

	"github.com/gocarina/gocsv"

	model "github.com/alexandredsa/poke-grpc/pkg/domains/pokemon/model"
)

type PokemonFileReader interface {
	Load(path string) ([]*model.Pokemon, error)
}

type PokemonCSVReader struct{}

func NewPokemonCSVReader() PokemonFileReader {
	return PokemonCSVReader{}
}

func (r PokemonCSVReader) Load(path string) ([]*model.Pokemon, error) {
	datasetFile, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer datasetFile.Close()

	pokemons := make([]*model.Pokemon, 0)

	if err := gocsv.UnmarshalFile(datasetFile, &pokemons); err != nil {
		return nil, err
	}

	return pokemons, nil
}
