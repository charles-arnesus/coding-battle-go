package flight_service

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (s *flightService) GetDestinations(name string) (destinations []flight_model.Destination, err error) {
	if name == "" {
		destinations, err = s.flightRepository.GetAllDestinations()
	} else {
		destination, errMessage := s.flightRepository.FindDestinationByName(name)
		if errMessage != nil {
			err = errMessage
			return
		}

		if destination.ID > 0 {
			destinations = append(destinations, destination)
		}
	}
	return
}
