package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) FindFlightRoutesByCity(cityID uint, departureDay int, departureTime string) (flightRoutes []flight_model.FlightRoute, err error) {
	err = r.db.
		Preload("Aircraft").
		Preload("DepartureCity").
		Preload("DestinationCity").
		Find(&flightRoutes, "departure_city_id = ? AND departure_day = ? AND departure_time = ?", cityID, departureDay, departureTime).
		Order("scheduled_day, aircraft_id, departure_city_id").
		Error

	return
}
