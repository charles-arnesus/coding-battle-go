package flight_repository

import (
	"errors"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"gorm.io/gorm"
)

func (r *flightRepository) FindAircraftByName(name string) (aircraft flight_model.Aircraft, err error) {
	err = r.db.
		First(&aircraft, "name = ?", name).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}

	return
}
