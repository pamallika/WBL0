package core

import (
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid" db:"uid"`
	TrackNumber       string    `json:"track_number" db:"track_number"`
	Entry             string    `json:"entry" db:"entry"`
	Delivery          Delivery  `json:"delivery" db:"delivery" marshal:"1"`
	Payment           Payment   `json:"payment" db:"payment" marshal:"2"`
	Items             []Item    `json:"items" db:"items" marshal:"3"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerID        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	ShardKey          string    `json:"shardkey" db:"shard_key"`
	SMID              int       `json:"sm_id" db:"sm_id"`
	DateCreated       time.Time `json:"date_created" db:"date_created"`
	OOFShard          string    `json:"oof_shard" db:"oof_shard"`
}
