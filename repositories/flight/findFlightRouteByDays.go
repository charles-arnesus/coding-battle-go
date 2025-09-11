package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) FindFlightRoutesByDay(minDay, maxDay int) (flightRoutes []flight_model.FlightRoute, err error) {
	err = r.db.
		Preload("Aircraft").
		Preload("DepartureCity").
		Preload("DestinationCity").
		Find(&flightRoutes, "scheduled_day >= ? AND scheduled_day <= ?", minDay, maxDay).
		Order("scheduled_day, aircraft_id, departure_city_id").
		Error

	return
}
