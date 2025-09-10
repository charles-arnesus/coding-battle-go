package flight_service

import flight_repository "github.com/charles-arnesus/coding-battle-go/repositories/flight"

type flightService struct {
	flightRepository flight_repository.FlightRepository
}

func NewFlightService(flightRepository flight_repository.FlightRepository) *flightService {
	return &flightService{
		flightRepository: flightRepository,
	}
}
