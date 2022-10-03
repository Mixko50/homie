package service

import (
	"server/types/request"
	"server/types/response"
)

type AccessoryService interface {
	GetAllAccessories() ([]response.GetAccessoryResponse, error)
	GetAccessoryById(uint64) (*response.GetAccessoryResponse, error)
	CreateAccessory(request request.CreateAccessoryRequest, groupId uint64) error
	GetAllAccessoriesInGroup(uint64) ([]response.GetAccessoryResponse, error)
}
