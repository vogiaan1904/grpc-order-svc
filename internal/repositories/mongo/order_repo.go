package mongo

import (
	"context"
	"time"

	"github.com/vogiaan1904/order-svc/internal/models"
	repo "github.com/vogiaan1904/order-svc/internal/repositories"
	"github.com/vogiaan1904/order-svc/pkg/log"
	"github.com/vogiaan1904/order-svc/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type implOrderRepository struct {
	l     log.Logger
	db    mongo.Database
	clock func() time.Time
}

func NewOrderRepository(l log.Logger, db mongo.Database) repo.OrderRepository {
	return &implOrderRepository{
		l:     l,
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

func (r *implOrderRepository) CreateOrder(ctx context.Context, opt repo.CreateOrderOptions) (models.Order, error) {
	col := r.getCollection()
	o := r.buildOrderModel(opt)

	_, err := col.InsertOne(ctx, o)
	if err != nil {
		r.l.Errorf(ctx, "Error InsertOne: %v", err)
		return models.Order{}, err
	}

	return o, nil
}

func (r *implOrderRepository) GetOrders(ctx context.Context, opts repo.GetOrdersOptions) ([]models.Order, error) {
	col := r.getCollection()

	ft := r.buildFindQuery(opts.GetOrdersFilter)
	r.l.Debugf(ctx, "filter: %-v", ft)
	cur, err := col.Find(ctx, ft, options.Find().
		SetSort(bson.D{
			{Key: "created_at", Value: -1},
			{Key: "_id", Value: -1},
		}))
	if err != nil {
		r.l.Errorf(ctx, "Error Find: %v", err)
		return []models.Order{}, err
	}
	defer cur.Close(ctx)

	var os []models.Order
	err = cur.All(ctx, &os)
	if err != nil {
		r.l.Errorf(ctx, "Error cur.All: %v", err)
		return []models.Order{}, err
	}

	return os, nil
}

func (r *implOrderRepository) FindOneOrder(ctx context.Context, opt repo.FindOneOrderOptions) (models.Order, error) {
	col := r.getCollection()

	var o models.Order

	ft := r.buildFindOneQuery(opt)
	err := col.FindOne(ctx, ft).Decode(&o)
	if err != nil {
		r.l.Errorf(ctx, "Error FindOne: %v", err)
		return models.Order{}, err
	}

	return o, nil
}

func (r *implOrderRepository) UpdateOrder(ctx context.Context, o models.Order) error {
	col := r.getCollection()

	_, err := col.UpdateOne(ctx, bson.M{"_id": o.ID}, bson.M{"$set": o})
	if err != nil {
		r.l.Errorf(ctx, "Error UpdateOne: %v", err)
		return err
	}

	return nil
}
