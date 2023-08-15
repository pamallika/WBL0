package handler

import (
	"github.com/gin-gonic/gin"
	wbl0 "github.com/pamallika/WBL0/internal/core"
	"log"
	"net/http"
)

func (h *Handler) getOrderById(c *gin.Context) (int, error) {
	return 0, nil
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var input wbl0.Order
	//fmt.Printf("%+v\n", c.Request.Body)
	if err := c.BindJSON(&input); err != nil {
		log.Fatalf("error validate data: %s", err)
	}
	if (input.Delivery != wbl0.Delivery{}) {
		_, err := h.services.Order.CreateDelivery(input.Delivery)
		if err != nil {
			log.Fatalf("error create delivery: %s", err)
		}
	}
	if (input.Payment != wbl0.Payment{}) {
		_, err := h.services.Order.CreatePayment(input.Payment)
		if err != nil {
			log.Fatalf("error create payment: %s", err)
		}
	}
	//log.Println(input.Items[] == wbl0.Item{})
	//fmt.Printf("%+v\n", input)
	id, err := h.services.Order.CreateOrder(input)
	if err != nil {
		log.Fatalf("Error creating order: %s", err)
	}
	for i := 0; i < len(input.Items); i++ {
		input.Items[i].OrderId = id
		_, err = h.services.Order.CreateItem(input.Items[i])
		if err != nil {
			log.Fatalf("error create items: %s", err)
		}
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

//func CreateDelivery(delivery wbl0.Delivery) {
//}
