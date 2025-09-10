package authentication_service

import (
	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
)

var loggedUser = user_model.User{}

func (r *authenticationService) GetLoggedUser() (user_model.User, error) {
	return loggedUser, nil
}
