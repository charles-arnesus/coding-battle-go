package flight_service

import flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"

type FlightService interface {
	AddAircraft(aircraft flight_model.Aircraft) (err error)
	AddDestination(destination flight_model.Destination) (err error)
	AddFlightRoute(in flight_model.UpsertFlightRouteRequest) (err error)
	UpdateFlightRouteStatus(in flight_model.UpsertFlightRouteRequest) (err error)

	GetAircrafts(name string) (aircrafts []flight_model.Aircraft, err error)
	GetDestinations(name string) (destinations []flight_model.Destination, err error)
	GetFlightRoutes(minDay, maxDay int) (flightRoutes []flight_model.FlightRoute, err error)
	GetAvailableFlightRoutesByCity(in flight_model.GetAvailableFlightRoutesByCityRequest) (out flight_model.GetAvailableFlightRoutesByCityResponse, err error)
	GetAvailableFlightRoute(in flight_model.GetAvailableFlightRouteRequest) (out flight_model.GetAvailableFlightRouteResponse, err error)
	GetFlightRouteByParams(in flight_model.GetFlightRouteByRequest) (flightRoute []flight_model.FlightRoute, err error)
}
