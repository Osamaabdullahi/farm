package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	FirstName  string    `gorm:"not null" json:"first_name"`
	LastName   string    `gorm:"not null" json:"last_name"`
	Email      string    `gorm:"unique;not null" json:"email"`
	Password   string    `gorm:"not null" json:"-"` // Hide password in JSON responses
	Onboarding bool      `gorm:"default:false" json:"onboarding"`
	Role       string    `gorm:"type:user_role;default:'farmer'" json:"role"`   // Correct ENUM reference
	Status     string    `gorm:"type:user_status;default:'pending'" json:"status"` // Correct ENUM reference
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Farmer *Farmer `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"farmer,omitempty"`
	Business   *Business  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"business,omitempty"` // Foreign Key
	Logistics *Logistics `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"logistics,omitempty"`

}

// BeforeCreate hook to generate UUID before saving
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}



