package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) FindFlightRouteSeatsUserIDFlightRouteID(userID uint, flightRouteID uint) (flightRouteSeats []flight_model.FlightRouteSeat, err error) {
	err = r.db.
		Find(&flightRouteSeats, "user_id = ? AND flight_route_id = ?", userID, flightRouteID).
		Error

	return
}
