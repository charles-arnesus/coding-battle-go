package booking_repository

import (
	booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"
)

func (r *bookingRepository) DeleteBookings(bookingIDs []uint) error {
	err := r.db.
		Where("id IN ?", bookingIDs).
		Unscoped().
		Delete(&booking_model.Booking{}).Error

	return err
}
