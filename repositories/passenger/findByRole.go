package passenger_repository

import (
	"fmt"

	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
)

func (r *passengerRepository) FindByRole(in *user_model.FindByRoleDto) (user_model.User, error) {
	fmt.Println("masuk ke passenger repository via find by role")
	fmt.Println(in)
	return user_model.User{}, nil
}
