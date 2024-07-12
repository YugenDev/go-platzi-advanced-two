package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/YugenDev/go-platzi-advanced-two/database"
	"github.com/YugenDev/go-platzi-advanced-two/repository"
	"github.com/YugenDev/go-platzi-advanced-two/websocket"
	"github.com/gorilla/mux"
)

type Config struct {
	Port         string
	JWTSecretKey string
	DatabaseURL  string
}

type Server interface {
	Config() *Config
	Hub() *websocket.Hub
}

type Broker struct {
	config *Config
	router *mux.Router
	hub    *websocket.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecretKey == "" {
		return nil, errors.New("jwt secret is required")
	}
	if config.DatabaseURL == "" {
		return nil, errors.New("database url is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websocket.NewHub(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)

	repo, err := database.NewPostgresRepository(b.config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	go b.hub.Run()
	repository.SetRepository(repo)
	log.Println("starting server on port", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Println("error starting server:", err)
	} else {
		log.Fatalf("server stopped")
	}

}
