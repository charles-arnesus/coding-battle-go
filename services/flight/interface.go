package flight_service

type FlightService interface {
	RegisterAircraft() error
	AddDestination() error
}
