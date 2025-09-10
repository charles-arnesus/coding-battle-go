package passenger_repository

import (
	"strings"

	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

func (r *passengerRepository) RegisterUser(user user_model.User) (err error) {
	err = r.db.Create(&user).Error

	if err != nil && strings.Contains(err.Error(), utils.UniqueViolationCodePostgres) {
		err = utils.ErrUsernameAlreadyExistMsg
	}

	return
}
