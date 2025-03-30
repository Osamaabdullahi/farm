package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID       uuid.UUID    `gorm:"type:uuid;not null" json:"user_id"`
	User         User         `gorm:"foreignKey:UserID" json:"user"`
	TotalAmount  float64      `gorm:"not null" json:"total_amount"`
	Status       string       `gorm:"default:'pending'" json:"status"` // pending, completed, canceled
	PaymentMethod string      `gorm:"not null" json:"payment_method"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Items        []OrderItem  `gorm:"foreignKey:OrderID" json:"items"`
	FarmerID   uuid.UUID    `gorm:"type:uuid;not null" json:"farmer_id"`
}

type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null" json:"order_id"`
	// Order     Order     `gorm:"foreignKey:OrderID" json:"order"`
	ProduceID uuid.UUID `gorm:"type:uuid;not null" json:"produce_id"`
	Produce   Produce   `gorm:"foreignKey:ProduceID" json:"produce"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"not null" json:"price"`
	TotalCost float64   `gorm:"not null" json:"total_cost"`
}

