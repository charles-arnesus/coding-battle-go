package booking_repository

import booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"

func (r *bookingRepository) SaveBooking() (booking booking_model.Booking, err error) {
	booking = booking_model.Booking{}
	err = r.db.Save(&booking).Error

	return
}
