package service

import (
	"server/repository"
	"server/types/error_response"
	"server/types/request"
	"server/types/response"
)

type accessoryStateService struct {
	accessoryStateRepository repository.AccessoryStateRepository
}

func NewAccessoryStateService(accessoryStateRepository repository.AccessoryStateRepository) accessoryStateService {
	return accessoryStateService{accessoryStateRepository: accessoryStateRepository}
}

func (s accessoryStateService) GetAllAccessoryStates() ([]response.GetAccessoryStateResponse, error) {
	accessoryStates, err := s.accessoryStateRepository.GetAll()
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to get all accessory states",
		}
	}

	// * Convert to response
	var accessoryStateResponse []response.GetAccessoryStateResponse
	for _, v := range accessoryStates {
		accessoryStateResponse = append(accessoryStateResponse, response.GetAccessoryStateResponse{
			Id:          v.Id,
			AccessoryId: v.AccessoryId,
			MemberId:    v.MemberId,
			State:       v.State,
			CreatedAt:   v.CreatedAt,
			GroupId:     v.GroupId,
		})
	}

	return accessoryStateResponse, nil
}

func (s accessoryStateService) GetAccessoryStateById(id uint64) (*response.GetAccessoryStateResponse, error) {
	accessoryState, err := s.accessoryStateRepository.GetById(id)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to accessory state",
		}
	}

	// * Convert to response
	accessoryStateResponse := response.GetAccessoryStateResponse{
		Id:          accessoryState.Id,
		AccessoryId: accessoryState.AccessoryId,
		MemberId:    accessoryState.MemberId,
		State:       accessoryState.State,
		CreatedAt:   accessoryState.CreatedAt,
		GroupId:     accessoryState.GroupId,
	}

	return &accessoryStateResponse, nil
}

func (s accessoryStateService) GetAllAccessoryStatesInGroup(group uint64) ([]response.GetAccessoryStateResponse, error) {
	accessoryState, err := s.accessoryStateRepository.GetAllInGroup(group)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to get all accessory states",
		}
	}

	// * Convert to response
	var accessoryStateResponse []response.GetAccessoryStateResponse
	for _, v := range accessoryState {
		accessoryStateResponse = append(accessoryStateResponse, response.GetAccessoryStateResponse{
			Id:          v.Id,
			AccessoryId: v.AccessoryId,
			MemberId:    v.MemberId,
			State:       v.State,
			CreatedAt:   v.CreatedAt,
			GroupId:     v.GroupId,
		})
	}

	return accessoryStateResponse, nil
}

func (s accessoryStateService) GetAllAccessoryStatesInGroupByMember(groupId, memberId uint64) ([]response.GetAccessoryStateResponse, error) {
	accessoryState, err := s.accessoryStateRepository.GetAllInGroupByMember(groupId, memberId)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to get all accessory states",
		}
	}

	// * Convert to response
	var accessoryStateResponse []response.GetAccessoryStateResponse
	for _, v := range accessoryState {
		accessoryStateResponse = append(accessoryStateResponse, response.GetAccessoryStateResponse{
			Id:          v.Id,
			AccessoryId: v.AccessoryId,
			MemberId:    v.MemberId,
			State:       v.State,
			CreatedAt:   v.CreatedAt,
			GroupId:     v.GroupId,
		})
	}

	return accessoryStateResponse, nil
}

func (s accessoryStateService) GetAllAccessoryStatesInGroupByAccessory(accessoryId, groupId uint64) ([]response.GetAccessoryStateResponse, error) {
	accessoryState, err := s.accessoryStateRepository.GetAllByAccessoryIdInGroup(accessoryId, groupId)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to get all accessory states",
		}
	}

	// * Convert to response
	var accessoryStateResponse []response.GetAccessoryStateResponse
	for _, v := range accessoryState {
		accessoryStateResponse = append(accessoryStateResponse, response.GetAccessoryStateResponse{
			Id:          v.Id,
			AccessoryId: v.AccessoryId,
			MemberId:    v.MemberId,
			State:       v.State,
			CreatedAt:   v.CreatedAt,
			GroupId:     v.GroupId,
		})
	}

	return accessoryStateResponse, nil
}

func (s accessoryStateService) GetAllAccessoryStatesInGroupByMemberAndAccessory(accessoryId, groupId, memberId uint64) ([]response.GetAccessoryStateResponse, error) {
	accessoryState, err := s.accessoryStateRepository.GetAllByAccessoryIdAndMemberIdInGroup(accessoryId, groupId, memberId)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to get all accessory states",
		}
	}

	// * Convert to response
	var accessoryStateResponse []response.GetAccessoryStateResponse
	for _, v := range accessoryState {
		accessoryStateResponse = append(accessoryStateResponse, response.GetAccessoryStateResponse{
			Id:          v.Id,
			AccessoryId: v.AccessoryId,
			MemberId:    v.MemberId,
			State:       v.State,
			CreatedAt:   v.CreatedAt,
			GroupId:     v.GroupId,
		})
	}

	return accessoryStateResponse, nil
}

func (s accessoryStateService) CreateAccessoryState(request request.CreateAccessoryStateRequest, groupId, memberId uint64) error {
	err := s.accessoryStateRepository.CreateAccessoryState(request.AccessoryId, groupId, memberId, request.State)
	if err != nil {
		return &error_response.Error{
			Message: "Unable get all accessory states",
		}
	}
	return nil
}
