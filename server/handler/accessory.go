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

type accessoryHandler struct {
	accessoryService service.AccessoryService
}

func NewAccessoryHandler(accessoryService service.AccessoryService) accessoryHandler {
	return accessoryHandler{accessoryService: accessoryService}
}

func (h accessoryHandler) GetAllAccessories(c *fiber.Ctx) error {
	accessories, err := h.accessoryService.GetAllAccessories()
	if err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(accessories))
}

func (h accessoryHandler) GetAllAccessoryInGroup(c *fiber.Ctx) error {
	claims := c.Locals("home").(*jwt.Token).Claims.(*secure.ClaimsStruct)
	accessories, err := h.accessoryService.GetAllAccessoriesInGroup(claims.GroupId)
	if err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(accessories))
}

func (h accessoryHandler) GetAccessoryById(c *fiber.Ctx) error {
	accessoryId, err := strconv.ParseUint(c.Params("accessory_id"), 10, 64)
	if err != nil {
		return err
	}
	accessories, err := h.accessoryService.GetAccessoryById(accessoryId)
	if err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(accessories))
}

func (h accessoryHandler) CreateAccessory(c *fiber.Ctx) error {
	claims := c.Locals("home").(*jwt.Token).Claims.(*secure.ClaimsStruct)
	body := new(request.CreateAccessoryRequest)
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

	// * Service
	if err := h.accessoryService.CreateAccessory(*body, claims.GroupId); err != nil {
		return err
	}

	return c.JSON(info_response.NewInfoResponse("Accessory created"))
}
