package booking_repository

import (
	booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"
)

func (r *bookingRepository) FindBookingFlightRoutesByFlightRouteID(flightRouteID uint) (bookingFlightRoute booking_model.BookingFlightRoute, err error) {
	err = r.db.
		First(&bookingFlightRoute, "flight_route_id = ?", flightRouteID).Error

	return
}
