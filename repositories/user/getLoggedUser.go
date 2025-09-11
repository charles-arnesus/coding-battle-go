package user_repository

import user_model "github.com/charles-arnesus/coding-battle-go/models/user"

var loggedUser = user_model.User{}

func (r *userRepository) GetLoggedUser() user_model.User {
	return loggedUser
}
