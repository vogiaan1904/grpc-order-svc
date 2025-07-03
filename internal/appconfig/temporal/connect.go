package temporal

import (
	"fmt"
	"log"

	"github.com/vogiaan1904/order-svc/config"
	"go.temporal.io/sdk/client"
)

func Connect(cfg config.TemporalConfig) (client.Client, error) {
	tCli, err := client.Dial(client.Options{
		HostPort:  cfg.HostPort,
		Namespace: cfg.Namespace,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create Temporal Client: %w", err)
	}
	defer tCli.Close()

	log.Println("Connected to Temporal!")

	return tCli, nil
}

func Disconnect(tCli client.Client) {
	if tCli == nil {
		return
	}

	tCli.Close()
	log.Println("Connection to Temporal closed.")
}
