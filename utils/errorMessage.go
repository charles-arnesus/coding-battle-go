package utils

import (
	"errors"
)

// error message
const InputInvalidMessage = "the input is invalid"
const CommandInvalidMessage = "command is invalid"
const RolePermissionMessage = "this role don't have a permission to go to this page"
const SeatsAircraftCapacityMessage = "the seat capacity should be number and greater than 0"
const NameAircraftRequiredMessage = "aircraft name is required"
const NameAircraftAlreadyExistMessage = "aircraft with this name already exists"
const NameDestinationAlreadyExistMessage = "destination with this name already exists"
const NameDestinationRequiredMsg = "destination name is required"
const UsernameAlreadyExistMessage = "username already exists"
const RecordNotFoundMessage = "record not found"
const SomethingWentWrongGetMessage = "something went wrong when getting the data"
const DestinationsDataEmptyMessage = "the data destination is still empty"
const FlightRouteAlreadyExistsMessage = "the flight route is already exists"
const AircraftDataEmptyMessage = "the data aircraft is still empty"
const ScheduledDayInvalidMessage = "the scheduled day should be between current day (%d) - (%d)"
const AircraftInOtherCityMessage = "cant create the flight route, because on that time aircraft in other city"
const DepartureDestinationSameMessage = "the departure and destination city should be different"
const BookingServiceDisabledMsg = "booking service is disabled. contact admin to enable booking service"
const NoDirectFlightFoundMsg = "No direct flights found, searching for transit options..."
const NoFlightFoundMsg = "No flights found for this route"
const BookingCancelledMsg = "booking cancelled"
const NoSeatsAvailableMsg = "no seats available"
const NoActiveBookingMsg = "you do not have any active booking"
const DuplicateFlightRouteBookingMsg = "you already book this flight route"
const DuplicateFlightRouteMsg = "this aircraft already have schedule in selected departure day and time"

// error object
var ErrInputInvalid = errors.New(InputInvalidMessage)
var ErrCommandInvalid = errors.New(CommandInvalidMessage)
var ErrRolePermission = errors.New(RolePermissionMessage)
var ErrSeatsAircraftCapacity = errors.New(SeatsAircraftCapacityMessage)
var ErrNameAircraftRequired = errors.New(NameAircraftRequiredMessage)
var ErrNameAircraftAlreadyExist = errors.New(NameAircraftAlreadyExistMessage)
var ErrNameDestinationAlreadyExist = errors.New(NameDestinationAlreadyExistMessage)
var ErrNameDestinationRequired = errors.New(NameDestinationRequiredMsg)
var ErrUsernameAlreadyExistMsg = errors.New(UsernameAlreadyExistMessage)
var ErrDuplicateFlightRouteBookingMsg = errors.New(DuplicateFlightRouteBookingMsg)
var ErrDuplicateFlightRouteMsg = errors.New(DuplicateFlightRouteMsg)
var ErrSomethingWentWrongGet = errors.New(SomethingWentWrongGetMessage)
var ErrDestinationsDataEmpty = errors.New(DestinationsDataEmptyMessage)
var ErrAircraftDataEmpty = errors.New(AircraftDataEmptyMessage)
var ErrRecordNotFound = errors.New(RecordNotFoundMessage)
var ErrDepartureDestinationSame = errors.New(DepartureDestinationSameMessage)
var ErrBookingServiceDisabledMsg = errors.New(BookingServiceDisabledMsg)
var ErrBookingCancelledMsg = errors.New(BookingCancelledMsg)
var ErrNoSeatsAvailableMsg = errors.New(NoSeatsAvailableMsg)
var ErrFlightRouteAlreadyExistMsg = errors.New(FlightRouteAlreadyExistsMessage)
var ErrAircraftInOtherCity = errors.New(AircraftInOtherCityMessage)

// error code message
const UniqueViolationCodePostgres = "23505"
