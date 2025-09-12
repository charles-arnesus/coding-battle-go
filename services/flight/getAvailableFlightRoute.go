package flight_service

import (
	"slices"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (s *flightService) GetAvailableFlightRoute(in flight_model.GetAvailableFlightRouteRequest) (flight_model.GetAvailableFlightRouteResponse, error) {
	response := flight_model.GetAvailableFlightRouteResponse{}
	flightRoutes, err := s.flightRepository.FindFlightRoutesByCities(in.DepartureCityID, in.DestinationCityID, in.DepartureDay, in.DepartureTime)
	if err != nil {
		return response, err
	}
	for _, flightRoute := range flightRoutes {
		if flightRoute.DepartureDay <= in.CurrentDay {
			continue
		}
		aircraft, err := s.flightRepository.FindAircraftByID(flightRoute.AircraftID)
		if err != nil {
			return response, err
		}

		flightRouteSeats, err := s.flightRepository.FindFlightRouteSeats(flightRoute.ID)
		if err != nil {
			return response, err
		}

		availableSeats := int(aircraft.Seats) - len(flightRouteSeats)
		if availableSeats > 0 && !slices.Contains(utils.UnavailableFlightStatus, flightRoute.Status) {
			response.FlightRoute = flightRoute
			response.AvailableSeats = availableSeats
			break
		}
	}

	return response, err
}
