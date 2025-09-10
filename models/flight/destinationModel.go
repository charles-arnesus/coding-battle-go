package flight_model

import "gorm.io/gorm"

type Destination struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null"`
}
