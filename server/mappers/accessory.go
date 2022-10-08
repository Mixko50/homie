package mappers

import (
	"server/repository"
	"server/types/response"
)

func MapAccessoriesResponse(accessories []repository.Accessory) []response.GetAccessoryResponse {
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

	return accessoriesResponse
}

func MapAccessoryResponse(accessory repository.Accessory) response.GetAccessoryResponse {
	accessoryResponse := response.GetAccessoryResponse{
		Id:        accessory.Id,
		Name:      accessory.Name,
		GroupId:   accessory.GroupId,
		GroupName: accessory.Group.Name,
		UpdatedAt: accessory.UpdatedAt,
		CreatedAt: accessory.CreatedAt,
	}
	return accessoryResponse
}
