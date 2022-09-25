package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/service"
	"server/types/error_response"
	"server/types/info_response"
	"server/types/request"
	"server/utils/validator"
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
	// * Get id
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	// * Get group
	group, err := h.groupService.GetGroupById(uint(id))
	if err != nil {
		return err
	}

	return c.JSON(info_response.NewInfoResponse(group))
}

func (h groupHandler) CreateGroup(c *fiber.Ctx) error {
	body := new(request.CreateGroupRequest)
	if err := c.BodyParser(body); err != nil {
		return &error_response.Error{
			Message: "Unable to parse information",
		}
	}

	// * Validate
	validateErr := validator.Validate.Struct(body)
	if validateErr != nil {
		return validateErr
	}

	err := h.groupService.CreateGroup(*body)
	if err != nil {
		return err
	}

	return c.JSON(info_response.NewInfoResponse("Group created"))
}
