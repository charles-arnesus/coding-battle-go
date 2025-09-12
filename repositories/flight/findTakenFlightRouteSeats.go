package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) FindTakenFlightRouteSeats(flightRouteID uint) (takenSeats []int, err error) {
	if err := r.db.
		Model(&flight_model.FlightRouteSeat{}).
		Where("flight_route_id = ?", flightRouteID).
		Pluck("seat_number", &takenSeats).Error; err != nil {
		return takenSeats, err
	}

	return
}
