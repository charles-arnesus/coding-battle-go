package booking_service

import (
	"fmt"

	booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
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

		//Seat Number Logic

		flightRouteSeatRequest := flight_model.FlightRouteSeat{
			FlightRouteID: bookingFlightRoute.ID,
			UserID:        in.UserID,
			SeatNumber:    1,
		}
		err = r.flightRepository.InsertFlightRouteSeat(flightRouteSeatRequest)
		if err != nil {
			fmt.Println("disini? 3")
			return
		}

		out.FligthRouteSeats = append(out.FligthRouteSeats, flightRouteSeatRequest)
	}

	out.BookingID = booking.ID

	return

}
