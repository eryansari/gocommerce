package repositories

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gocommerce/microservice-authentication/pkg/model"
)

type UserAuthRepo struct {
	db  *mongo.Database
	col *mongo.Collection
	log *logrus.Logger
}

// Repo Auth
func NewUserAuthRepo(db *mongo.Database, log *logrus.Logger) *UserAuthRepo {
	return &UserAuthRepo{
		db:  db,
		col: db.Collection(model.UserAuthCollection),
		log: log,
	}
}

// Create new User
func (r *UserAuthRepo) Create(c context.Context, t *model.UserAuth) (model.UserAuth, error) {
	r.log.Debugf("Creating (%s)", t)
	ctx, cancel := TimeOutContextWithParent(c)
	defer cancel()

	var userAuth model.UserAuth
	now := bson.Now()
	t.TimestampLog.CreatedDt = &now
	res, err := r.col.InsertOne(ctx, t)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return userAuth, err
	}

	err = r.col.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&userAuth)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return userAuth, err
	}

	r.log.Debugf("Saved User (%v)", userAuth)

	return userAuth, nil
}

// Create Token
func (r *UserAuthRepo) CreateToken(c context.Context, t *model.UserToken) (model.UserToken, error) {
	r.log.Debugf("Creating Token (%s)", t)
	ctx, cancel := TimeOutContextWithParent(c)
	defer cancel()

	col := r.db.Collection(model.UserTokenCollection)

	var utoken model.UserToken
	res, err := col.InsertOne(ctx, t)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return utoken, err
	}

	err = col.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&utoken)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return utoken, err
	}

	r.log.Debugf("Saving Token (%v)", utoken)

	return utoken, err
}

// Delete Token By User & DeviceID
func (r *UserAuthRepo) DeleteTokenByUserAndDevice(c context.Context, username, deviceID string) (bool, error) {
	r.log.Debugf("Delete Token By User & DeviceID (%s,%s)", username, deviceID)
	ctx, cancel := TimeOutContextWithParent(c)
	defer cancel()

	col := r.db.Collection(model.UserTokenCollection)
	filter := bson.M{
		"username":  username,
		"device_id": deviceID,
	}

	res, err := col.DeleteOne(ctx, filter)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return false, err
	}

	return res.DeletedCount > 0, nil
}

// Find Token Not Expired
func (r *UserAuthRepo) FindTokenByUserAndDevice(c context.Context, username, deviceID string) (model.UserToken, error) {
	r.log.Debugf("Find Token By User & DeviceID (%s,%s)", username, deviceID)
	ctx, cancel := TimeOutContextWithParent(c)
	defer cancel()

	col := r.db.Collection(model.UserTokenCollection)
	filter := bson.M{
		"username":  username,
		"device_id": deviceID,
		"expired_date": bson.M{
			"$gt": bson.Now(),
		},
	}

	var utoken model.UserToken
	err := col.FindOne(ctx, filter).Decode(&utoken)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return utoken, err
	}

	return utoken, nil
}

// Find By Email
func (r *UserAuthRepo) FindByEmail(c context.Context, email string) (model.UserAuth, error) {
	r.log.Debugf("Find By Email (%s) \n", email)
	ctx, cancel := TimeOutContextWithParent(c)
	defer cancel()

	var userAuth model.UserAuth
	err := r.col.FindOne(ctx, bson.M{"email": email}).Decode(&userAuth)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return userAuth, err
	}

	return userAuth, nil
}

// Find By Phone
func (r *UserAuthRepo) FindByPhone(c context.Context, phone string) (model.UserAuth, error) {
	r.log.Debugf("Find By Phone (%s) \n", phone)
	ctx, cancel := TimeOutContextWithParent(c)
	defer cancel()

	var userAuth model.UserAuth
	err := r.col.FindOne(ctx, bson.M{"phone": phone}).Decode(&userAuth)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return userAuth, err
	}

	return userAuth, nil
}

// Find By Username
func (r *UserAuthRepo) FindByUsername(c context.Context, username string) (model.UserAuth, error) {
	r.log.Debugf("Find By Username (%s) \n", username)
	ctx, cancel := TimeOutContextWithParent(c)
	defer cancel()

	var userAuth model.UserAuth
	err := r.col.FindOne(ctx, bson.M{"username": username}).Decode(&userAuth)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return userAuth, err
	}

	return userAuth, nil
}

// Find By ID
func (r *UserAuthRepo) FindByID(c context.Context, id string) (model.UserAuth, error) {
	r.log.Debugf("Find By ID(%s) \n", id)
	ctx, cancel := TimeOutContextWithParent(c)
	defer cancel()

	var userAuth model.UserAuth
	oid, _ := primitive.ObjectIDFromHex(id)
	err := r.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&userAuth)
	if err != nil {
		r.log.Errorf(ErrMgoOpsFail, err)
		return userAuth, err
	}

	return userAuth, nil
}
