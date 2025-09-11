package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) FindFlightRoutesByDay(day int) (flightRoutes []flight_model.FlightRoute, err error) {
	err = r.db.
		Preload("Aircraft").
		Preload("DepartureCity").
		Preload("DestinationCity").
		Find(&flightRoutes, "scheduled_day = ?", day).Error

	return
}
