package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID uint `gorm:"index;not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`

	AddressID uint    `gorm:"not null" json:"address_id"`
	Address   Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address,omitempty"`

	TotalAmount     float64     `gorm:"type:decimal(10,2)" json:"total_amount"`
	Status          string      `gorm:"size:50;not null;default:'pending'" json:"status"`
	ShippingAddress string      `gorm:"size:255" json:"shipping_address"`
	OrderItems      []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
	Payment         Payment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"payment,omitempty"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null;index" json:"order_id"`
	ProductID uint    `gorm:"not null;index" json:"product_id"`
	Quantity  int     `gorm:"not null;default:1" json:"quantity"`
	Price     float64 `gorm:"type:decimal(10,2);not null" json:"price"` // product price at time of order

	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
}

type CreateOrderRequest struct {
	// UserID    uint `json:"user_id" validate:"required"`
	AddressID uint `json:"address_id" validate:"required"`
	Items     []struct {
		ProductID uint `json:"product_id" validate:"required"`
		Quantity  int  `json:"quantity" validate:"required,min=1"`
	} `json:"items" validate:"required,dive"`
}
