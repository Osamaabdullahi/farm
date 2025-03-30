package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Farmer Model
type Farmer struct {
	ID               uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID           uuid.UUID   `gorm:"type:uuid;not null" json:"user_id"`
	FarmName         string      `gorm:"not null" json:"farm_name"`
	FarmLocation     string      `gorm:"not null" json:"farm_location"`
	TypesOfCrops     string 	 `gorm:"not null" json:"types_of_crops"`
	HarvestFrequency string      `gorm:"not null" json:"harvest_frequency"`
	VerificationID   string      `gorm:"not null;unique" json:"verification_id"`
	PreferredPayment string 	 `gorm:"not null" json:"preferred_payment"`
	PhoneNumber      string      `gorm:"not null;unique" json:"phone_number"`
	FarmPhotos       StringSlice `gorm:"type:jsonb" json:"farm_photos"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}




// BeforeCreate hook to ensure UUIDs are set properly
func (f *Farmer) BeforeCreate(tx *gorm.DB) (err error) {
	f.ID = uuid.New()
	return
}
