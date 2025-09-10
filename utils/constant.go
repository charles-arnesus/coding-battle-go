package utils

import "fmt"

const RoleAdmin = "admin"
const RoleAdminLabel = "Admin"
const RolePassenger = "passenger"
const RolePassengerLabel = "Passenger"

// ID service
var AdminRegisterAircraftID = fmt.Sprintf("%s_1", RoleAdmin)
var AdminAddDestinationID = fmt.Sprintf("%s_2", RoleAdmin)

// menu
const RegisterAircraftLabel = "Register Aircraft"
const AddDestinationLabel = "Add Destination"
const CreateFlightRouteLabel = "Create Flight Route"
const RunBookingServiceLabel = "Run Booking Service"
const GoToNextDayLabel = "Go to Next Day"
const RunFlightLabel = "Run Flight"
const ExitLabel = "Exit"

const BookAFlightLabel = "Book a Flight"
const CancelABookingLabel = "Cancel a Booking"

var MenuAdmin = []string{
	RegisterAircraftLabel, AddDestinationLabel, CreateFlightRouteLabel, RunBookingServiceLabel, GoToNextDayLabel, RunFlightLabel,
}

var MenuPassenger = []string{
	BookAFlightLabel, CancelABookingLabel,
}
