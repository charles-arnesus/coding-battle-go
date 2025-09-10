package handler

import (
	"github.com/charles-arnesus/coding-battle-go/command"
	"github.com/charles-arnesus/coding-battle-go/utils"
)

type Handler struct {
	commands map[string]command.Command
}

func NewHandler() *Handler {
	return &Handler{
		commands: make(map[string]command.Command),
	}
}

func (h *Handler) RegisterCommand(cmd command.Command) {
	h.commands[cmd.ID()] = cmd
}

func (h *Handler) ExecuteCommand(input, userRole string) error {
	ID := utils.ConvertInputToIDService(input, userRole)
	cmd, exists := h.commands[ID]
	if !exists {
		return utils.ErrCommandInvalid
	}

	isAllowed := utils.ContainsString(cmd.AllowedRole(), userRole)
	if !isAllowed {
		return utils.ErrRolePermission
	}

	return cmd.Execute()
}
