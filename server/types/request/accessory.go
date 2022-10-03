package request

type CreateAccessoryRequest struct {
	Name          string `json:"name" validate:"required"`
	GroupPassword string `json:"group_password" validate:"required"`
}
