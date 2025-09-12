package flight_service

import (
	"errors"
	"fmt"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (s *flightService) AddFlightRoute(in flight_model.AddFlightRouteDTO) (err error) {

	// should between current day - max in year
	isPreviousDay := in.FlightRoute.DepartureDay < in.CurrentDay
	if in.CurrentDay == utils.MaxDaysInYear {
		isPreviousDay = false
	}

	if isPreviousDay || in.FlightRoute.DepartureDay > utils.MaxDaysInYear {
		errMessage := fmt.Sprintf(utils.ScheduledDayInvalidMessage, in.CurrentDay, utils.MaxDaysInYear)
		err = errors.New(errMessage)
		return
	}

	err = s.flightRepository.InsertFlightRoute(in.FlightRoute)

	return
}
