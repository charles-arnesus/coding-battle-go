package user_repository

import (
	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
)

func (r *userRepository) FindByRole(in *user_model.FindByRoleDto) (user_model.User, error) {
	var user user_model.User
	err := r.db.First(&user, "role = ?", in.Role).Error
	return user, err
}
