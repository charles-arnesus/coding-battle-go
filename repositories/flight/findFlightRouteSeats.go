package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) FindFlightRouteSeats(flightRouteID uint) (flightRouteSeats []flight_model.FlightRouteSeat, err error) {
	err = r.db.
		Find(&flightRouteSeats, "flight_route_id = ?", flightRouteID).
		Error

	return
}
