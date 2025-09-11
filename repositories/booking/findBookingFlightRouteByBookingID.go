package booking_repository

import (
	booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"
)

func (r *bookingRepository) FindBookingFlightRoutesByBookingID(bookingID uint) (bookingFlightRoutes []booking_model.BookingFlightRoute, err error) {
	err = r.db.
		Preload("FlightRoute").
		Find(&bookingFlightRoutes, "booking_id = ?", bookingID).
		Error

	return
}
