package command

import (
	"fmt"

	flight_service "github.com/charles-arnesus/coding-battle-go/services/flight"
	system_operation_service "github.com/charles-arnesus/coding-battle-go/services/systemOperation"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

type GoToNextDayCommand struct {
	systemOperation system_operation_service.SystemOperationService
	flightService   flight_service.FlightService
}

func NewGoToNextDayCommand(systemOperation system_operation_service.SystemOperationService, flightService flight_service.FlightService) *GoToNextDayCommand {
	return &GoToNextDayCommand{
		systemOperation: systemOperation,
		flightService:   flightService,
	}
}

func (h *GoToNextDayCommand) Execute() (err error) {

	fmt.Println("=== NEXT DAY ===")
	fmt.Println("Advancing to the next day...")

	currentDay := h.systemOperation.SetNextDay()
	fmt.Printf("Current day is now: %d\n", currentDay)

	flightRoutes, err := h.flightService.GetFlightRoutes(currentDay + 1)
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	for _, flightRoute := range flightRoutes {
		fmt.Printf("Flight %s -> %s is scheduled for tomorrow.\n", flightRoute.DepartureCity.Name, flightRoute.DestinationCity.Name)
	}

	if len(flightRoutes) == 0 {
		fmt.Println("There is no flight schedule for tomorrow")
	}

	return
}

func (h *GoToNextDayCommand) ID() string {
	return utils.AdminGoToNextDaySystemID
}

func (h *GoToNextDayCommand) AllowedRole() []string {
	return []string{
		utils.RoleAdmin,
	}
}
