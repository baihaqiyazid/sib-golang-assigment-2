package web

import (
	"assignment-2/models"
	"time"
)

type OrderPayload struct {
	OrderedAt    time.Time     `json:"ordered_at"`
	CustomerName string        `json:"customer_name"`
	Items        []models.Item `json:"items"`
}
