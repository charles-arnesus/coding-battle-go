package command

import (
	"bufio"
	"fmt"
	"os"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
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

	// confirm to user that there is possibility flight no run yet
	currentDay := h.systemOperation.GetCurrentDay()
	isSkip, err := h.cancelPreviousScheduledFlight(currentDay, utils.SCHEDULED)
	if err != nil || isSkip {
		return
	}

	fmt.Println("Advancing to the next day...")

	currentDay = h.systemOperation.SetNextDay()
	fmt.Printf("Current day is now: %d\n", currentDay)

	flightRoutes, err := h.flightService.GetFlightRoutes(currentDay, currentDay+1)
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	strDay := ""
	for _, flightRoute := range flightRoutes {
		// if no user use this route today, make it cancelled
		if len(flightRoute.FlightRouteSeat) == 0 {
			flightRoute.Status = utils.CANCELLED
			updatedRecord := flight_model.UpsertFlightRouteRequest{
				FlightRoute: flightRoute,
			}
			_ = h.flightService.UpdateFlightRouteStatus(updatedRecord)

			fmt.Printf("Flight %s -> %s is cancelled, because no passenger\n", flightRoute.DepartureCity.Name, flightRoute.DestinationCity.Name)
			continue
		}

		switch flightRoute.DepartureDay {
		case currentDay:
			strDay = "today"
		case currentDay + 1:
			strDay = "tomorrow"
		default:
			strDay = fmt.Sprintf("Day %d", flightRoute.DepartureDay)
		}

		fmt.Printf("Flight %s -> %s is scheduled for %s.\n", flightRoute.DepartureCity.Name, flightRoute.DestinationCity.Name, strDay)
	}

	if len(flightRoutes) == 0 {
		if strDay == "today" {
			strDay = ""
		} else if strDay != "" {
			strDay = "until " + strDay
		}

		fmt.Printf("There is no flight schedule for today %s", strDay)
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

func (h *GoToNextDayCommand) cancelPreviousScheduledFlight(currentDay int, status string) (isSkip bool, err error) {
	params := flight_model.GetFlightRouteByRequest{
		DepartureDay: currentDay,
		Status:       status,
	}

	flightRoutes, err := h.flightService.GetFlightRouteByParams(params)
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	if len(flightRoutes) > 0 {
		reader := bufio.NewReaderSize(os.Stdin, 1)

		fmt.Printf("There are %d flights have no run yet\n", len(flightRoutes))
		fmt.Printf("Continue this process will be cancelled the flight route status...\n")
		fmt.Print("Confirm (y/n): ")
		userConfirm, errInput := reader.ReadString('\n')
		if errInput != nil {
			err = utils.ErrInputInvalid
			return
		}

		if userConfirm == utils.No {
			isSkip = true
			return
		}
	}

	for _, flightRoute := range flightRoutes {
		flightRoute.Status = utils.CANCELLED
		updatedRecord := flight_model.UpsertFlightRouteRequest{
			FlightRoute: flightRoute,
		}
		_ = h.flightService.UpdateFlightRouteStatus(updatedRecord)

		fmt.Printf("Flight %s -> %s is cancelled by admin\n", flightRoute.DepartureCity.Name, flightRoute.DestinationCity.Name)
	}

	return
}
