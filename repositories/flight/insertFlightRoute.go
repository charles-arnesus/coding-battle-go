package flight_repository

import (
	"strings"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (r *flightRepository) InsertFlightRoute(flightRoute flight_model.FlightRoute) (err error) {
	err = r.db.Create(&flightRoute).Error

	if err != nil && strings.Contains(err.Error(), utils.UniqueViolationCodePostgres) {
		err = utils.ErrDuplicateFlightRouteMsg
	}

	return
}
