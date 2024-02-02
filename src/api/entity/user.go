package entity

import "time"

type User struct {
	ID             string    `bson:"_id,omitempty" json:"id,omitempty"`
	Username       string    `bson:"username" json:"username"`
	Password       string    `bson:"password" json:"password"`
	Seed           string    `bson:"seed" json:"seed"`
	Roles          []string  `bson:"roles" json:"roles"`
	Permissions    []string  `bson:"permissions" json:"permissions"`
	CreateAt       time.Time `bson:"create_at" json:"create_at"`
	UpdateAt       time.Time `bson:"update_at,omitempty" json:"update_at,omitempty"`
	CreateBy       string    `bson:"create_by,omitempty" json:"create_by,omitempty"`
	ClientCreateBy string    `bson:"client_create_by" json:"client_create_by"`
	UpdateBy       string    `bson:"update_by,omitempty" json:"update_by,omitempty"`
	ClientUpdateBy string    `bson:"client_update_by,omitempty" json:"client_update_by,omitempty"`
	Active         bool      `bson:"active" json:"active"`
	TwoFactorCode  string    `bson:"two_factory_code,omitempty" json:"two_factory_code,omitempty"`
}
