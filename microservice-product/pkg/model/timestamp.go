package model

import "time"

// TimestampLog to Insert CreatedAt & UpdatedAt
type TimestampLog struct {
	CreatedDt *time.Time `bson:"created_at"`
	UpdatedDt *time.Time `bson:"updated_at"`
}
