package flight_service

import (
	"errors"
	"fmt"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (s *flightService) AddFlightRoute(flightRoute flight_model.FlightRoute) (err error) {

	// should between current day - max in year
	if flightRoute.ScheduledDay < 0 || flightRoute.ScheduledDay > utils.MaxDaysInYear {
		errMessage := fmt.Sprintf(utils.ScheduledDayInvalidMessage, 0, utils.MaxDaysInYear)
		err = errors.New(errMessage)
		return
	}

	err = s.flightRepository.InsertFlightRoute(flightRoute)

	return
}
