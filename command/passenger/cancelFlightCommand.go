package command

import (
	"fmt"

	booking_service "github.com/charles-arnesus/coding-battle-go/services/booking"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

type CancelFlightCommand struct {
	bookingService booking_service.BookingService
}

func NewCancelFlightCommand(bookingService booking_service.BookingService) *CancelFlightCommand {
	return &CancelFlightCommand{
		bookingService: bookingService,
	}
}

func (h *CancelFlightCommand) Execute() (err error) {
	bookingSystem, err := h.bookingService.GetBookingSystem()
	if err != nil {
		err = utils.ErrSomethingWentWrongGet
		return
	}

	if !bookingSystem.IsActive {
		err = utils.ErrBookingServiceDisabledMsg
		return
	}

	fmt.Println("TODO: Cancel a flight")

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
