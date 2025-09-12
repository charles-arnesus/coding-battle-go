package booking_repository

import (
	booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"
)

func (r *bookingRepository) DeleteBookingFlightRoutes(bookingFlightRouteIDs []uint) error {
	err := r.db.
		Where("id IN ?", bookingFlightRouteIDs).
		Unscoped().
		Delete(&booking_model.BookingFlightRoute{}).Error

	return err
}
