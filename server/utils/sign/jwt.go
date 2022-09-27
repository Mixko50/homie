package sign

import (
	"github.com/golang-jwt/jwt/v4"
	"server/types/secure"
	"server/utils/config"
)

func SignJwt(claims secure.JwtHomieClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.C.JwtSecret))
}
