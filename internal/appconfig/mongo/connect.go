package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vogiaan1904/order-svc/pkg/mongo"
)

const (
	ctxTimeout = 10 * time.Second
)

func Connect(uri string) (mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	client, err := mongo.NewClient(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to create mongo client: %w", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	err = client.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping to DB: %w", err)
	}

	log.Println("Connected to MongoDB.")

	return client, nil
}

func Disconnect(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
