package flight_service

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (s *flightService) GetFlightRouteByParams(in flight_model.GetFlightRouteByRequest) (flightRoute []flight_model.FlightRoute, err error) {
	flightRoute, err = s.flightRepository.FindFlightRouteByParams(in)

	return
}
