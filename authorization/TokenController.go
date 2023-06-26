package authorization

import (
	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
type JWTHandle struct {
		
}
