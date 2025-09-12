package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	booking_model "github.com/charles-arnesus/coding-battle-go/models/booking"
	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	authentication_service "github.com/charles-arnesus/coding-battle-go/services/authentication"
	booking_service "github.com/charles-arnesus/coding-battle-go/services/booking"
	flight_service "github.com/charles-arnesus/coding-battle-go/services/flight"
	system_operation_service "github.com/charles-arnesus/coding-battle-go/services/systemOperation"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

type BookFlightCommand struct {
	authenticationService authentication_service.AuthenticationService
	bookingService        booking_service.BookingService
	flightService         flight_service.FlightService
	systemOperation       system_operation_service.SystemOperationService
}

func NewBookFlightCommand(authenticationService authentication_service.AuthenticationService, bookingService booking_service.BookingService, flightService flight_service.FlightService, systemOperation system_operation_service.SystemOperationService) *BookFlightCommand {
	return &BookFlightCommand{
		authenticationService: authenticationService,
		bookingService:        bookingService,
		flightService:         flightService,
		systemOperation:       systemOperation,
	}
}

func (h *BookFlightCommand) Execute() (err error) {
	currentDay := h.systemOperation.GetCurrentDay()
	transitRoutes := []flight_model.GetAvailableFlightRouteResponse{}
	bookedRoutes := []flight_model.FlightRoute{}
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

	fmt.Printf("Enter departure Day [%d (current day) - %d]: ", currentDay, utils.MaxDaysInYear)
	departureDayStr, err := reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}
	departureDay, _ := strconv.Atoi(strings.TrimSpace(departureDayStr))

	fmt.Printf("Enter departure time [%s/%s]: ", utils.MORNING, utils.EVENING)
	departureTimeStr, err := reader.ReadString('\n')
	if err != nil {
		err = utils.ErrInputInvalid
		return
	}
	departureTime := strings.ToUpper(strings.TrimSpace(departureTimeStr))

	if !utils.ContainsString([]string{utils.MORNING, utils.EVENING}, departureTime) {
		fmt.Println(utils.ErrInputInvalid)
		return
	}

	if departureDay <= currentDay {
		errMessage := fmt.Sprintf(utils.ScheduledDayInvalidMessage, currentDay+1, utils.MaxDaysInYear)
		err = errors.New(errMessage)
		return
	}

	getAvailableFlightRouteRequest := flight_model.GetAvailableFlightRouteRequest{
		DepartureCityID:   departureCityObj[0].ID,
		DestinationCityID: destinationCityObj[0].ID,
		CurrentDay:        currentDay,
		DepartureDay:      departureDay,
		DepartureTime:     departureTime,
	}

	getAvailableFlightRouteResponse, err := h.flightService.GetAvailableFlightRoute(getAvailableFlightRouteRequest)
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	if getAvailableFlightRouteResponse.FlightRoute.ID == 0 {
		fmt.Println(utils.NoDirectFlightFoundMsg)

		getAvailableFlightRoutesByCityRequest := flight_model.GetAvailableFlightRoutesByCityRequest{
			DepartureCityID: departureCityObj[0].ID,
			CurrentDay:      currentDay,
			DepartureDay:    departureDay,
			DepartureTime:   departureTime,
		}
		availableFlightRoutesByCity, errGetAvailableFlightRoutesByCity := h.flightService.GetAvailableFlightRoutesByCity(getAvailableFlightRoutesByCityRequest)
		if errGetAvailableFlightRoutesByCity != nil {
			err = utils.ErrSomethingWentWrongGet
			return err
		}
		for _, flightRoute := range availableFlightRoutesByCity.GetAvailableFlightRouteResponses {
			var departureTimeTransit = departureTime
			var departureDayTransit = departureDay
			if flightRoute.FlightRoute.DepartureTime == utils.MORNING {
				departureTimeTransit = utils.EVENING
			} else {
				departureTimeTransit = utils.MORNING
				departureDayTransit += 1
			}
			req := flight_model.GetAvailableFlightRouteRequest{
				DepartureCityID:   flightRoute.FlightRoute.DestinationCityID,
				DestinationCityID: destinationCityObj[0].ID,
				CurrentDay:        flightRoute.FlightRoute.DepartureDay - 1,
				DepartureDay:      departureDayTransit,
				DepartureTime:     departureTimeTransit,
			}
			resp, err := h.flightService.GetAvailableFlightRoute(req)
			if err != nil {
				return err
			}

			if resp.FlightRoute.ID == 0 {
				continue
			}

			transitRoutes = append(transitRoutes, flightRoute, resp)
			bookedRoutes = append(bookedRoutes, flightRoute.FlightRoute, resp.FlightRoute)
		}

		if len(transitRoutes) > 0 {
			fmt.Printf(utils.TransitFlightFoundMessage,
				transitRoutes[0].FlightRoute.DepartureCity.Name,
				transitRoutes[0].FlightRoute.DestinationCity.Name,
				transitRoutes[1].FlightRoute.DestinationCity.Name,
				transitRoutes[0].FlightRoute.DepartureDay)
			for _, transiteRoute := range transitRoutes {
				fmt.Printf(utils.AvailableSeatsMessage,
					transiteRoute.AvailableSeats,
					transiteRoute.FlightRoute.DepartureCity.Name,
					transiteRoute.FlightRoute.DestinationCity.Name,
					transiteRoute.FlightRoute.DepartureDay,
					transiteRoute.FlightRoute.DepartureTime,
					transiteRoute.FlightRoute.Aircraft.Name)
			}
		}
	} else {
		fmt.Printf(utils.DirectFlightFoundMessage,
			getAvailableFlightRouteResponse.FlightRoute.DepartureCity.Name,
			getAvailableFlightRouteResponse.FlightRoute.DestinationCity.Name,
			getAvailableFlightRouteResponse.FlightRoute.DepartureDay)
		fmt.Printf(utils.AvailableSeatsMessage,
			getAvailableFlightRouteResponse.AvailableSeats,
			getAvailableFlightRouteResponse.FlightRoute.DepartureCity.Name,
			getAvailableFlightRouteResponse.FlightRoute.DestinationCity.Name,
			getAvailableFlightRouteResponse.FlightRoute.DepartureDay,
			getAvailableFlightRouteResponse.FlightRoute.DepartureTime,
			getAvailableFlightRouteResponse.FlightRoute.Aircraft.Name)
		bookedRoutes = append(bookedRoutes, getAvailableFlightRouteResponse.FlightRoute)
	}

	if getAvailableFlightRouteResponse.FlightRoute.ID == 0 && len(transitRoutes) == 0 {
		fmt.Println(utils.NoFlightFoundMsg)
		return
	}

	fmt.Print("Confirm booking? (y/n): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(utils.ErrInputInvalid)
		return
	}
	if !utils.ContainsString([]string{utils.Yes, utils.No}, strings.ToLower(strings.TrimSpace(input))) {
		fmt.Println(utils.ErrInputInvalid)
		return
	}

	if strings.TrimSpace(input) == utils.No {
		err = utils.ErrBookingCancelledMsg
		return
	}

	loggedUser := h.authenticationService.GetLoggedUser()

	saveBookingRequest := booking_model.SaveBookingRequest{
		FlightRoutes: bookedRoutes,
		UserID:       loggedUser.ID,
	}
	bookingResponse, err := h.bookingService.SaveBooking(saveBookingRequest)
	if err != nil {
		return
	}

	fmt.Printf(utils.BookingSuccessMessage, bookingResponse.BookingID)
	for _, flightRoute := range bookingResponse.FlightRoutes {
		for _, flightRouteSeat := range bookingResponse.FligthRouteSeats {
			if flightRoute.ID == flightRouteSeat.FlightRouteID {
				fmt.Printf(utils.BookingDetailMessage,
					flightRouteSeat.SeatNumber,
					flightRoute.DepartureCity.Name,
					flightRoute.DestinationCity.Name,
					flightRoute.DepartureDay,
					flightRoute.DepartureTime,
					flightRoute.Aircraft.Name,
				)
			}
		}
	}

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
