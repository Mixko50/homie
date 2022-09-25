package request

type CreateGroupRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}
