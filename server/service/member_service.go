package service

import (
	"server/repository"
	"server/types/error_response"
	"server/types/request"
	"server/types/response"
	"server/types/secure"
	"server/utils/bcrypt"
	"server/utils/sign"
	"time"
)

type memberService struct {
	memberRepository repository.MemberRepository
	groupRepository  repository.GroupRepository
}

func NewMemberService(memberRepository repository.MemberRepository, groupRepository repository.GroupRepository) memberService {
	return memberService{memberRepository: memberRepository, groupRepository: groupRepository}
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
			Id:        v.Id,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
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
		Id:        member.Id,
		Name:      member.Name,
		CreatedAt: member.CreatedAt,
		UpdatedAt: member.UpdatedAt,
	}

	return &memberResponse, nil
}

func (s memberService) CreateMember(request request.CreateMemberRequest, agent string) (*response.CreateMemberResponse, error) {
	// * Check group password
	group, err := s.groupRepository.GetByIdWithPassword(request.GroupId)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Group not found",
			Err:     err,
		}
	}
	if !bcrypt.ComparePassword(group.Password, request.GroupPassword) {
		return nil, &error_response.Error{
			Message: "Group password is incorrect",
			Err:     err,
		}
	}

	isDuplicate, err := s.memberRepository.CheckDuplicateNameInMember(request.Name, request.GroupId)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to check duplicate name",
			Err:     err,
		}
	}
	if isDuplicate {
		return nil, &error_response.Error{
			Message: "Duplicated name",
			Err:     err,
		}
	}

	// * Hash password
	hashedPassword, err := bcrypt.HashPassword(request.MemberPassword)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Unable to perform the password",
		}
	}

	member, err := s.memberRepository.CreateMember(request.Name, hashedPassword, request.GroupId, agent, time.Now())
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

func (s memberService) GetMemberToken(request request.GetMemberTokenRequest) (*response.GetMemberTokenResponse, error) {
	members, err := s.memberRepository.GetMembersByName(request.Name)
	if err != nil {
		return nil, &error_response.Error{
			Message: "Member not found",
			Err:     err,
		}
	}

	for _, member := range members {
		if bcrypt.ComparePassword(member.Password, request.Password) && request.Name == member.Name {
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
			payload := response.GetMemberTokenResponse{
				Token: signedToken,
			}

			return &payload, nil
		}
	}

	return nil, &error_response.Error{
		Message: "Member not found",
		Err:     err,
	}
}
