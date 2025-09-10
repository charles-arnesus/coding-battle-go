package flight_service

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (s *flightService) RegisterAircraft(aircraft flight_model.Aircraft) (err error) {
	if aircraft.Seats < 1 {
		return utils.ErrSeatsAircraftCapacity
	}

	if aircraft.Name == "" {
		return utils.ErrNameDestinationRequired
	}

	err = s.flightRepository.InsertAircraft(aircraft)

	return
}
