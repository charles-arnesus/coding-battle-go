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
}

type BookingFlightRoute struct {
	gorm.Model
	BookingID     uint
	FlightRouteID uint
	Booking       Booking                  `gorm:"foreignKey:BookingID;"`
	FlightRoute   flight_model.FlightRoute `gorm:"foreignKey:FlightRouteID;"`
}

type SaveBookingRequest struct {
	FlightRoutes []flight_model.FlightRoute
	UserID       uint
}

type SaveBookingResponse struct {
	BookingID        uint
	FlightRoutes     []flight_model.FlightRoute
	FligthRouteSeats []flight_model.FlightRouteSeat
}

type GetBookingDetailsResponse struct {
	BookingDetails []BookingDetail
}

type BookingDetail struct {
	BookingID           uint
	BookingFlightRoutes []BookingFlightRoute
	FlightRoutes        []flight_model.FlightRoute
	FlightRouteSeats    []flight_model.FlightRouteSeat
}
