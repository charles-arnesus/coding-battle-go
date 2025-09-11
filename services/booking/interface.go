package booking_service

import booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"

type BookingService interface {
	SetBookingSystem(in booking_model.BookingSystem) error
	GetBookingSystem() (booking_model.BookingSystem, error)
}
