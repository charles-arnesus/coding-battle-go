package authentication_service

import user_model "github.com/charles-arnesus/coding-battle-go/models/user"

type AuthenticationService interface {
	GetLoggedUser() user_model.User
	LoginUser(in *user_model.LoginDto) error
	RegisterUser(in *user_model.User) error
}
