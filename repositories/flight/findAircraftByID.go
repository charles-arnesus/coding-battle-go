package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) FindAircraftByID(ID uint) (aircraft flight_model.Aircraft, err error) {
	err = r.db.
		First(&aircraft, "id = ?", ID).Error

	return
}
