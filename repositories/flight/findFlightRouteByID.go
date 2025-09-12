package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) FindFlightRouteByID(ID uint) (flightRoute flight_model.FlightRoute, err error) {
	err = r.db.
		Preload("Aircraft").
		Preload("DepartureCity").
		Preload("DestinationCity").
		First(&flightRoute, "id = ?", ID).Error

	return
}
