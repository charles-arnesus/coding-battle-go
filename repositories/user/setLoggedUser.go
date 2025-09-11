package user_repository

import user_model "github.com/charles-arnesus/coding-battle-go/models/user"

func (r *userRepository) SetLoggedUser(user user_model.User) {
	loggedUser = user
}
