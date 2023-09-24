package service

import (
	"assignment-2/models"
	"assignment-2/web"

	"github.com/gin-gonic/gin"
)


type OrderService interface {
	CreateOrder(ctx *gin.Context, request web.OrderPayload) (models.Order, error)
	GetAll(ctx *gin.Context) []models.Order
	UpdateOrder(ctx *gin.Context, request web.OrderPayload) (models.Order, error)
	GetOrderById(ctx *gin.Context, id int) error
	DeleteOrder(ctx *gin.Context, id int) error
}