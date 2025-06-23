package mongo

import (
	repo "github.com/vogiaan1904/order-svc/internal/repositories"
	"github.com/vogiaan1904/order-svc/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *implOrderRepository) buildFindQuery(filter repo.GetOrdersFilter) bson.M {
	ft := bson.M{}
	ft = mongo.BuildQueryWithSoftDelete(ft)

	if filter.UserID != "" {
		ft["user_id"] = filter.UserID
	}

	if filter.Status != "" {
		ft["status"] = filter.Status
	}

	return ft
}

func (r *implOrderRepository) buildFindOneQuery(opts repo.FindOneOrderOptions) bson.M {
	ft := r.buildFindQuery(opts.GetOrdersFilter)

	if opts.ID != "" {
		ft["_id"] = opts.ID
	}

	if opts.Code != "" {
		ft["code"] = opts.Code
	}

	return ft
}
