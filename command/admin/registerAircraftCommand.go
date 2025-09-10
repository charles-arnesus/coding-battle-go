package command

import (
	flight_service "github.com/charles-arnesus/coding-battle-go/services/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

type RegisterAircraftCommand struct {
	flightService flight_service.FlightService
}

func NewRegisterAircraftCommand(flightService flight_service.FlightService) *RegisterAircraftCommand {
	return &RegisterAircraftCommand{
		flightService: flightService,
	}
}

func (h *RegisterAircraftCommand) Execute() error {
	err := h.flightService.RegisterAircraft()
	return err
}

func (h *RegisterAircraftCommand) ID() string {
	return utils.AdminRegisterAircraftID
}

func (h *RegisterAircraftCommand) AllowedRole() []string {
	return []string{
		utils.RoleAdmin,
	}
}
