package flight_service

import flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"

type FlightService interface {
	RegisterAircraft(aircraft flight_model.Aircraft) (err error)
	AddDestination(destination flight_model.Destination) (err error)
}
