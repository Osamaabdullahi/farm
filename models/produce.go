// package models

// import (
// 	"time"

// 	"github.com/google/uuid"
// )

// // Produce Model
// type Produce struct {
// 	ID                uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
// 	Title             string      `gorm:"not null" json:"title"`
// 	Description       string      `json:"description"`
// 	Category          string      `gorm:"not null" json:"category"`
// 	Images           StringSlice `gorm:"type:jsonb" json:"images"`
// 	PricePerUnit      float64     `gorm:"not null" json:"price_per_unit"`
// 	UnitType          string      `gorm:"not null" json:"unit_type"`
// 	MinimumOrder      int         `gorm:"default:1" json:"minimum_order"`
// 	AvailableStock    int         `gorm:"default:0" json:"available_stock"`
// 	HarvestDate       time.Time   `json:"harvest_date"`
// 	FarmerID          uuid.UUID   `gorm:"not null" json:"farmer_id"`
// 	Farmer            *Farmer     `gorm:"foreignKey:FarmerID;constraint:OnDelete:CASCADE" json:"farmer"`
// 	DeliveryOptions   string      `gorm:"not null" json:"delivery_options"`
// 	EstimatedDelivery string      `json:"estimated_delivery"`
// 	CreatedAt         time.Time   `json:"created_at"`
// 	UpdatedAt         time.Time   `json:"updated_at"`
// }

package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Custom type to handle array of image objects in JSON
type Image struct {
	Path string `json:"path"`
	URL  string `json:"url"`
	Name string `json:"name"`
}

type ImageSlice []Image

// Convert ImageSlice to JSON before storing in DB
func (s ImageSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Convert JSON from DB back into Go struct
func (s *ImageSlice) Scan(value interface{}) error {
	if value == nil {
		*s = []Image{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type for ImageSlice")
	}

	return json.Unmarshal(bytes, s)
}

// Produce struct
type Produce struct {
	ID                uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title             string     `gorm:"not null" json:"title"`
	Description       string     `json:"description"`
	Category         string     `gorm:"not null" json:"category"`
	Images           ImageSlice `gorm:"type:jsonb" json:"images"` // Store images as JSON array

	PricePerUnit      float64    `gorm:"not null" json:"price_per_unit"`
	UnitType          string     `gorm:"not null" json:"unit_type"`
	MinimumOrder      int        `gorm:"default:1" json:"minimum_order"`
	AvailableStock    int        `gorm:"default:0" json:"available_stock"`
	HarvestDate       time.Time  `json:"harvest_date"`

	FarmerID          uuid.UUID  `gorm:"not null" json:"farmer_id"`
	Farmer            *Farmer    `gorm:"foreignKey:FarmerID;constraint:OnDelete:CASCADE" json:"farmer"`

	DeliveryOptions   string     `gorm:"not null" json:"delivery_options"`
	EstimatedDelivery string     `json:"estimated_delivery"`

	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
