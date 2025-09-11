package flight_model

import (
	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
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
	Status            string      `gorm:"not null"`
}

type FlightRouteSeat struct {
	gorm.Model
	FlightRouteID uint
	UserID        uint
	SeatNumber    int
	FlightRoute   FlightRoute     `gorm:"foreignKey:FlightRouteID;"`
	User          user_model.User `gorm:"foreignKey:UserID;"`
}

type AddFlightRouteDTO struct {
	FlightRoute FlightRoute
	CurrentDay  int
}

type GetAvailableFlightRouteDto struct {
	DepartureCityID   uint
	DestinationCityID uint
	CurrentDay        int
}
