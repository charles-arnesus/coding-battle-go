package command

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
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

func (h *RegisterAircraftCommand) Execute() (err error) {
	reader := bufio.NewReaderSize(os.Stdin, 1)

	fmt.Print("Enter aircraft name:")

	name, err := reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}
	name = strings.ToUpper(strings.TrimSpace(name))

	fmt.Print("Enter seat capacity: ")

	strSeat, err := reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}
	seat, _ := strconv.Atoi(strings.TrimSpace(strSeat))

	aircraft := flight_model.Aircraft{
		Name:  name,
		Seats: int64(seat),
	}

	err = h.flightService.RegisterAircraft(aircraft)
	if err == nil {
		fmt.Printf(utils.RegisterAircraftSuccessMessage, name, seat)
	}

	return
}

func (h *RegisterAircraftCommand) ID() string {
	return utils.AdminRegisterAircraftID
}

func (h *RegisterAircraftCommand) AllowedRole() []string {
	return []string{
		utils.RoleAdmin,
	}
}
