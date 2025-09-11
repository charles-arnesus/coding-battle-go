package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) GetAllAircraft() (aircrafts []flight_model.Aircraft, err error) {
	err = r.db.Find(&aircrafts).Error

	return
}
