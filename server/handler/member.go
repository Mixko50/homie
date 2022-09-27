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

type memberHandler struct {
	memberService service.MemberService
}

func NewMemberHandler(memberService service.MemberService) memberHandler {
	return memberHandler{memberService: memberService}
}

func (h memberHandler) GetAllMembers(c *fiber.Ctx) error {
	members, err := h.memberService.GetAllMembers()
	if err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(members))
}

func (h memberHandler) GetMemberById(c *fiber.Ctx) error {
	// * Get id
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	// * Get member
	member, err := h.memberService.GetMemberById(id)
	if err != nil {
		return err
	}
	return c.JSON(info_response.NewInfoResponse(member))
}

func (h memberHandler) CreateMember(c *fiber.Ctx) error {
	body := new(request.CreateMemberRequest)
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

	result, err := h.memberService.CreateMember(*body, c.Get("user-agent"))
	if err != nil {
		return err
	}

	return c.JSON(info_response.NewInfoResponse(result))
}
