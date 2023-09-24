package models

import "time"

type Item struct {
	ItemID      uint      `gorm:"primaryKey" json:"item_id"`
	ItemCode    string    `gorm:"not null;type:varchar(255)" json:"item_code"`
	Description string    `gorm:"type:varchar" json:"description"`
	Quantity    uint      `gorm:"type:integer" json:"quantity"`
	OrderID     uint      `json:"order_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
