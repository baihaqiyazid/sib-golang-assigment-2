package repository

import (
	"assignment-2/helper"
	"assignment-2/models"


	"gorm.io/gorm"
)

type OrderRepositoryImpl struct{
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		db: db,
	}
}

func (repository *OrderRepositoryImpl) Create(order models.Order) (models.Order, error){
	err := repository.db.Create(&order).Error
	helper.LogIfError(err)

	return order, nil
}

func (repository *OrderRepositoryImpl) GetAll() []models.Order{
	var orders []models.Order

	err := repository.db.Preload("Items").Find(&orders).Error
	helper.LogIfError(err)

	return orders
}

func (repository *OrderRepositoryImpl) UpdateOrder(order models.Order, id int) (models.Order, error) {
	err := repository.db.Where("id = ?", id).Updates(&order).Error
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (repository *OrderRepositoryImpl) UpdateItem(item models.Item, id int) (*models.Item, error) {
	err := repository.db.Where("item_id = ?", id).Updates(item).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}


func (repository *OrderRepositoryImpl) GetOrderById(id int) (models.Order, error) {
	var order models.Order
	err := repository.db.Where("id", id).Preload("Items").First(&order).Error
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (repository *OrderRepositoryImpl) GetItemById(id int) (*[]models.Item, error) {
	var item []models.Item
	err := repository.db.Where("order_id", id).Find(&item).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (repository *OrderRepositoryImpl) DeleteOrder(id int) {
	var order models.Order
	repository.db.Where("id", id).Delete(&order)
}

func (repository *OrderRepositoryImpl) DeleteItem(id int) {
	var item models.Item
	repository.db.Where("order_id", id).Delete(&item)
}


