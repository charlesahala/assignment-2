package models

import "time"

type Orders struct {
	OrderID uint `gorm:"primaryKey" json:"orderID"`
	CustomerName string `gorm:"not null type:VARCHAR(50)" json:"customerName"`
	OrderedAt time.Time `gorm:"not null" json:"orderedAt"`
	Items []Items `json:"items"`
}