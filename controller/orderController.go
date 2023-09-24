package controller

import "github.com/gin-gonic/gin"

type OrderController interface{
	CreateOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	GetAllOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}
