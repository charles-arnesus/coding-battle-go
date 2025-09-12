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

	// get available city and aircraft
	destinationsStrings, aircraftsStrings, err := h.availableCityAircraftSection()
	if err != nil {
		return
	}

	fmt.Println("=== CREATE FLIGHT ROUTE ===")
	fmt.Printf("Available city: %s\n", strings.Join(destinationsStrings, ", "))
	fmt.Printf("Available aircraft: %s\n", strings.Join(aircraftsStrings, ", "))

	// get departure and destination city input
	departureCityObj, destinationCityObj, err := h.cityInputSection(reader)
	if err != nil {
		return
	}

	currentDay := h.systemOperation.GetCurrentDay()

	// get the scheduled flight time input
	departureDay, arrivalDay, departureTime, arrivalTime, err := scheduledTimeSection(currentDay, reader)
	if err != nil {
		return
	}

	// get aircraft input
	aircraftObj, err := h.aircraftInputSection(reader)
	if err != nil {
		return
	}

	// validate the flight route with existing
	err = h.checkExistingFlightRoute(departureDay, aircraftObj.ID, departureCityObj.ID, destinationCityObj.ID)
	if err != nil {
		return
	}

	// add flight route
	flightRoute := flight_model.FlightRoute{
		Aircraft:        aircraftObj,
		DepartureCity:   departureCityObj,
		DestinationCity: destinationCityObj,
		DepartureDay:    departureDay,
		DepartureTime:   departureTime,
		ArrivalDay:      arrivalDay,
		ArrivalTime:     arrivalTime,
		Status:          utils.SCHEDULED,
	}

	err = h.flightService.AddFlightRoute(flight_model.UpsertFlightRouteRequest{
		FlightRoute: flightRoute,
		CurrentDay:  currentDay,
	})

	if err == nil {
		fmt.Printf(utils.CreateFlightRouteSuccessMessage, departureCityObj.Name, destinationCityObj.Name, aircraftObj.Name, departureDay, departureTime)
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

func (h *CreateFlightRouteCommand) availableCityAircraftSection() (destinationsStrings, aircraftsStrings []string, err error) {
	// get all aircrafts
	aircrafts, err := h.flightService.GetAircrafts("")
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	if len(aircrafts) == 0 {
		err = utils.ErrAircraftDataEmpty
		return
	}

	// get all destinations
	destinations, err := h.flightService.GetDestinations("")
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	if len(destinations) == 0 {
		err = utils.ErrDestinationsDataEmpty
		return
	}

	// join the destination in string
	destinationsStrings = make([]string, len(destinations))
	for i, destination := range destinations {
		destinationsStrings[i] = destination.Name
	}

	// join the aircraft in string
	aircraftsStrings = make([]string, len(aircrafts))
	for i, aircraft := range aircrafts {
		aircraftsStrings[i] = aircraft.Name
	}

	return
}

func (h *CreateFlightRouteCommand) cityInputSection(reader *bufio.Reader) (departure, destination flight_model.Destination, err error) {
	fmt.Println()
	fmt.Print("Enter departure city: ")
	departureCity, err := reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}
	departureCity = strings.ToUpper(strings.TrimSpace(departureCity))

	// get departure city information
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

	// get destination city information
	destinationCityObj, err := h.flightService.GetDestinations(destinationCity)
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}

	if len(destinationCityObj) == 0 {
		err = utils.ErrRecordNotFound
		return
	}

	departure, destination = departureCityObj[0], destinationCityObj[0]

	return
}

func scheduledTimeSection(currentDay int, reader *bufio.Reader) (departureDay, arrivalDay int, departureTime, arrivalTime string, err error) {
	possibleDay := currentDay + 1
	if possibleDay >= utils.MaxDaysInYear {
		possibleDay = 1
	}
	fmt.Printf("Enter departure day [%d - %d]: ", possibleDay, utils.MaxDaysInYear)
	departureDayStr, err := reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}
	departureDay, _ = strconv.Atoi(strings.TrimSpace(departureDayStr))

	fmt.Printf("Enter departure time [%s, %s]: ", utils.MORNING, utils.EVENING)
	departureTime, err = reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}

	departureTime = strings.ToUpper(strings.TrimSpace(departureTime))
	if !utils.ContainsString([]string{utils.MORNING, utils.EVENING}, departureTime) {
		err = utils.ErrInputInvalid
		return
	}

	if strings.EqualFold(departureTime, utils.EVENING) {
		arrivalDay = departureDay + 1
		arrivalTime = utils.MORNING
		departureTime = utils.EVENING
	}

	if strings.EqualFold(departureTime, utils.MORNING) {
		arrivalDay = departureDay
		arrivalTime = utils.EVENING
		departureTime = utils.MORNING
	}

	return
}

func (h *CreateFlightRouteCommand) aircraftInputSection(reader *bufio.Reader) (aircraftObj flight_model.Aircraft, err error) {
	fmt.Print("Select aircraft: ")
	aircraftName, err := reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}
	aircraftName = strings.ToUpper(strings.TrimSpace(aircraftName))

	// get aircraft information
	aircraftObjs, err := h.flightService.GetAircrafts(aircraftName)
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}

	if len(aircraftObjs) == 0 {
		err = utils.ErrRecordNotFound
		return
	}

	aircraftObj = aircraftObjs[0]

	return
}

func (h *CreateFlightRouteCommand) checkExistingFlightRoute(departureDay int, aircraftID, departureCityID, destinationCityID uint) (err error) {
	// check the flight route with aircraft and scheduled day information
	params := flight_model.GetFlightRouteByRequest{
		AircraftID:   aircraftID,
		DepartureDay: departureDay,
	}
	existingFlightRoutes, err := h.flightService.GetFlightRouteByParams(params)
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	for _, existingFlightRoute := range existingFlightRoutes {
		// if departure city same, it means the flight route already exists
		if existingFlightRoute.DepartureCityID == departureCityID &&
			existingFlightRoute.DepartureCityID == destinationCityID {
			err = utils.ErrFlightRouteAlreadyExistMsg
			return
		}

		if existingFlightRoute.DepartureCityID == departureCityID {
			err = utils.ErrDuplicateFlightRouteMsg
			return
		}

		// if departure different, the new departure city should be same with previous destination and only possible in evening
		if strings.EqualFold(existingFlightRoute.DepartureTime, utils.EVENING) ||
			departureCityID != existingFlightRoute.DestinationCityID {
			err = utils.ErrAircraftInOtherCity
			return
		}
	}

	return
}
