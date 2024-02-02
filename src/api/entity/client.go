package entity

import "time"

type Client struct {
	ID             string    `bson:"_id,omitempty" json:"id,omitemptys"`
	ClientID       string    `bson:"client_id" json:"client_id"`
	ClientSecret   string    `bson:"client_secret" json:"client_secret"`
	Name           string    `bson:"name" json:"name"`
	Description    string    `bson:"description" json:"description"`
	CreateAt       time.Time `bson:"create_at" json:"create_at"`
	UpdateAt       time.Time `bson:"update_at,omitempty" json:"update_at,omitemptys"`
	CreateBy       string    `bson:"create_by,omitempty" json:"create_by,omitempty"`
	ClientCreateBy string    `bson:"client_create_by" json:"client_create_by"`
	UpdateBy       string    `bson:"update_by,omitempty" json:"updatse_by,omitempty"`
	ClientUpdateBy string    `bson:"client_update_by,omitempty" json:"client_update_by,omitempty"`
	Active         bool      `bson:"active" json:"active"`
}
