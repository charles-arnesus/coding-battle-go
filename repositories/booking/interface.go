package booking_repository

import booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"

type BookingRepository interface {
	SetBookingSystem(bookingSystem booking_model.BookingSystem) error
	GetBookingSystem() (booking_model.BookingSystem, error)
	SaveBooking() (booking booking_model.Booking, err error)
	SaveBookingFlightRoute(in booking_model.BookingFlightRoute) (err error)
	FindBookingFlightRoutesByBookingID(bookingID uint) (bookingFlightRoutes []booking_model.BookingFlightRoute, err error)
	FindBookingFlightRoutesByFlightRouteID(FlightRouteID uint) (bookingFlightRoute booking_model.BookingFlightRoute, err error)
}
