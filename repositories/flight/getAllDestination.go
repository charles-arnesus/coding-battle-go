package flight_repository

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (r *flightRepository) GetAllDestinations() (destinations []flight_model.Destination, err error) {
	err = r.db.Find(&destinations).Error

	return
}
