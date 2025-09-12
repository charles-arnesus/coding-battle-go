package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) UpdateFlightRouteStatus(flightRoute flight_model.FlightRoute) (err error) {
	err = r.db.
		Model(&flightRoute).
		Updates(flight_model.FlightRoute{Status: flightRoute.Status}).
		Error

	return
}
