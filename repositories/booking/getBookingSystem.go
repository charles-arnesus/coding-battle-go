package booking_repository

import booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"

func (r *bookingRepository) GetBookingSystem() (booking_model.BookingSystem, error) {
	var bookingSystem booking_model.BookingSystem
	err := r.db.First(&bookingSystem).Error
	return bookingSystem, err
}
