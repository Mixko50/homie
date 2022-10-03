package service

import (
	"server/types/request"
	"server/types/response"
)

type AccessoryStateService interface {
	GetAllAccessoryStates() ([]response.GetAccessoryStateResponse, error)
	GetAccessoryStateById(uint64) (*response.GetAccessoryStateResponse, error)
	GetAllAccessoryStatesInGroup(uint64) ([]response.GetAccessoryStateResponse, error)
	GetAllAccessoryStatesInGroupByMember(uint64, uint64) ([]response.GetAccessoryStateResponse, error)
	GetAllAccessoryStatesInGroupByAccessory(uint64, uint64) ([]response.GetAccessoryStateResponse, error)
	GetAllAccessoryStatesInGroupByMemberAndAccessory(uint64, uint64, uint64) ([]response.GetAccessoryStateResponse, error)
	CreateAccessoryState(request request.CreateAccessoryStateRequest, groupId, memberId uint64) error
}
