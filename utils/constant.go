package utils

import "fmt"

// max days
const MaxDaysInYear = 365

const RoleAdmin = "admin"
const RoleAdminLabel = "Admin"
const RolePassenger = "passenger"
const RolePassengerLabel = "Passenger"

// ID service
// admin
var AdminRegisterAircraftID = fmt.Sprintf("%s_1", RoleAdmin)
var AdminAddDestinationID = fmt.Sprintf("%s_2", RoleAdmin)
var AdminCreateFligthRouteID = fmt.Sprintf("%s_3", RoleAdmin)
var AdminSetBookingSystemID = fmt.Sprintf("%s_4", RoleAdmin)
var AdminGoToNextDaySystemID = fmt.Sprintf("%s_5", RoleAdmin)

// passenger
var PassengerBookFlightID = fmt.Sprintf("%s_1", RolePassenger)
var PassengerCancelFlightID = fmt.Sprintf("%s_2", RolePassenger)

// menu
const RegisterAircraftLabel = "Register Aircraft"
const AddDestinationLabel = "Add Destination"
const CreateFlightRouteLabel = "Create Flight Route"
const RunBookingServiceLabel = "Enable/Disable Booking Service"
const GoToNextDayLabel = "Go to Next Day"
const RunFlightLabel = "Run Flight"
const ExitLabel = "Exit"

const BookAFlightLabel = "Book a Flight"
const CancelABookingLabel = "Cancel a Booking"

// status
const SCHEDULED = "SCHEDULED"
const DEPARTED = "DEPARTED"
const ARRIVED = "ARRIVED"

// yes no
const Yes = "y"
const No = "n"

var MenuAdmin = []string{
	RegisterAircraftLabel, AddDestinationLabel, CreateFlightRouteLabel, RunBookingServiceLabel, GoToNextDayLabel, RunFlightLabel,
}

var MenuPassenger = []string{
	BookAFlightLabel, CancelABookingLabel,
}

var FlightStatus = []string{
	SCHEDULED, DEPARTED, ARRIVED,
}

var UnavailableFlightStatus = []string{
	DEPARTED, ARRIVED,
}
