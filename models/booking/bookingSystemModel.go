package booking_model

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"gorm.io/gorm"
)

type BookingSystem struct {
	gorm.Model
	IsActive bool `gorm:"not null"`
}

type Booking struct {
	gorm.Model
	Code string `gorm:"uniqueIndex; not null"`
}

type BookingFlightRoute struct {
	gorm.Model
	BookingID     uint
	FlightRouteID uint
	Booking       Booking                  `gorm:"foreignKey:BookingID;"`
	FlightRoute   flight_model.FlightRoute `gorm:"foreignKey:FlightRouteID;"`
}
