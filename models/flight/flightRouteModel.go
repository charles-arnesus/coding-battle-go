package flight_model

import (
	user_model "github.com/charles-arnesus/coding-battle-go/models/user"
	"gorm.io/gorm"
)

type FlightRoute struct {
	gorm.Model
	AircraftID        uint `gorm:"uniqueIndex:idx_departure_arrival_aircraft"`
	DepartureCityID   uint
	DestinationCityID uint
	Aircraft          Aircraft    `gorm:"foreignKey:AircraftID;"`
	DepartureCity     Destination `gorm:"foreignKey:DepartureCityID;"`
	DestinationCity   Destination `gorm:"foreignKey:DestinationCityID;"`
	DepartureTime     string      `gorm:"not null;uniqueIndex:idx_departure_arrival_aircraft"`
	ArrivalTime       string      `gorm:"not null"`
	DepartureDay      int         `gorm:"not null;uniqueIndex:idx_departure_arrival_aircraft"`
	ArrivalDay        int         `gorm:"not null"`
	Status            string      `gorm:"not null"`
}

type FlightRouteSeat struct {
	gorm.Model
	FlightRouteID uint `gorm:"uniqueIndex:idx_flightroute_user"`
	UserID        uint `gorm:"uniqueIndex:idx_flightroute_user"`
	SeatNumber    int
	FlightRoute   FlightRoute     `gorm:"foreignKey:FlightRouteID;"`
	User          user_model.User `gorm:"foreignKey:UserID;"`
}

type AddFlightRouteDTO struct {
	FlightRoute FlightRoute
	CurrentDay  int
}

type GetAvailableFlightRouteRequest struct {
	DepartureCityID   uint
	DestinationCityID uint
	CurrentDay        int
	DepartureDay      int
	DepartureTime     string
}

type GetAvailableFlightRouteResponse struct {
	FlightRoute    FlightRoute
	AvailableSeats int
}

type GetAvailableFlightRoutesByCityRequest struct {
	DepartureCityID uint
	CurrentDay      int
	DepartureDay    int
	DepartureTime   string
}

type GetAvailableFlightRoutesByCityResponse struct {
	GetAvailableFlightRouteResponses []GetAvailableFlightRouteResponse
}
