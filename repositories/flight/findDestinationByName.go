package flight_repository

import (
	"errors"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"gorm.io/gorm"
)

func (r *flightRepository) FindDestinationByName(name string) (destination flight_model.Destination, err error) {

	err = r.db.
		First(&destination, "name = ?", name).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}

	return
}
