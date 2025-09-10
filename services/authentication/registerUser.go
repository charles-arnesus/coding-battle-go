package authentication_service

import (
	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
)

func (r *authenticationService) RegisterUser(user user_model.User) (err error) {

	err = r.passengerRepository.RegisterUser(user)

	if err == nil {
		loggedUser = user
	}

	return
}
