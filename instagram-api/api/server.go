package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	client *mongo.Client
}

func NewServer() *Server {

	clientOptions := options.Client().
		ApplyURI("mongodb+srv://bis123:bis123@cluster0.g6knc.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return &Server{client}
}

func (s *Server) StartServer(port string) {
	fmt.Printf("Starting a server on port: %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
