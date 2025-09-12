package flight_repository

import (
	"strings"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (r *flightRepository) InsertFlightRouteSeat(in flight_model.FlightRouteSeat) (err error) {
	err = r.db.Create(&in).Error

	if err != nil && strings.Contains(err.Error(), utils.UniqueViolationCodePostgres) {
		err = utils.ErrDuplicateFlightRouteBookingMsg
	}

	return
}
