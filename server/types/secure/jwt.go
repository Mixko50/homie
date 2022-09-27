package secure

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtHomieClaims struct {
	MemberId uint64 `json:"member_id"`
	GroupId  uint64 `json:"group_id"`
	jwt.RegisteredClaims
}
