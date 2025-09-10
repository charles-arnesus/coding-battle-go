package authentication_service

import (
	passenger_repository "github.com/charles-arnesus/coding-battle-go/repositories/passenger"
)

type authenticationService struct {
	passengerRepository passenger_repository.PassengerRepository
}

func NewAuthenticationService(passengerRepository passenger_repository.PassengerRepository) *authenticationService {
	return &authenticationService{
		passengerRepository: passengerRepository,
	}
}
