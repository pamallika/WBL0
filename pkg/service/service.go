package service

import (
	"github.com/pamallika/WBL0/internal/core"
	"github.com/pamallika/WBL0/pkg/repository"
)

type Order interface {
	CreateOrder(order core.Order) (int, error)
	CreateDelivery(delivery core.Delivery) (int, error)
	CreatePayment(payment core.Payment) (int, error)
	CreateItem(item core.Item) (int, error)
	GetById(order core.Order) (int, error)
}
type Delivery interface {
}
type Item interface {
	CreateItem(item core.Item) (int, error)
}
type Payment interface {
	CreatePayment(payment core.Payment) (int, error)
}

type Service struct {
	Order
	Delivery
	Item
	Payment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order),
	}
}
