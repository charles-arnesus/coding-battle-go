package command

import (
	flight_service "github.com/charles-arnesus/coding-battle-go/services/flight"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

type AddDestinationCommand struct {
	flightService flight_service.FlightService
}

func NewAddDestinationCommand(flightService flight_service.FlightService) *AddDestinationCommand {
	return &AddDestinationCommand{
		flightService: flightService,
	}
}

func (h *AddDestinationCommand) Execute() error {
	err := h.flightService.AddDestination()
	return err
}

func (h *AddDestinationCommand) ID() string {
	return utils.AdminAddDestinationID
}

func (h *AddDestinationCommand) AllowedRole() []string {
	return []string{
		utils.RoleAdmin,
	}
}
