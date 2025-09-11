package flight_service

import (
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
)

func (s *flightService) GetAircrafts(name string) (aircrafts []flight_model.Aircraft, err error) {
	if name == "" {
		aircrafts, err = s.flightRepository.GetAllAircraft()
	} else {
		aircraft, errMessage := s.flightRepository.FindAircraftByName(name)
		if errMessage != nil {
			err = errMessage
			return
		}

		if aircraft.ID > 0 {
			aircrafts = append(aircrafts, aircraft)
		}
	}
	return
}
