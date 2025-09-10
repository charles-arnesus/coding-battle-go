package passenger_repository

import (
	"fmt"

	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
)

func (r *passengerRepository) FindByUsername(in *user_model.FindByUsernameDto) (user_model.User, error) {
	fmt.Println("masuk ke passenger repository via find by username")
	fmt.Println(in)
	return user_model.User{}, nil
}
