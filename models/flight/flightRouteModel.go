package flight_model

import (
	"gorm.io/gorm"
)

type FlightRoute struct {
	gorm.Model
	AircraftID        uint
	DepartureCityID   uint
	DestinationCityID uint
	Aircraft          Aircraft    `gorm:"foreignKey:AircraftID;"`
	DepartureCity     Destination `gorm:"foreignKey:DepartureCityID;"`
	DestinationCity   Destination `gorm:"foreignKey:DestinationCityID;"`
	ScheduledDay      int         `gorm:"not null"`
}
