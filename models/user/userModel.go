package user_model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Name     string `gorm:"not null"`
	Role     string `gorm:"not null"`
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
