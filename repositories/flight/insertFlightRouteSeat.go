package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) InsertFlightRouteSeat(in flight_model.FlightRouteSeat) (err error) {
	err = r.db.Create(&in).Error

	return
}
