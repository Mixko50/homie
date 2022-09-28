package request

type CreateMemberRequest struct {
	Name          string `json:"name" validate:"required"`
	GroupId       uint64 `json:"group_id" validate:"required"`
	GroupPassword string `json:"group_password" validate:"required"`
}
