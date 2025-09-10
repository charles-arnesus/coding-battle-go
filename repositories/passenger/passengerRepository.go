package passenger_repository

import "gorm.io/gorm"

type passengerRepository struct {
	db *gorm.DB
}

func NewPassengerRepository(db *gorm.DB) *passengerRepository {
	return &passengerRepository{
		db: db,
	}
}
