package passenger_repository

import user_model "github.com/charles-arnesus/coding-battle-go/models/user"

type PassengerRepository interface {
	FindByRole(in *user_model.FindByRoleDto) (user_model.User, error)
	FindByUsername(in *user_model.FindByUsernameDto) (user_model.User, error)
}
