package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) FindFlightRouteSeatsUserID(userID uint) (flightRouteSeats []flight_model.FlightRouteSeat, err error) {
	err = r.db.
		Find(&flightRouteSeats, "user_id = ?", userID).
		Error

	return
}
