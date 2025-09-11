package flight_service

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (s *flightService) AddDestination(destination flight_model.Destination) (err error) {

	if destination.Name == "" {
		err = utils.ErrNameDestinationRequired
		return
	}

	err = s.flightRepository.InsertDestination(destination)

	return
}
