package service

import (
	"gorm.io/gorm"
	"server/repository"
	"server/types/error_response"
	"server/types/response"
)

type groupService struct {
	groupRepository repository.GroupRepository
}

func NewGroupService(groupRepository repository.GroupRepository) groupService {
	return groupService{groupRepository: groupRepository}
}

func (s groupService) GetAllGroups() ([]response.GetGroupResponse, error) {
	groups, err := s.groupRepository.GetAll()
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, &error_response.Error{
				Message: "Group not found",
			}
		}
		return nil, &error_response.Error{
			Message: "Unable to get groups",
		}
	}

	// * Convert to response
	var groupResponse []response.GetGroupResponse
	for _, v := range groups {
		groupResponse = append(groupResponse, response.GetGroupResponse{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return groupResponse, nil
}

func (s groupService) GetGroupById(id uint) (*response.GetGroupResponse, error) {
	group, err := s.groupRepository.GetById(id)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, &error_response.Error{
				Message: "Group not found",
			}
		}
		return nil, &error_response.Error{
			Message: "Unable to get group",
		}
	}

	// * Convert to response
	groupResponse := response.GetGroupResponse{
		Id:   group.Id,
		Name: group.Name,
	}

	return &groupResponse, nil
}
