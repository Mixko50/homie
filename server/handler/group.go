package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/service"
	"server/types/info_response"
	"strconv"
)

type groupHandler struct {
	groupService service.GroupService
}

func NewGroupHandler(groupService service.GroupService) groupHandler {
	return groupHandler{groupService: groupService}
}

func (h groupHandler) GetAllGroups(c *fiber.Ctx) error {
	groups, err := h.groupService.GetAllGroups()
	if err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(groups))
}

func (h groupHandler) GetGroupById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	group, err := h.groupService.GetGroupById(uint(id))
	if err != nil {
		return err
	}

	return c.JSON(info_response.NewInfoResponse(group))
}
