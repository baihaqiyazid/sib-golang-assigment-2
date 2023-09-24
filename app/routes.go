package app

import (
	"assignment-2/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func Route(orderController controller.OrderController){
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	
	router.POST("/api/order/create", orderController.CreateOrder)
	router.GET("/api/orders", orderController.GetAllOrder)
	router.PUT("/api/order/:id", orderController.UpdateOrder)
	router.DELETE("/api/order/:id", orderController.DeleteOrder)

	router.Run(":4001")
}