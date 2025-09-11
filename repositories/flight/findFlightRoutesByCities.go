package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) FindFlightRoutesByCities(departurecityID, destinationCityID uint) (flightRoutes []flight_model.FlightRoute, err error) {
	err = r.db.
		Preload("Aircraft").
		Preload("DepartureCity").
		Preload("DestinationCity").
		Find(&flightRoutes, "departure_city_id = ? AND destination_city_id = ?", departurecityID, destinationCityID).
		Order("scheduled_day, aircraft_id, departure_city_id").
		Error

	return
}
