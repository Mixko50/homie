package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"server/service"
	"server/types/error_response"
	"server/types/info_response"
	"server/types/request"
	"server/types/secure"
	"server/utils/validator"
	"strconv"
)

type accessoryStateHandler struct {
	accessoryStateService service.AccessoryStateService
}

func NewAccessoryStateHandler(accessoryStateService service.AccessoryStateService) accessoryStateHandler {
	return accessoryStateHandler{accessoryStateService: accessoryStateService}
}

func (h accessoryStateHandler) GetAllAccessoryStates(c *fiber.Ctx) error {
	accessoryStates, err := h.accessoryStateService.GetAllAccessoryStates()
	if err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(accessoryStates))
}

func (h accessoryStateHandler) GetAccessoryStateById(c *fiber.Ctx) error {
	accessoryStateId, err := strconv.ParseUint(c.Params("accessory_state_id"), 10, 64)
	if err != nil {
		return err
	}
	accessoryState, err := h.accessoryStateService.GetAccessoryStateById(accessoryStateId)
	if err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(accessoryState))
}

func (h accessoryStateHandler) GetAllAccessoryStatesInGroup(c *fiber.Ctx) error {
	claims := c.Locals("home").(*jwt.Token).Claims.(*secure.ClaimsStruct)
	accessoryStates, err := h.accessoryStateService.GetAllAccessoryStatesInGroup(claims.GroupId)
	if err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(accessoryStates))
}

func (h accessoryStateHandler) GetAllAccessoryStatesInGroupByMember(c *fiber.Ctx) error {
	claims := c.Locals("home").(*jwt.Token).Claims.(*secure.ClaimsStruct)
	accessoryStates, err := h.accessoryStateService.GetAllAccessoryStatesInGroupByMember(claims.GroupId, claims.MemberId)
	if err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(accessoryStates))
}

func (h accessoryStateHandler) GetAllAccessoryStatesInGroupByAccessory(c *fiber.Ctx) error {
	claims := c.Locals("home").(*jwt.Token).Claims.(*secure.ClaimsStruct)
	accessoryId, err := strconv.ParseUint(c.Params("accessory_id"), 10, 64)
	if err != nil {
		return err
	}
	accessoryStates, err := h.accessoryStateService.GetAllAccessoryStatesInGroupByAccessory(claims.GroupId, accessoryId)
	if err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(accessoryStates))
}

func (h accessoryStateHandler) GetAllAccessoryStatesInGroupByMemberAndAccessory(c *fiber.Ctx) error {
	claims := c.Locals("home").(*jwt.Token).Claims.(*secure.ClaimsStruct)
	accessoryId, err := strconv.ParseUint(c.Params("accessory_id"), 10, 64)
	if err != nil {
		return err
	}
	accessoryStates, err := h.accessoryStateService.GetAllAccessoryStatesInGroupByMemberAndAccessory(claims.GroupId, claims.MemberId, accessoryId)
	if err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(accessoryStates))
}

func (h accessoryStateHandler) CreateAccessoryState(c *fiber.Ctx) error {
	claims := c.Locals("home").(*jwt.Token).Claims.(*secure.ClaimsStruct)

	// * Parse body
	body := new(request.CreateAccessoryStateRequest)
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

	if err := h.accessoryStateService.CreateAccessoryState(*body, claims.GroupId, claims.MemberId); err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(nil))
}
