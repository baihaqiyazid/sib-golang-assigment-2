package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `gorm:"type:varchar(255)" json:"customer_name"`
	Items        []Item    `json:"items"`
	OrderedAt    time.Time `json:"ordered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
