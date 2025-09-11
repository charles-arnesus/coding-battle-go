package booking_repository

import "gorm.io/gorm"

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *bookingRepository {
	return &bookingRepository{
		db: db,
	}
}
