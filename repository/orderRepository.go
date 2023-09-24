package repository

import (
	"assignment-2/models"
)

type OrderRepository interface{
	Create(item models.Order) (models.Order, error)
	GetAll() []models.Order
	UpdateOrder(order models.Order, id int) (models.Order, error)
	UpdateItem(item models.Item, id int) (*models.Item, error)
	GetOrderById(id int) (models.Order, error)
	GetItemById(id int) (*[]models.Item, error)
	DeleteOrder(id int) 
	DeleteItem(id int)
}