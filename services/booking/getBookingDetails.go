package booking_service

import (
	booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *bookingService) GetBookingDetails(userID uint, currentDay int) (booking_model.GetBookingDetailsResponse, error) {
	response := booking_model.GetBookingDetailsResponse{}
	flightRouteSeats, err := r.flightRepository.FindFlightRouteSeatsUserID(userID)
	if err != nil {
		return response, err
	}
	for _, flightRouteSeat := range flightRouteSeats {
		flightRoute, err := r.flightRepository.FindFlightRouteByID(flightRouteSeat.FlightRouteID)
		if err != nil {
			return response, err
		}
		if flightRoute.DepartureDay < currentDay {
			continue
		}

		bookingFlightRoute, err := r.bookingRepository.FindBookingFlightRoutesByFlightRouteID(flightRoute.ID)
		if err != nil {
			return response, err
		}

		bookingIDfound := false
		var bookingDetailIdx int
		for idx, bookingDetail := range response.BookingDetails {
			if bookingDetail.BookingID == bookingFlightRoute.BookingID {
				bookingIDfound = true
				bookingDetailIdx = idx
			}
		}

		if bookingIDfound {
			response.BookingDetails[bookingDetailIdx].BookingFlightRoutes =
				append(response.BookingDetails[bookingDetailIdx].BookingFlightRoutes, bookingFlightRoute)
			response.BookingDetails[bookingDetailIdx].FlightRoutes =
				append(response.BookingDetails[bookingDetailIdx].FlightRoutes, flightRoute)
			response.BookingDetails[bookingDetailIdx].FlightRouteSeats =
				append(response.BookingDetails[bookingDetailIdx].FlightRouteSeats, flightRouteSeat)
		} else {
			bookingFlightRouteData := []booking_model.BookingFlightRoute{bookingFlightRoute}
			flightRouteData := []flight_model.FlightRoute{flightRoute}
			flightRouteSeatData := []flight_model.FlightRouteSeat{flightRouteSeat}
			bookingDetail := booking_model.BookingDetail{
				BookingID:           bookingFlightRoute.BookingID,
				BookingFlightRoutes: bookingFlightRouteData,
				FlightRoutes:        flightRouteData,
				FlightRouteSeats:    flightRouteSeatData,
			}

			response.BookingDetails = append(response.BookingDetails, bookingDetail)
		}

	}

	return response, err
}
