package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

const (
	// UserCollection Auth
	UserAuthCollection = "userauth"
	// Token Auth
	UserTokenCollection = "usertoken"
)

type UserStatus string

const (
	UserStatusActive UserStatus = "active"
)

// User
type UserAuth struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string        `bson:"username"`
	Email    string        `bson:"email"`
	Phone    string        `bson:"phone"`
	Password string        `bson:"password"`
	Status   UserStatus    `bson:"status"`
	TimestampLog
}

// UserToken
type UserToken struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	UserID      bson.ObjectId `bson:"user_id"`
	Username    string        `bson:"username"`
	DeviceID    string        `bson:"device_id"`
	LoginDate   *time.Time    `bson:"login_date"`
	ExpiredDate *time.Time    `bson:"expired_date"`
	Token       string        `bson:"token"`
}
