package user_model

type User struct {
	Username string
	Name     string
	Role     string
}

type LoginDto struct {
	Username string
	Role     string
}

type FindByRoleDto struct {
	Role string
}

type FindByUsernameDto struct {
	Username string
}
