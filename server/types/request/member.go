package request

type CreateMemberRequest struct {
	Name       string `json:"name" validate:"required"`
	DeviceName string `json:"device_name" validate:"required"`
	GroupId    uint64 `json:"group_id" validate:"required"`
}
