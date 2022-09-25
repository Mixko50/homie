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

	// * Paths ------------------------------------------------
	group := router.Group("group/")
	group.Get("info/all", groupHandler.GetAllGroups)
	group.Get("info/:id", groupHandler.GetGroupById)
	group.Post("create", groupHandler.CreateGroup)
}
