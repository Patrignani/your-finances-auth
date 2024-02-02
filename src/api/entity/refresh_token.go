package entity

import "time"

type RefreshToken struct {
	ID             string    `bson:"_id,omitempty" json:"id,omitempty"`
	ExpirationDate time.Time `bson:"expiration_date" json:"expiration_date"`
	UserID         string    `bson:"user_id" json:"user_id"`
	Active         bool      `bson:"active" json:"active"`
	CreateAt       time.Time `bson:"create_at" json:"create_at"`
	UpdateAt       time.Time `bson:"update_at,omitempty" json:"update_at,omitempty"`
}
