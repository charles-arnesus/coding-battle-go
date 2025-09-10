package passenger_repository

import (
	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
)

func (r *passengerRepository) FindByUsername(in *user_model.FindByUsernameDto) (user_model.User, error) {
	var user user_model.User
	err := r.db.First(&user, "username = ?", in.Username).Error
	return user, err
}
