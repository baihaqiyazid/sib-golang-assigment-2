package service

import (
	"assignment-2/helper"
	"assignment-2/models"
	"assignment-2/repository"
	"assignment-2/web"
	"log"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderServiceImpl struct {
	DB              *gorm.DB
	OrderRepository repository.OrderRepository
}

func NewOrderService(db *gorm.DB, OrderRepository repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{
		DB:              db,
		OrderRepository: OrderRepository,
	}
}

func (service *OrderServiceImpl) CreateOrder(ctx *gin.Context, request web.OrderPayload) (models.Order, error) {
	var order models.Order

	order.CustomerName = request.CustomerName
	order.OrderedAt = request.OrderedAt
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	for _, itemPayload := range request.Items {
		item := models.Item{
			ItemCode:    itemPayload.ItemCode,
			Description: itemPayload.Description,
			Quantity:    itemPayload.Quantity,
		}

		order.Items = append(order.Items, item)
	}

	data, err := service.OrderRepository.Create(order)
	helper.LogIfError(err)

	return data, nil
}

func (service *OrderServiceImpl) GetAll(ctx *gin.Context) []models.Order {
	data := service.OrderRepository.GetAll()

	return data
}

func (service *OrderServiceImpl) UpdateOrder(ctx *gin.Context, request web.OrderPayload) (models.Order, error) {
	var order models.Order

	id := ctx.Param("id")
	orderId, _ := strconv.Atoi(id)

	// Mulai transaksi
	tx := service.DB.Begin()
	if tx.Error != nil {
		return models.Order{}, tx.Error
	}

	order.CustomerName = request.CustomerName
	order.UpdatedAt = time.Now()

	itemResponse, err := service.OrderRepository.GetItemById(orderId)
	if err != nil {
		tx.Rollback()
		return models.Order{}, err
	}

	var itemToUpdate models.Item

	for i, item := range *itemResponse {

		itemToUpdate.ItemID = item.ItemID
		itemToUpdate.ItemCode = request.Items[i].ItemCode
		itemToUpdate.Description = request.Items[i].Description
		itemToUpdate.Quantity = request.Items[i].Quantity

		log.Println("itemToUpdate: ", itemToUpdate)

		_, err := service.OrderRepository.UpdateItem(itemToUpdate, int(itemToUpdate.ItemID))
		if err != nil {
			tx.Rollback()
			return models.Order{}, err
		}

	}

	_, err = service.OrderRepository.UpdateOrder(order, orderId)
	if err != nil {
		tx.Rollback()
		return models.Order{}, err
	}

	tx.Commit()

	data, err := service.OrderRepository.GetOrderById(orderId)
	if err != nil {
		return models.Order{}, err
	}

	return data, nil
}



func (service *OrderServiceImpl) GetOrderById(ctx *gin.Context, id int) error {
	_, err := service.OrderRepository.GetOrderById(id)
	if err != nil {
		return err
	}

	return nil
}


func (service *OrderServiceImpl) DeleteOrder(ctx *gin.Context, id int) error {
	
	service.OrderRepository.DeleteItem(id)
	service.OrderRepository.DeleteOrder(id)
	
	return nil
}
