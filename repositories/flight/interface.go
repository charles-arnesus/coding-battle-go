package flight_repository

import flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"

type FlightRepository interface {
	InsertAircraft(aircraft flight_model.Aircraft) error
	InsertDestination(destination flight_model.Destination) error
	InsertFlightRoute(flightRoute flight_model.FlightRoute) (err error)

	GetAllAircraft() (aircrafts []flight_model.Aircraft, err error)
	GetAllDestinations() (destinations []flight_model.Destination, err error)
	FindAircraftByName(name string) (aircraft flight_model.Aircraft, err error)
	FindAircraftByID(ID uint) (aircraft flight_model.Aircraft, err error)
	FindDestinationByName(name string) (destination flight_model.Destination, err error)
	FindFlightRoutesByDay(minDay, maxDay int) (flightRoutes []flight_model.FlightRoute, err error)
	FindFlightRoutesByCity(cityID uint) (flightRoutes []flight_model.FlightRoute, err error)
	FindFlightRoutesByCities(departureCityID, destinationCityID uint) (flightRoute []flight_model.FlightRoute, err error)
	FindFlightRouteSeats(flightRouteID uint) (flightRouteSeats []flight_model.FlightRouteSeat, err error)
}
