package authentication_service

import (
	"fmt"

	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
)

func (r *authenticationService) LoginUser(in *user_model.LoginDto) error {
	fmt.Println("masuk ke authentication service")

	if in.Role != "" {
		findByRoleDto := &user_model.FindByRoleDto{
			Role: in.Role,
		}
		user, err := r.passengerRepository.FindByRole(findByRoleDto)
		if err != nil {
			return err
		}
		loggedUser = user
	} else {
		findByUsernameDto := &user_model.FindByUsernameDto{
			Username: in.Username,
		}
		user, err := r.passengerRepository.FindByUsername(findByUsernameDto)
		if err != nil {
			return err
		}
		loggedUser = user
	}

	return nil
}
