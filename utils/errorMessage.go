package utils

import "errors"

const InputInvalidMsg = "cannot read the input"
const CommandInvalidMsg = "command is invalid"
const RolePermissionMsg = "this role don't have a permission to go to this page"
const SeatsAircraftCapacityMsg = "the seat capacity should be number and greater than 0"
const NameAircraftRequiredMsg = "aircraft name is required"
const NameAircraftAlreadyExistMsg = "aircraft with this name already exists"
const NameDestinationAlreadyExistMsg = "destination with this name already exists"
const NameDestinationRequiredMsg = "destination name is required"
const UsernameAlreadyExistMsg = "username already exists"

var ErrInputInvalid = errors.New(InputInvalidMsg)
var ErrCommandInvalid = errors.New(CommandInvalidMsg)
var ErrRolePermission = errors.New(RolePermissionMsg)
var ErrSeatsAircraftCapacity = errors.New(SeatsAircraftCapacityMsg)
var ErrNameAircraftRequired = errors.New(NameAircraftRequiredMsg)
var ErrNameAircraftAlreadyExist = errors.New(NameAircraftAlreadyExistMsg)
var ErrNameDestinationAlreadyExist = errors.New(NameDestinationAlreadyExistMsg)
var ErrNameDestinationRequired = errors.New(NameDestinationRequiredMsg)
var ErrUsernameAlreadyExistMsg = errors.New(UsernameAlreadyExistMsg)

const UniqueViolationCodePostgres = "23505"
