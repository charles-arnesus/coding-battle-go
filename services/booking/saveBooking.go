package booking_service

import (
	booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (r *bookingService) SaveBooking(in booking_model.SaveBookingRequest) (out booking_model.SaveBookingResponse, err error) {
	booking, err := r.bookingRepository.SaveBooking()
	if err != nil {
		return
	}

	for _, bookingFlightRoute := range in.FlightRoutes {
		bookingFlightRouteRequest := booking_model.BookingFlightRoute{
			BookingID:     booking.ID,
			FlightRouteID: bookingFlightRoute.ID,
		}
		err = r.bookingRepository.SaveBookingFlightRoute(bookingFlightRouteRequest)
		if err != nil {
			return
		}

		out.FlightRoutes = append(out.FlightRoutes, bookingFlightRoute)

		// Seat Number Logic

		// Get Aircraft maximum seats
		aircraft, errFindAircraft := r.flightRepository.FindAircraftByID(bookingFlightRoute.AircraftID)
		if errFindAircraft != nil {
			err = errFindAircraft
			return
		}

		// Get all taken seat numbers
		takenSeats, errFindFlightRouteSeat := r.flightRepository.FindTakenFlightRouteSeats(bookingFlightRoute.ID)
		if errFindFlightRouteSeat != nil {
			err = errFindFlightRouteSeat
			return
		}

		// Find smallest available seat
		seatMap := make(map[int]bool)
		for _, s := range takenSeats {
			seatMap[s] = true
		}

		seatNumber := 0
		for i := 1; i <= int(aircraft.Seats); i++ {
			if !seatMap[i] {
				seatNumber = i
				break
			}
		}

		if seatNumber == 0 {
			err = utils.ErrNoSeatsAvailableMsg
			return
		}

		flightRouteSeatRequest := flight_model.FlightRouteSeat{
			FlightRouteID: bookingFlightRoute.ID,
			UserID:        in.UserID,
			SeatNumber:    seatNumber,
		}
		err = r.flightRepository.InsertFlightRouteSeat(flightRouteSeatRequest)
		if err != nil {
			return
		}

		out.FligthRouteSeats = append(out.FligthRouteSeats, flightRouteSeatRequest)
	}

	out.BookingID = booking.ID

	return

}
