package booking_repository

import booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"

type BookingRepository interface {
	SetBookingSystem(bookingSystem booking_model.BookingSystem) error
	GetBookingSystem() (booking_model.BookingSystem, error)
}
