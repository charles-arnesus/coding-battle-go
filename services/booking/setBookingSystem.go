package booking_service

import booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"

func (r *bookingService) SetBookingSystem(in booking_model.BookingSystem) error {
	in.IsActive = !in.IsActive
	return r.bookingRepository.SetBookingSystem(in)
}
