package service

import (
	"server/types/request"
	"server/types/response"
)

type MemberService interface {
	GetAllMembers() ([]response.GetMemberResponse, error)
	GetMemberById(uint64) (*response.GetMemberResponse, error)
	CreateMember(request request.CreateMemberRequest, agent string) (*response.CreateMemberResponse, error)
}
