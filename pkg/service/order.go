package service

import (
	wbl0 "github.com/pamallika/WBL0/internal/core"
	"github.com/pamallika/WBL0/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order wbl0.Order) (int, error) {
	return s.repo.CreateOrder(order)
}
func (s *OrderService) CreateDelivery(delivery wbl0.Delivery) (int, error) {
	return s.repo.CreateDelivery(delivery)
}
func (s *OrderService) CreatePayment(payment wbl0.Payment) (int, error) {
	return s.repo.CreatePayment(payment)
}
func (s *OrderService) CreateItem(item wbl0.Item) (int, error) {
	return s.repo.CreateItem(item)
}
func (s *OrderService) GetById(order wbl0.Order) (int, error) {
	return s.repo.CreateOrder(order)
}
