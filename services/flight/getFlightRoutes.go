package flight_service

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (s *flightService) GetFlightRoutes(day int) (flightRoutes []flight_model.FlightRoute, err error) {
	flightRoutes, err = s.flightRepository.FindFlightRoutesByDay(day)
	return
}
