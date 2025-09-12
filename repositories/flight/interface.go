package flight_repository

import flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"

type FlightRepository interface {
	InsertAircraft(aircraft flight_model.Aircraft) error
	InsertDestination(destination flight_model.Destination) error
	InsertFlightRoute(flightRoute flight_model.FlightRoute) (err error)
	InsertFlightRouteSeat(in flight_model.FlightRouteSeat) error

	flightRepositoryRead
}

type flightRepositoryRead interface {
	GetAllAircraft() (aircrafts []flight_model.Aircraft, err error)
	GetAllDestinations() (destinations []flight_model.Destination, err error)
	FindAircraftByName(name string) (aircraft flight_model.Aircraft, err error)
	FindAircraftByID(ID uint) (aircraft flight_model.Aircraft, err error)
	FindDestinationByName(name string) (destination flight_model.Destination, err error)
	FindFlightRouteByID(ID uint) (flightRoute flight_model.FlightRoute, err error)
	FindFlightRoutesByDay(minDay, maxDay int) (flightRoutes []flight_model.FlightRoute, err error)
	FindFlightRoutesByCity(cityID uint, departureDay int, departureTime string) (flightRoutes []flight_model.FlightRoute, err error)
	FindFlightRoutesByCities(departureCityID, destinationCityID uint, departureDay int, departureTime string) (flightRoute []flight_model.FlightRoute, err error)
	FindFlightRouteSeats(flightRouteID uint) (flightRouteSeats []flight_model.FlightRouteSeat, err error)
	FindFlightRouteSeatsUserID(userID uint) (flightRouteSeats []flight_model.FlightRouteSeat, err error)
	FindTakenFlightRouteSeats(flightRouteID uint) (takenSeats []int, err error)
	FindFlightRouteByParams(in flight_model.GetFlightRouteByRequest) (flightRoute []flight_model.FlightRoute, err error)
}
