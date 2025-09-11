package authentication_service

import (
	"errors"

	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
	"gorm.io/gorm"
)

func (r *authenticationService) LoginUser(in *user_model.LoginDto) error {
	// If logged in as admin
	if in.Role != "" {
		//Find user by role = admin
		findByRoleDto := &user_model.FindByRoleDto{
			Role: in.Role,
		}
		user, err := r.userRepository.FindByRole(findByRoleDto)
		if err != nil {
			return err
		}
		// save user into loggedUser variable
		r.userRepository.SetLoggedUser(user)

		// If logged in as passenger
	} else {
		findByUsernameDto := &user_model.FindByUsernameDto{
			Username: in.Username,
		}
		user, err := r.userRepository.FindByUsername(findByUsernameDto)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		r.userRepository.SetLoggedUser(user)
	}

	return nil
}
