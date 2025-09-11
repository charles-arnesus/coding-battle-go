package flight_service

import (
	"slices"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (s *flightService) GetAvailableFlightRoute(in flight_model.GetAvailableFlightRouteDto) (flight_model.FlightRoute, error) {
	finalFlightRoute := flight_model.FlightRoute{}
	flightRoutes, err := s.flightRepository.FindFlightRoutesByCities(in.DepartureCityID, in.DestinationCityID)
	for _, flightRoute := range flightRoutes {
		if flightRoute.ScheduledDay <= in.CurrentDay {
			continue
		}
		aircraft, err := s.flightRepository.FindAircraftByID(flightRoute.AircraftID)
		if err != nil {
			return finalFlightRoute, err
		}

		flightRouteSeats, err := s.flightRepository.FindFlightRouteSeats(flightRoute.ID)
		if err != nil {
			return finalFlightRoute, err
		}

		if int(aircraft.Seats)-len(flightRouteSeats) > 0 && !slices.Contains(utils.UnavailableFlightStatus, flightRoute.Status) {
			finalFlightRoute = flightRoute
			break
		}
	}

	return finalFlightRoute, err
}
