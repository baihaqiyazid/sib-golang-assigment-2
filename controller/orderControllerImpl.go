package controller

import (
	"assignment-2/helper"
	"assignment-2/service"
	"assignment-2/web"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(service service.OrderService) *OrderControllerImpl {
	return &OrderControllerImpl{OrderService: service}
}

func (controller *OrderControllerImpl) CreateOrder(ctx *gin.Context) {

	var request web.OrderPayload
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	data, err := controller.OrderService.CreateOrder(ctx, request)
	if err != nil{
		helper.ResponseBadRequest(ctx, err)
	}
	

	helper.ResponseSuccess(ctx, data)

	return
}

func (controller *OrderControllerImpl) UpdateOrder(ctx *gin.Context) {

	var request web.OrderPayload
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	// Check Order by Id
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := controller.OrderService.GetOrderById(ctx, id)
	if err != nil{
		fmt.Println(err)
		helper.ResponseNotFound(ctx, "Data Not Found")
		panic(err)
	}

	data, err := controller.OrderService.UpdateOrder(ctx, request)
	if err != nil{
		helper.ResponseBadRequest(ctx, err)
	}
	

	helper.ResponseSuccess(ctx, data)

	return
}

func (controller *OrderControllerImpl) GetAllOrder(ctx *gin.Context)  {
	data := controller.OrderService.GetAll(ctx)

	helper.ResponseSuccess(ctx, data)

	return
}

func (controller *OrderControllerImpl) DeleteOrder(ctx *gin.Context) {


	// Check Order by Id
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := controller.OrderService.GetOrderById(ctx, id)
	if err != nil{
		fmt.Println(err)
		helper.ResponseNotFound(ctx, "Data Not Found")
		panic(err)
	}
	
	controller.OrderService.DeleteOrder(ctx, id)	

	helper.ResponseSuccess(ctx, nil)

	return
}