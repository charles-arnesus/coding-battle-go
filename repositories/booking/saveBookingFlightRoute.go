package booking_repository

import booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"

func (r *bookingRepository) SaveBookingFlightRoute(in booking_model.BookingFlightRoute) (err error) {
	err = r.db.Save(&in).Error

	return
}
