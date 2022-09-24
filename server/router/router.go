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
	group := router.Group("group/info/")
	group.Get("all", groupHandler.GetAllGroups)
	group.Get(":id", groupHandler.GetGroupById)
}
