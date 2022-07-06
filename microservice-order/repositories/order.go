package repositories

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gocommerce/helpers"
	"gocommerce/microservice-order/pkg/model"
)

type OrderRepo struct {
	db  *mongo.Database
	col *mongo.Collection
	log *logrus.Logger
}

// Repo Order
func NewOrderRepo(db *mongo.Database, log *logrus.Logger) *OrderRepo {
	return &OrderRepo{
		db:  db,
		col: db.Collection(model.OrderCollection),
		log: log,
	}
}

// Create new Order
func (r *OrderRepo) Create(c context.Context, t *model.Order) (model.Order, error) {
	r.log.Debugf("Creating (%s)", t)
	ctx, cancel := TimeOutContextWithParent(c)
	defer cancel()

	var (
		order model.Order
		now   = bson.Now()
	)

	t.TimestampLog.CreatedDt = &now
	t.UserID = helpers.GetUserID(c)
	res, err := r.col.InsertOne(ctx, t)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return order, err
	}

	err = r.col.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&order)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return order, err
	}

	return order, err
}

// FindByID
func (r *OrderRepo) FindByID(id string) (model.Order, error) {
	r.log.Debugf("Find By ID (%s) \n", id)
	ctx, cancel := TimeOutContext()
	defer cancel()

	var order model.Order
	oid, _ := primitive.ObjectIDFromHex(id)
	err := r.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&order)
	if err != nil {
		r.log.Println(err)
		return order, err
	}

	return order, nil
}

// FindByUserID
func (r *OrderRepo) FindByUserID(userID string) ([]model.Order, error) {
	r.log.Debugf("Find By UserID (%s)\n", userID)
	ctx, cancel := TimeOutContext()
	defer cancel()

	var orders []model.Order
	opts := options.Find()
	opts.SetSort(bson.M{"no": 1})

	rcur, err := r.col.Find(ctx, bson.M{"user_id": userID}, opts)
	if err != nil {
		r.log.Println(err)
		return orders, err
	}
	defer rcur.Close(ctx)
	for rcur.Next(ctx) {
		var order model.Order
		err := rcur.Decode(&order)
		if err != nil {
			r.log.Errorf(ErrMgoOpsFail, err)
		}
		orders = append(orders, order)
	}
	if err := rcur.Err(); err != nil {
		r.log.Error(err)
	}

	return orders, nil
}
