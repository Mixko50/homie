package service

import (
	"gorm.io/gorm"
	"server/repository"
	"server/types/error_response"
	"server/types/request"
	"server/types/response"
	"server/utils/bcrypt"
	"server/utils/text"
	"time"
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

func (s groupService) GetGroupById(id uint64) (*response.GetGroupResponse, error) {
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

func (s groupService) CreateGroup(request request.CreateGroupRequest) error {
	// * Hash password
	hashedPassword, err := bcrypt.HashPassword(request.Password)
	if err != nil {
		return &error_response.Error{
			Message: "Unable to perform the password",
		}
	}

	// * Check duplicate name
	isDuplicate, err := s.groupRepository.CheckDuplicateName(request.Name)
	if err != nil {
		return &error_response.Error{
			Message: "Unable to check duplicate name",
		}
	} else {
		if isDuplicate {
			return &error_response.Error{
				Message: "Duplicate name",
			}
		}
	}

	// * Create group
	if err := s.groupRepository.CreateGroup(request.Name, hashedPassword, text.FormatTime(time.Now())); err != nil {
		return &error_response.Error{
			Message: "Unable to create group",
		}
	}
	return nil
}
