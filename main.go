package main

import (
	"assignment-2/app"
	"assignment-2/controller"
	"assignment-2/repository"
	"assignment-2/service"
)

func main()  {
	app.StartDB()

	db := app.GetDB()

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(db, orderRepository)
	orderController := controller.NewOrderController(orderService)

	app.Route(orderController)

}