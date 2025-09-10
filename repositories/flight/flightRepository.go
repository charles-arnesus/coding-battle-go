package flight_repository

import "gorm.io/gorm"

type flightRepository struct {
	db *gorm.DB
}

func NewFlightRepository(db *gorm.DB) *flightRepository {
	return &flightRepository{
		db: db,
	}
}
