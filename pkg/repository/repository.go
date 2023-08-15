package repository

import (
	"github.com/jmoiron/sqlx"
	wbl0 "github.com/pamallika/WBL0/internal/core"
)

type Order interface {
	CreateOrder(order wbl0.Order) (int, error)
	CreateDelivery(order wbl0.Delivery) (int, error)
	CreatePayment(order wbl0.Payment) (int, error)
	CreateItem(order wbl0.Item) (int, error)
}
type Delivery interface {
}
type Item interface {
}
type Payment interface {
}

type Repository struct {
	Order
	Delivery
	Item
	Payment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db),
	}
}
