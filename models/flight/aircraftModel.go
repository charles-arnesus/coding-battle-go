package flight_model

import (
	"gorm.io/gorm"
)

type Aircraft struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex; not null"`
	Seats int64  `gorm:"not null"`
}
