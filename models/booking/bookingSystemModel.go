package booking_model

import "gorm.io/gorm"

type BookingSystem struct {
	gorm.Model
	IsActive bool `gorm:"not null"`
}
