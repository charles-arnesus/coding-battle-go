package booking_repository

import booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"

func (r *bookingRepository) SetBookingSystem(in booking_model.BookingSystem) (err error) {
	err = r.db.Save(&in).Error

	return
}
