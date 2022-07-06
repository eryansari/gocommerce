package model

import (
	"github.com/globalsign/mgo/bson"
)

const (
	OrderCollection = "order"
)

// Order
type Order struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	UserID      string        `bson:"user_id"`
	NoInvoice   string        `bson:"no_inv"`
	Status      bool          `bson:"status"`
	OrderDetail []OrderDetail
	TimestampLog
}

// OrderDetail
type OrderDetail struct {
	ProductID string `bson:"product_id"`
	Qty       int64  `bson:"qty"`
}
