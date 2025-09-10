package command

type Command interface {
	ID() string
	AllowedRole() []string
	Execute() error
}
