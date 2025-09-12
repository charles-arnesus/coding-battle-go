package booking_service

import booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"

type BookingService interface {
	SetBookingSystem(in booking_model.BookingSystem) error
	GetBookingSystem() (booking_model.BookingSystem, error)
	GetBookingDetails(userID uint, currentDay int) (booking_model.GetBookingDetailsResponse, error)
	SaveBooking(in booking_model.SaveBookingRequest) (out booking_model.SaveBookingResponse, err error)
}
