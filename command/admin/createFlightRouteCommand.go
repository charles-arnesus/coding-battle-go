package command

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	flight_service "github.com/charles-arnesus/coding-battle-go/services/flight"
	system_operation_service "github.com/charles-arnesus/coding-battle-go/services/systemOperation"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

type CreateFlightRouteCommand struct {
	flightService   flight_service.FlightService
	systemOperation system_operation_service.SystemOperationService
}

func NewCreateFlightRouteCommand(flightService flight_service.FlightService, systemOperation system_operation_service.SystemOperationService) *CreateFlightRouteCommand {
	return &CreateFlightRouteCommand{
		flightService:   flightService,
		systemOperation: systemOperation,
	}
}

func (h *CreateFlightRouteCommand) Execute() (err error) {
	reader := bufio.NewReaderSize(os.Stdin, 1)

	aircrafts, err := h.flightService.GetAircrafts("")
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	if len(aircrafts) == 0 {
		err = utils.ErrAircraftDataEmpty
		return
	}

	destinations, err := h.flightService.GetDestinations("")
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	if len(destinations) == 0 {
		err = utils.ErrDestinationsDataEmpty
		return
	}

	destinationsStrings := make([]string, len(destinations))
	for i, destination := range destinations {
		destinationsStrings[i] = destination.Name
	}

	aircraftsStrings := make([]string, len(aircrafts))
	for i, aircraft := range aircrafts {
		aircraftsStrings[i] = aircraft.Name
	}

	fmt.Println("=== CREATE FLIGHT ROUTE ===")
	fmt.Printf("Available city: %s\n", strings.Join(destinationsStrings, ", "))
	fmt.Printf("Available aircraft: %s\n", strings.Join(aircraftsStrings, ", "))

	fmt.Println()
	fmt.Print("Enter departure city: ")
	departureCity, err := reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}
	departureCity = strings.ToUpper(strings.TrimSpace(departureCity))
	departureCityObj, err := h.flightService.GetDestinations(departureCity)
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}

	if len(departureCityObj) == 0 {
		err = utils.ErrRecordNotFound
		return
	}

	fmt.Print("Enter destination city: ")
	destinationCity, err := reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}
	destinationCity = strings.ToUpper(strings.TrimSpace(destinationCity))
	if destinationCity == departureCity {
		err = utils.ErrDepartureDestinationSame
		return
	}

	destinationCityObj, err := h.flightService.GetDestinations(destinationCity)
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}

	if len(destinationCityObj) == 0 {
		err = utils.ErrRecordNotFound
		return
	}

	fmt.Print("Select aircraft: ")
	aircraftName, err := reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}
	aircraftName = strings.ToUpper(strings.TrimSpace(aircraftName))
	aircraftObj, err := h.flightService.GetAircrafts(aircraftName)
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}

	if len(aircraftObj) == 0 {
		err = utils.ErrRecordNotFound
		return
	}

	currentDay := h.systemOperation.GetCurrentDay()
	fmt.Printf("Enter scheduled Day [%d (current day) - %d]: ", currentDay, utils.MaxDaysInYear)
	scheduledDayStr, err := reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}
	scheduledDay, _ := strconv.Atoi(strings.TrimSpace(scheduledDayStr))

	flightRoute := flight_model.FlightRoute{
		Aircraft:        aircraftObj[0],
		DepartureCity:   departureCityObj[0],
		DestinationCity: destinationCityObj[0],
		ScheduledDay:    scheduledDay,
		Status:          utils.SCHEDULED,
	}

	err = h.flightService.AddFlightRoute(flight_model.AddFlightRouteDTO{
		FlightRoute: flightRoute,
		CurrentDay:  currentDay,
	})

	if err == nil {
		fmt.Printf(utils.CreateFlightRouteSuccessMessage, departureCity, destinationCity, aircraftName, scheduledDay)
	}

	return
}

func (h *CreateFlightRouteCommand) ID() string {
	return utils.AdminCreateFligthRouteID
}

func (h *CreateFlightRouteCommand) AllowedRole() []string {
	return []string{
		utils.RoleAdmin,
	}
}
