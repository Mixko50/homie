package service

import (
	"server/repository"
	"server/types/error_response"
	"server/types/request"
	"server/types/response"
	"server/utils/bcrypt"
	"time"
)

type accessoryService struct {
	accessoryRepository repository.AccessoryRepository
	groupRepository     repository.GroupRepository
}

func NewAccessoryService(accessoryRepository repository.AccessoryRepository, groupRepository repository.GroupRepository) accessoryService {
	return accessoryService{accessoryRepository: accessoryRepository, groupRepository: groupRepository}
}

func (s accessoryService) GetAllAccessories() ([]response.GetAccessoryResponse, error) {
	accessories, err := s.accessoryRepository.GetAll()
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to get accessories",
			Err:     err,
		}
	}

	// * Convert to response
	var accessoriesResponse []response.GetAccessoryResponse
	for _, accessory := range accessories {
		accessoriesResponse = append(accessoriesResponse, response.GetAccessoryResponse{
			Id:        accessory.Id,
			Name:      accessory.Name,
			GroupId:   accessory.GroupId,
			GroupName: accessory.Group.Name,
			UpdatedAt: accessory.UpdatedAt,
			CreatedAt: accessory.CreatedAt,
		})
	}

	return accessoriesResponse, nil
}

func (s accessoryService) GetAccessoryById(id uint64) (*response.GetAccessoryResponse, error) {
	accessory, err := s.accessoryRepository.GetById(id)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to get group",
			Err:     err,
		}
	}

	// * Convert to response
	accessoryResponse := response.GetAccessoryResponse{
		Id:        accessory.Id,
		Name:      accessory.Name,
		GroupId:   accessory.GroupId,
		GroupName: accessory.Group.Name,
		UpdatedAt: accessory.UpdatedAt,
		CreatedAt: accessory.CreatedAt,
	}

	return &accessoryResponse, nil
}

func (s accessoryService) CreateAccessory(request request.CreateAccessoryRequest, groupId uint64) error {
	group, err := s.groupRepository.GetByIdWithPassword(groupId)
	if err != nil {
		return &error_response.Error{
			Message: "Unable to get group",
			Err:     err,
		}
	}
	if !bcrypt.ComparePassword(group.Password, request.GroupPassword) {
		return &error_response.Error{
			Message: "Group password is incorrect",
			Err:     err,
		}
	}

	// * Check duplicate name
	accessory, err := s.accessoryRepository.CheckDuplicateNameInGroup(request.Name, groupId)
	if err != nil {
		return &error_response.Error{
			Message: "Unable to check the name in accessory",
			Err:     err,
		}
	}
	if accessory {
		return &error_response.Error{
			Message: "Duplicated name",
			Err:     err,
		}
	}

	if result := s.accessoryRepository.CreateAccessory(request.Name, groupId, time.Now()); result != nil {
		return &error_response.Error{
			Message: "Unable to add accessory",
			Err:     err,
		}
	}
	return nil
}

func (s accessoryService) GetAllAccessoriesInGroup(id uint64) ([]response.GetAccessoryResponse, error) {
	accessories, err := s.accessoryRepository.GetAllInGroup(id)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to get accessories",
			Err:     err,
		}
	}

	// * Convert to response
	var accessoriesResponse []response.GetAccessoryResponse
	for _, accessory := range accessories {
		accessoriesResponse = append(accessoriesResponse, response.GetAccessoryResponse{
			Id:        accessory.Id,
			Name:      accessory.Name,
			GroupId:   accessory.GroupId,
			GroupName: accessory.Group.Name,
			UpdatedAt: accessory.UpdatedAt,
			CreatedAt: accessory.CreatedAt,
		})
	}

	return accessoriesResponse, nil
}
