package repository

import (
	"github.com/vogiaan1904/order-svc/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *implOrderRepository) buildFindQuery(filter FindFilter) (bson.M, error) {
	ft := bson.M{}
	ft = mongo.BuildQueryWithSoftDelete(ft)

	if filter.UserID != "" {
		ft["user_id"] = filter.UserID
	}

	if filter.Status != "" {
		ft["status"] = filter.Status
	}

	return ft, nil
}
