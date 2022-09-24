package service

import "server/types/response"

type GroupService interface {
	GetAllGroups() ([]response.GetGroupResponse, error)
	GetGroupById(uint) (*response.GetGroupResponse, error)
}
