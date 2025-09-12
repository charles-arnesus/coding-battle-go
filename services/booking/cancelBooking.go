package booking_service

import "strconv"

func (r *bookingService) CancelBooking(bookingIDStr string, userID uint) error {
	flightRouteSeatIDs := []uint{}
	bookingFlightRouteIDs := []uint{}
	bookingID, err := strconv.ParseUint(bookingIDStr, 10, 32)
	if err != nil {
		return err
	}

	bookingFlightRoutes, err := r.bookingRepository.FindBookingFlightRoutesByBookingID(uint(bookingID))
	if err != nil {
		return err
	}
	for _, bookingFlightRoute := range bookingFlightRoutes {
		flightRouteSeats, err := r.flightRepository.FindFlightRouteSeatsUserIDFlightRouteID(userID, bookingFlightRoute.FlightRoute.ID)
		if err != nil {
			return err
		}
		for _, flightRouteSeat := range flightRouteSeats {
			flightRouteSeatIDs = append(flightRouteSeatIDs, flightRouteSeat.ID)
		}
		bookingFlightRouteIDs = append(bookingFlightRouteIDs, bookingFlightRoute.ID)
	}

	err = r.flightRepository.DeleteFlightRouteSeats(flightRouteSeatIDs)
	if err != nil {
		return err
	}
	err = r.bookingRepository.DeleteBookingFlightRoutes(bookingFlightRouteIDs)
	if err != nil {
		return err
	}
	err = r.bookingRepository.DeleteBookings([]uint{uint(bookingID)})
	if err != nil {
		return err
	}

	return err

}
