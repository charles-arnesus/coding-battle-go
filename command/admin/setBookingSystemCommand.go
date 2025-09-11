package command

import (
	"fmt"

	booking_service "github.com/charles-arnesus/coding-battle-go/services/booking"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

type SetBookingSystemCommand struct {
	bookingService booking_service.BookingService
}

func NewSetBookingSystemCommand(bookingService booking_service.BookingService) *SetBookingSystemCommand {
	return &SetBookingSystemCommand{
		bookingService: bookingService,
	}
}

func (h *SetBookingSystemCommand) Execute() (err error) {
	bookingSystem, err := h.bookingService.GetBookingSystem()
	if err != nil {
		return
	}

	err = h.bookingService.SetBookingSystem(bookingSystem)
	if err != nil {
		return
	}

	if bookingSystem.IsActive {
		fmt.Println(utils.BookingServiceDisabledMessage)
	} else {
		fmt.Println(utils.BookingServiceEnabledMessage)
	}

	return
}

func (h *SetBookingSystemCommand) ID() string {
	return utils.AdminSetBookingSystemID
}

func (h *SetBookingSystemCommand) AllowedRole() []string {
	return []string{
		utils.RoleAdmin,
	}
}
