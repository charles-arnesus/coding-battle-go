package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	flight_model "github.com/charles-arnesus/coding-battle-go/models/flight"
	authentication_service "github.com/charles-arnesus/coding-battle-go/services/authentication"
	booking_service "github.com/charles-arnesus/coding-battle-go/services/booking"
	system_operation_service "github.com/charles-arnesus/coding-battle-go/services/systemOperation"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

type CancelFlightCommand struct {
	authenticationService authentication_service.AuthenticationService
	bookingService        booking_service.BookingService
	systemOperation       system_operation_service.SystemOperationService
}

func NewCancelFlightCommand(authenticationService authentication_service.AuthenticationService, bookingService booking_service.BookingService, systemOperation system_operation_service.SystemOperationService) *CancelFlightCommand {
	return &CancelFlightCommand{
		authenticationService: authenticationService,
		bookingService:        bookingService,
		systemOperation:       systemOperation,
	}
}

func (h *CancelFlightCommand) Execute() (err error) {
	currentDay := h.systemOperation.GetCurrentDay()

	bookingSystem, err := h.bookingService.GetBookingSystem()
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	if !bookingSystem.IsActive {
		err = utils.ErrBookingServiceDisabledMsg
		return
	}

	loggedUser := h.authenticationService.GetLoggedUser()

	response, err := h.bookingService.GetBookingDetails(loggedUser.ID, currentDay)
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	if len(response.BookingDetails) == 0 {
		fmt.Println(utils.NoActiveBookingMsg)
		return
	}

	fmt.Println("Your bookings: ")

	bookingIDs := []string{}
	for idx, bookingDetail := range response.BookingDetails {
		bookingIDs = append(bookingIDs, fmt.Sprintf("%d", bookingDetail.BookingID))
		fmt.Printf("%d. Booking ID: %d\n", idx+1, bookingDetail.BookingID)

		// Build maps for faster lookup
		flightRoutesByID := make(map[uint]flight_model.FlightRoute)
		for _, fr := range bookingDetail.FlightRoutes {
			flightRoutesByID[fr.ID] = fr
		}

		seatsByRouteID := make(map[uint][]flight_model.FlightRouteSeat)
		for _, seat := range bookingDetail.FlightRouteSeats {
			seatsByRouteID[seat.FlightRouteID] = append(seatsByRouteID[seat.FlightRouteID], seat)
		}

		// Loop booking flight routes
		for _, bookingFlightRoute := range bookingDetail.BookingFlightRoutes {
			flightRoute, ok := flightRoutesByID[bookingFlightRoute.FlightRouteID]
			if !ok {
				continue
			}

			// Print seats for this flight route
			for _, seat := range seatsByRouteID[flightRoute.ID] {
				fmt.Printf(utils.BookingDetailMessage,
					seat.SeatNumber,
					flightRoute.DepartureCity.Name,
					flightRoute.DestinationCity.Name,
					flightRoute.DepartureDay,
					flightRoute.DepartureTime,
					flightRoute.Aircraft.Name,
				)
			}
		}

		fmt.Println("")
	}

	fmt.Print("Select booking to cancel (Enter ID): ")
	reader := bufio.NewReaderSize(os.Stdin, 1)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(utils.ErrInputInvalid)
		return
	}

	if !utils.ContainsString(bookingIDs, strings.TrimSpace(input)) {
		fmt.Println(utils.ErrInputInvalid)
		return
	}

	// TODO
	//err := h.bookingService.CancelBooking(input)

	return
}

func (h *CancelFlightCommand) ID() string {
	return utils.PassengerCancelFlightID
}

func (h *CancelFlightCommand) AllowedRole() []string {
	return []string{
		utils.RolePassenger,
	}
}
