package booking_service

import booking_repository "github.com/charles-arnesus/coding-battle-go/repositories/booking"

type bookingService struct {
	bookingRepository booking_repository.BookingRepository
}

func NewBookingRepository(bookingRepository booking_repository.BookingRepository) *bookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
	}
}
