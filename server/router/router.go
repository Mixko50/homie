package router

import (
	"github.com/gofiber/fiber/v2"
	"server/handler"
	"server/loaders/database"
	"server/repository"
	"server/service"
)

func Router(router fiber.Router) {
	// * Register --------------------------------------------

	// * Group
	groupRepository := repository.NewGroupRepositoryDb(database.DB)
	groupService := service.NewGroupService(groupRepository)
	groupHandler := handler.NewGroupHandler(groupService)

	// * Member
	memberRepository := repository.NewMemberRepositoryDb(database.DB)
	memberService := service.NewMemberService(memberRepository, groupRepository)
	memberHandler := handler.NewMemberHandler(memberService)

	// * Paths ------------------------------------------------
	group := router.Group("group/")
	group.Get("info/all", groupHandler.GetAllGroups)
	group.Get("info/:id", groupHandler.GetGroupById)
	group.Post("create", groupHandler.CreateGroup)

	member := router.Group("member/")
	member.Get("info/all", memberHandler.GetAllMembers)
	member.Get("info/:id", memberHandler.GetMemberById)
	member.Post("create", memberHandler.CreateMember)
}
