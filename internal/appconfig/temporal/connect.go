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
		return nil, fmt.Errorf("failed to initialize Temporal client: %v", err)
	}

	log.Printf("Temporal Client connected.")
	return tCli, nil
}

func Disconnect(cli client.Client) {
	if cli == nil {
		return
	}

	cli.Close()
	log.Printf("Temporal Client connection closed.")
}
