package models

import (
	"time"

	"github.com/google/uuid"
)

type Logistics struct {
	ID               uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID           uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	CompanyName       string `json:"company_name,omitempty"`
	VehicleType       string `json:"vehicle_type" binding:"required"`
	ServiceAreas      string `json:"service_areas" binding:"required"`
	VerificationID    string `json:"verification_id" binding:"required"`
	PhoneNumber       string `json:"phone_number" binding:"required"`
	AlternativeContact string `json:"alternative_contact,omitempty"`
	PaymentMethod     string `json:"payment_method" binding:"required"`
	AvailabilityStatus bool   `json:"availability_status" binding:"required"`
	Photos  		   string    `json:"photos"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}
