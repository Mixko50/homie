package request

type CreateMemberRequest struct {
	Name           string `json:"name" validate:"required"`
	GroupId        uint64 `json:"group_id" validate:"required"`
	GroupPassword  string `json:"group_password" validate:"required"`
	MemberPassword string `json:"member_password" validate:"required,min=8,max=255"`
}

type GetMemberTokenRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}
