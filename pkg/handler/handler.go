package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pamallika/WBL0/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h Handler) InitRouts() *gin.Engine {
	router := gin.New()
	orders := router.Group("/orders")
	{
		//orders.GET("/:id", h.getOrderById)
		orders.POST("/", h.CreateOrder)
	}
	return router
}
