package service

import (
	"server/repository"
	"server/types/error_response"
	"server/types/request"
	"server/types/response"
	"server/types/secure"
	"server/utils/sign"
	"server/utils/text"
	"time"
)

type memberService struct {
	memberRepository repository.MemberRepository
}

func NewMemberService(memberRepository repository.MemberRepository) memberService {
	return memberService{memberRepository: memberRepository}
}

func (s memberService) GetAllMembers() ([]response.GetMemberResponse, error) {
	members, err := s.memberRepository.GetAll()
	if err != nil {
		return nil, &error_response.Error{
			Message: "Member not found",
			Err:     err,
		}
	}

	// * Convert to response
	var memberResponse []response.GetMemberResponse
	for _, v := range members {
		memberResponse = append(memberResponse, response.GetMemberResponse{
			Id:         v.Id,
			Name:       v.Name,
			DeviceName: v.DeviceName,
			CreatedAt:  text.FormatTime(v.CreatedAt),
			UpdatedAt:  text.FormatTime(v.UpdatedAt),
		})
	}

	return memberResponse, nil
}

func (s memberService) GetMemberById(id uint64) (*response.GetMemberResponse, error) {
	member, err := s.memberRepository.GetById(id)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Member not found",
			Err:     err,
		}
	}

	// * Convert to response
	memberResponse := response.GetMemberResponse{
		Id:         member.Id,
		Name:       member.Name,
		DeviceName: member.DeviceName,
		CreatedAt:  text.FormatTime(member.CreatedAt),
		UpdatedAt:  text.FormatTime(member.UpdatedAt),
	}

	return &memberResponse, nil
}

func (s memberService) CreateMember(request request.CreateMemberRequest, agent string) (*response.CreateMemberResponse, error) {
	isDuplicate, err := s.memberRepository.CheckDuplicateNameInGroup(request.Name, request.GroupId)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to check duplicate name",
			Err:     err,
		}
	}
	if isDuplicate {
		return nil, &error_response.Error{
			Message: "Duplicate name",
			Err:     err,
		}
	}

	member, err := s.memberRepository.CreateMember(request.Name, request.DeviceName, request.GroupId, agent, time.Now())
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to create member",
			Err:     err,
		}
	}

	// * Sign token
	signedToken, err := sign.SignJwt(secure.JwtHomieClaims{
		MemberId: member.Id,
		GroupId:  member.GroupId,
	})
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to sign token",
			Err:     err,
		}
	}

	// * Convert to response
	payload := response.CreateMemberResponse{
		Token: signedToken,
	}

	return &payload, nil
}
