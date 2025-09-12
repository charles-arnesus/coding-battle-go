package flight_service

import (
	"slices"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (s *flightService) GetAvailableFlightRoutesByCity(in flight_model.GetAvailableFlightRoutesByCityRequest) (flight_model.GetAvailableFlightRoutesByCityResponse, error) {
	response := flight_model.GetAvailableFlightRoutesByCityResponse{}
	flightRoutes, err := s.flightRepository.FindFlightRoutesByCity(in.DepartureCityID, in.DepartureDay, in.DepartureTime)
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
			availableRoute := flight_model.GetAvailableFlightRouteResponse{
				FlightRoute:    flightRoute,
				AvailableSeats: availableSeats,
			}
			response.GetAvailableFlightRouteResponses = append(
				response.GetAvailableFlightRouteResponses,
				availableRoute,
			)
		}
	}
	return response, err
}
