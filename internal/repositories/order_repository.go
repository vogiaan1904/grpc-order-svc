package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/vogiaan1904/order-svc/internal/models"
	"github.com/vogiaan1904/order-svc/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type implOrderRepository struct {
	db    mongo.Database
	clock func() time.Time
}

func NewOrderRepository(db mongo.Database) OrderRepository {
	return &implOrderRepository{
		db:    db,
		clock: time.Now,
	}
}

const (
	ordersCollection = "orders"
)

func (r *implOrderRepository) getCollection() mongo.Collection {
	return r.db.Collection(ordersCollection)
}

func (r *implOrderRepository) CreateOrder(ctx context.Context, opt CreateOrderOptions) (models.Order, error) {
	col := r.getCollection()
	o := r.buildOrderModel(opt)

	_, err := col.InsertOne(ctx, o)
	if err != nil {
		fmt.Println("Error inserting order:", err)
		return models.Order{}, err
	}

	return o, nil
}

func (r *implOrderRepository) FindManyOrders(ctx context.Context, opt FindManyOrderOptions) ([]models.Order, error) {
	col := r.getCollection()

	ft, err := r.buildFindQuery(opt.FindFilter)
	if err != nil {
		fmt.Println("Error building find query:", err)
		return []models.Order{}, err
	}

	cur, err := col.Find(ctx, ft, options.Find().
		SetSort(bson.D{
			{Key: "created_at", Value: -1},
			{Key: "_id", Value: -1},
		}))
	if err != nil {
		fmt.Println("Error Find:", err)
		return []models.Order{}, err
	}
	defer cur.Close(ctx)

	var os []models.Order
	err = cur.All(ctx, &os)
	if err != nil {
		fmt.Println("Error decoding orders:", err)
		return []models.Order{}, err
	}

	return os, nil
}
