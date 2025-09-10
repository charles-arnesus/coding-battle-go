package utils

import "errors"

const CommandInvalidMsg = "command is invalid"
const RolePermissionMsg = "This role don't have a permission to go to this page"

var ErrCommandInvalid = errors.New(CommandInvalidMsg)
var ErrRolePermission = errors.New(RolePermissionMsg)
