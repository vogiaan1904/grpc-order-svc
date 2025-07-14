package mongo

import (
	"time"

	"github.com/vogiaan1904/order-svc/internal/repositories"
	"github.com/vogiaan1904/order-svc/pkg/log"
	"github.com/vogiaan1904/order-svc/pkg/mongo"
)

type implOrderRepository struct {
	l     log.Logger
	db    mongo.Database
	clock func() time.Time
}

func NewOrderRepo(l log.Logger, db mongo.Database) repositories.OrderRepository {
	return &implOrderRepository{
		l:     l,
		db:    db,
		clock: time.Now,
	}
}
