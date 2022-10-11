package models

import "time"

type Order struct {
	OrderID      uint      `gorm:"primaryKey" json:"-" form:"order_id" valid:"required~OrderID is required"`
	OrderedAt    time.Time `gorm:"not null;type:timestamp;autoCreateTime" json:"orderedAt" valid:"required"`
	CustomerName string    `gorm:"not null" json:"customer_name" form:"customer_name" valid:"required~Customer Name is required"`
	Items        []Item
}
