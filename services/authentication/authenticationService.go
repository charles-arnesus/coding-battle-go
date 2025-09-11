package authentication_service

import user_repository "github.com/charles-arnesus/coding-battle-go/repositories/user"

type authenticationService struct {
	userRepository user_repository.UserRepository
}

func NewAuthenticationService(userRepository user_repository.UserRepository) *authenticationService {
	return &authenticationService{
		userRepository: userRepository,
	}
}
