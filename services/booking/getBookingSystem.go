package booking_service

import booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"

func (r *bookingService) GetBookingSystem() (booking_model.BookingSystem, error) {
	return r.bookingRepository.GetBookingSystem()
}
