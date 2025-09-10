package authentication_service

import user_model "github.com/charles-arnesus/coding-battle-go/models/user"

type AuthenticationService interface {
	GetLoggedUser() (user_model.User, error)
	LoginUser(in *user_model.LoginDto) error
}
