package models

type Item struct {
	ItemID      uint   `gorm:"primaryKey" json:"-"`
	ItemCode    string `gorm:"foreignKey" json:"item_code"`
	Description string `json:"description" form:"description"`
	Quantity    uint   `json:"quantity" form:"quantity"`
	OrderID     uint   `json:"-"`
}
