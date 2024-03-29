package repositories

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gocommerce/microservice-product/pkg/model"
)

type ProductRepo struct {
	db  *mongo.Database
	col *mongo.Collection
	log *logrus.Logger
}

// Repo Product
func NewProductRepo(db *mongo.Database, log *logrus.Logger) *ProductRepo {
	return &ProductRepo{
		db:  db,
		col: db.Collection(model.ProductCollection),
		log: log,
	}
}

// Create new Product
func (r *ProductRepo) Create(c context.Context, t *model.Product) (model.Product, error) {
	r.log.Debugf("Creating (%s)", t)
	ctx, cancel := TimeOutContextWithParent(c)
	defer cancel()

	var (
		product model.Product
		now     = bson.Now()
	)

	t.TimestampLog.CreatedDt = &now
	t.Status = model.ProductActive
	res, err := r.col.InsertOne(ctx, t)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return product, err
	}

	err = r.col.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&product)

	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return product, err
	}

	return product, err
}

// Find By ID
func (r *ProductRepo) FindByID(id string) (model.Product, error) {
	r.log.Debugf("Find By ID (%s) \n", id)
	ctx, cancel := TimeOutContext()
	defer cancel()

	var product model.Product
	oid, _ := primitive.ObjectIDFromHex(id)
	err := r.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&product)
	if err != nil {
		r.log.Println(err)
		return product, err
	}

	return product, nil
}

// FindAll products
func (r *ProductRepo) FindAll() ([]model.Product, error) {
	r.log.Debugf("Find All \n")
	ctx, cancel := TimeOutContext()
	defer cancel()

	var products []model.Product
	opts := options.Find()
	opts.SetSort(bson.M{"no": 1})

	rcur, err := r.col.Find(ctx, bson.M{}, opts)
	if err != nil {
		r.log.Println(err)
		return products, err
	}
	defer rcur.Close(ctx)
	for rcur.Next(ctx) {
		var product model.Product
		err := rcur.Decode(&product)
		if err != nil {
			r.log.Errorf(ErrMgoOpsFail, err)
		}
		products = append(products, product)
	}
	if err := rcur.Err(); err != nil {
		r.log.Error(err)
	}

	return products, nil
}
