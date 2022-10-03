package router

import (
	"github.com/gofiber/fiber/v2"
	"server/handler"
	"server/loaders/database"
	"server/repository"
	"server/service"
	"server/utils/middleware"
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

	// * Accessory
	accessoryRepository := repository.NewAccessoryRepositoryDB(database.DB)
	accessoryService := service.NewAccessoryService(accessoryRepository, groupRepository)
	accessoryHandler := handler.NewAccessoryHandler(accessoryService)

	// * Accessory State
	accessoryStateRepository := repository.NewAccessoryStateRepository(database.DB)
	accessoryStateService := service.NewAccessoryStateService(accessoryStateRepository)
	accessoryStateHandler := handler.NewAccessoryStateHandler(accessoryStateService)

	// * Paths ------------------------------------------------
	group := router.Group("group/")
	group.Get("info/all", groupHandler.GetAllGroups)
	group.Get("info/:id", groupHandler.GetGroupById)
	group.Post("create", groupHandler.CreateGroup)

	member := router.Group("member/")
	member.Get("info/all", memberHandler.GetAllMembers)
	member.Get("info/:member_id", memberHandler.GetMemberById)
	member.Post("create", memberHandler.CreateMember)
	member.Post("info/token", memberHandler.GetMemberToken)

	accessory := router.Group("accessory/", middleware.Jwt())
	accessory.Get("info/all", accessoryHandler.GetAllAccessories)
	accessory.Get("info/group", accessoryHandler.GetAllAccessoryInGroup)
	accessory.Get("info/:accessory_id", accessoryHandler.GetAccessoryById)
	accessory.Post("create", accessoryHandler.CreateAccessory)

	accessoryState := router.Group("accessory/state/", middleware.Jwt())
	accessoryState.Get("info/all", accessoryStateHandler.GetAllAccessoryStates)
	accessoryState.Get("info/:accessory_state_id", accessoryStateHandler.GetAccessoryStateById)
	accessoryState.Get("info/group", accessoryStateHandler.GetAllAccessoryStatesInGroup)
	accessoryState.Get("info/group/member", accessoryStateHandler.GetAllAccessoryStatesInGroupByMember)
	accessoryState.Get("info/group/accessory", accessoryStateHandler.GetAllAccessoryStatesInGroupByAccessory)
	accessoryState.Get("info/group/member/accessory", accessoryStateHandler.GetAllAccessoryStatesInGroupByMemberAndAccessory)
	accessoryState.Post("create", accessoryStateHandler.CreateAccessoryState)
}
