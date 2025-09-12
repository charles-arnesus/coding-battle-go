package booking_service

import (
	booking_repository "github.com/charles-arnesus/coding-battle-go/repositories/booking"
	flight_repository "github.com/charles-arnesus/coding-battle-go/repositories/flight"
)

type bookingService struct {
	bookingRepository booking_repository.BookingRepository
	flightRepository  flight_repository.FlightRepository
}

func NewBookingRepository(bookingRepository booking_repository.BookingRepository, flightRepository flight_repository.FlightRepository) *bookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
		flightRepository:  flightRepository,
	}
}
