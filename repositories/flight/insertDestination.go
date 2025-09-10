package flight_repository

import (
	"strings"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (r *flightRepository) InsertDestination(destination flight_model.Destination) (err error) {
	err = r.db.Create(&destination).Error

	if err != nil && strings.Contains(err.Error(), utils.UniqueViolationCodePostgres) {
		err = utils.ErrNameDestinationAlreadyExist
	}

	return
}
