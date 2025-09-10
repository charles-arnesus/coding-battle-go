package authentication_service

import (
	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (r *authenticationService) GetLoggedUser() (user_model.User, error) {
	return user_model.User{
		Name: "Admin",
		Role: utils.RoleAdmin,
	}, nil
}
