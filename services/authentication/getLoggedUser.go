package authentication_service

import (
	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
)

func (r *authenticationService) GetLoggedUser() user_model.User {
	return r.userRepository.GetLoggedUser()
}
