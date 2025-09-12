package user_repository

import (
	"strings"

	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (r *userRepository) RegisterUser(user user_model.User) (userID uint, err error) {
	result := r.db.Create(&user)
	if result.Error != nil && strings.Contains(result.Error.Error(), utils.UniqueViolationCodePostgres) {
		err = utils.ErrUsernameAlreadyExistMsg
	}

	userID = user.ID

	return
}
