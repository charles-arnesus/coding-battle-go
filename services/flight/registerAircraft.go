package flight_service

import "fmt"

func (s *flightService) RegisterAircraft() error {
	fmt.Println("masuk ke service")
	return s.flightRepository.InsertAircraft()
}
