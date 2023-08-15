package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	wbl0 "github.com/pamallika/WBL0/internal/core"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) CreateOrder(order wbl0.Order) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (uid, track_number, entry, internal_signature, customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard, locate) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id", ordersTable)
	row := r.db.QueryRow(query, order.OrderUID, order.TrackNumber, order.Entry, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SMID, order.DateCreated, order.OOFShard, order.Locale)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OrderPostgres) CreateDelivery(delivery wbl0.Delivery) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, phone, zip, city, address, region, email) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", deliveriesTable)
	row := r.db.QueryRow(query, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OrderPostgres) CreateItem(item wbl0.Item) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_id) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id", itemsTable)
	row := r.db.QueryRow(query, item.ChrtID, item.TrackNumber, item.Price, item.RID, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status, item.OrderId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OrderPostgres) CreatePayment(payment wbl0.Payment) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id", paymentsTable)
	row := r.db.QueryRow(query, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDT, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
