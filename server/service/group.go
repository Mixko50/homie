package service

import (
	"server/types/request"
	"server/types/response"
)

type GroupService interface {
	GetAllGroups() ([]response.GetGroupResponse, error)
	GetGroupById(uint) (*response.GetGroupResponse, error)
	CreateGroup(request request.CreateGroupRequest) error
}
