package models

import (
	"time"

	"github.com/google/uuid"
)

// Business model
type Business struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID             uuid.UUID `gorm:"type:uuid;unique;not null" json:"user_id"` // Link to User
	BusinessName       string    `gorm:"not null" json:"business_name"`
	BusinessType       string    `gorm:"not null" json:"business_type"` // e.g., Restaurant, Grocery Store
	BusinessLocation   string    `gorm:"not null" json:"business_location"`
	VerificationID     string    `gorm:"not null;unique" json:"verification_id"` // Business Registration Number
	PreferredPayment   string    `gorm:"not null" json:"preferred_payment"`
	ContactPersonName  string    `gorm:"not null" json:"contact_person_name"`
	PhoneNumber        string    `gorm:"not null;unique" json:"phone_number"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Photos  		   string    `json:"photos"`
}
