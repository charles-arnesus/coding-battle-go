package flight_service

import flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"

type FlightService interface {
	AddAircraft(aircraft flight_model.Aircraft) (err error)
	AddDestination(destination flight_model.Destination) (err error)
	AddFlightRoute(in flight_model.AddFlightRouteDTO) (err error)

	GetAircrafts(name string) (aircrafts []flight_model.Aircraft, err error)
	GetDestinations(name string) (destinations []flight_model.Destination, err error)
	GetFlightRoutes(day int) (flightRoutes []flight_model.FlightRoute, err error)
}
