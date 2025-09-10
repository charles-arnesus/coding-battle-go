package flight_repository

import flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"

type FlightRepository interface {
	InsertAircraft(aircraft flight_model.Aircraft) error
	InsertDestination(destination flight_model.Destination) error
}
