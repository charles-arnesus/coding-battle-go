package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) FindFlightRouteByParams(in flight_model.GetFlightRouteByRequest) (flightRoute []flight_model.FlightRoute, err error) {

	trx := r.db.
		Preload("Aircraft").
		Preload("DepartureCity").
		Preload("DestinationCity").
		Preload("FlightRouteSeat").
		Preload("FlightRouteSeat.User")

	if in.AircraftID > 0 {
		trx = trx.Where("aircraft_id = ?", in.AircraftID)
	}

	if in.DepartureDay > 0 {
		trx = trx.Where("departure_day = ?", in.DepartureDay)
	}

	if in.DepartureTime != "" {
		trx = trx.Where("departure_time = ?", in.DepartureTime)
	}

	if in.DepartureCity > 0 {
		trx = trx.Where("departure_city_id = ?", in.DepartureCity)
	}

	if in.DestinationCity > 0 {
		trx = trx.Where("destination_city_id = ?", in.DestinationCity)
	}

	if in.Status != "" {
		trx = trx.Where("status = ?", in.Status)
	}

	err = trx.Find(&flightRoute).
		Error

	return
}
