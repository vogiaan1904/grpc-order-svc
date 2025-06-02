package repository

import (
	"github.com/vogiaan1904/order-svc/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *implOrderRepository) buildFindQuery(filter FindFilter) bson.M {
	ft := bson.M{}
	ft = mongo.BuildQueryWithSoftDelete(ft)

	if filter.UserID != "" {
		ft["user_id"] = filter.UserID
	}

	if filter.Status != nil {
		ft["status"] = filter.Status
	}

	return ft
}

func (r *implOrderRepository) buildFindOneQuery(opt FindOneOrderOptions) bson.M {
	ft := r.buildFindQuery(opt.FindFilter)

	if opt.ID != "" {
		ft["_id"] = opt.ID
	}

	if opt.Code != "" {
		ft["code"] = opt.Code
	}

	return ft
}
