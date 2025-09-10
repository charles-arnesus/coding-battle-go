package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
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

func (h *AddDestinationCommand) Execute() (err error) {
	reader := bufio.NewReaderSize(os.Stdin, 1)

	fmt.Print("Enter destination name: ")

	name, err := reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}
	name = strings.ToUpper(strings.TrimSpace(name))

	destination := flight_model.Destination{
		Name: name,
	}

	err = h.flightService.AddDestination(destination)
	if err == nil {
		fmt.Printf(utils.AddDestinationSuccessMessage, name)
	}

	return
}

func (h *AddDestinationCommand) ID() string {
	return utils.AdminAddDestinationID
}

func (h *AddDestinationCommand) AllowedRole() []string {
	return []string{
		utils.RoleAdmin,
	}
}
