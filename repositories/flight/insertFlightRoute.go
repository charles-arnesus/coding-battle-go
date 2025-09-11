package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) InsertFlightRoute(flightRoute flight_model.FlightRoute) (err error) {
	err = r.db.Create(&flightRoute).Error

	return
}
