package command

import (
	"fmt"

	booking_service "github.com/charles-arnesus/coding-battle-go/services/booking"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

type BookFlightCommand struct {
	bookingService booking_service.BookingService
}

func NewBookFlightCommand(bookingService booking_service.BookingService) *BookFlightCommand {
	return &BookFlightCommand{
		bookingService: bookingService,
	}
}

func (h *BookFlightCommand) Execute() (err error) {
	bookingSystem, err := h.bookingService.GetBookingSystem()
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	if !bookingSystem.IsActive {
		err = utils.ErrBookingServiceDisabledMsg
		return
	}

	fmt.Println("TODO: Book a flight")

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
