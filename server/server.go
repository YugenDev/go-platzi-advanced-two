package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/YugenDev/go-platzi-advanced-two/database"
	"github.com/YugenDev/go-platzi-advanced-two/repository"
	"github.com/gorilla/mux"
)

type Config struct {
	Port         string
	JWTSecretKey string
	DatabaseURL  string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	ConfigFile *Config
	Router     *mux.Router
}

func (b *Broker) Config() *Config {
	return b.ConfigFile
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecretKey == "" {
		return nil, errors.New("secret is required")
	}

	if config.DatabaseURL == "" {
		return nil, errors.New("database url is required")
	}

	broker := &Broker{
		ConfigFile: config,
		Router:     mux.NewRouter(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.Router = mux.NewRouter()
	binder(b, b.Router)

	repo, err := database.NewPostgresRepository(b.ConfigFile.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	repository.SetRepository(repo)

	log.Println("Initializing server on port: ", b.Config().Port)
	if err := http.ListenAndServe(b.ConfigFile.Port, b.Router); err != nil {
		log.Println("Failed to start server: ", err)
	}

}
