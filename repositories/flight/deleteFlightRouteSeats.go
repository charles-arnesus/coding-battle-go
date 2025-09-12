package flight_repository

import flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"

func (r *flightRepository) DeleteFlightRouteSeats(flightRouteSeatIDs []uint) error {
	err := r.db.
		Where("id IN ?", flightRouteSeatIDs).
		Unscoped().
		Delete(&flight_model.FlightRouteSeat{}).Error

	return err
}
