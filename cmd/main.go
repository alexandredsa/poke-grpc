package main

import (
	"context"
	"log"
	"net"
	"os"
	"strings"

	"github.com/alexandredsa/poke-grpc/internal/container"
	pokeGrpc "github.com/alexandredsa/poke-grpc/pkg/domains/pokemon/transport/grpc"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := loadDotEnv(); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logrus.WithField("port", "9000").Info("TCP Listening...")

	dep := container.Dependencies{}
	if err := dep.Setup(ctx); err != nil {
		log.Fatal(err)
	}

	pokeService := pokeGrpc.NewPokemonService(dep.Providers.PokemonRepository)

	srv := grpc.NewServer()
	pokeService.Register(srv)
	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func loadDotEnv() error {
	env := os.Getenv("ENVIRONMENT")
	if strings.ToUpper(env) != "DEVELOPMENT" && strings.ToUpper(env) != "" {
		return nil
	}

	err := godotenv.Load(".env")
	return err
}
