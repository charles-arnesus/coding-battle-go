package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	booking_service "github.com/charles-arnesus/coding-battle-go/services/booking"
	flight_service "github.com/charles-arnesus/coding-battle-go/services/flight"
	system_operation_service "github.com/charles-arnesus/coding-battle-go/services/systemOperation"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

type BookFlightCommand struct {
	bookingService  booking_service.BookingService
	flightService   flight_service.FlightService
	systemOperation system_operation_service.SystemOperationService
}

func NewBookFlightCommand(bookingService booking_service.BookingService, flightService flight_service.FlightService, systemOperation system_operation_service.SystemOperationService) *BookFlightCommand {
	return &BookFlightCommand{
		bookingService:  bookingService,
		flightService:   flightService,
		systemOperation: systemOperation,
	}
}

func (h *BookFlightCommand) Execute() (err error) {
	currentDay := h.systemOperation.GetCurrentDay()
	// check if booking service system is enabled
	bookingSystem, err := h.bookingService.GetBookingSystem()
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	if !bookingSystem.IsActive {
		err = utils.ErrBookingServiceDisabledMsg
		return
	}

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

	fmt.Printf("Available city: %s\n", strings.Join(destinationsStrings, ", "))

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

	getAvailableFlightRouteDto := flight_model.GetAvailableFlightRouteDto{
		DepartureCityID:   departureCityObj[0].ID,
		DestinationCityID: destinationCityObj[0].ID,
		CurrentDay:        currentDay,
	}

	directRoute, err := h.flightService.GetAvailableFlightRoute(getAvailableFlightRouteDto)
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	if directRoute.ID == 0 {
		// flightRoutes, errGetFlightRoutesByCity := h.flightService.GetFlightRoutesByCity(departureCityObj[0].ID)
		// if errGetFlightRoutesByCity != nil {
		// 	err = utils.ErrSomethingWentWrongGet
		// 	return err
		// }
		// h.flightService.GetAvailableFlightRoute(departureCityObj[0].ID, destinationCityObj[0].ID)

		fmt.Println("no direct routes")
	}

	fmt.Println(directRoute)

	return
}

func (h *BookFlightCommand) ID() string {
	return utils.PassengerBookFlightID
}

func (h *BookFlightCommand) AllowedRole() []string {
	return []string{
		utils.RolePassenger,
	}
}
