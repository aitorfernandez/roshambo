package jwt

import (
	"strconv"
	"time"

	"github.com/aitorfernandez/roshambo/pkg/env"
	"github.com/dgrijalva/jwt-go"
)

// JWT struct keeps details for JWT token generation.
type JWT struct {
	expires int
	secret  string
}

// New creates a JWT struct.
func New() *JWT {
	expires, _ := strconv.Atoi(env.MustHget("jwt", "expires"))
	return &JWT{
		expires: expires,
		secret:  env.MustHget("jwt", "secret"),
	}
}

// DefaultJWT is the default JWT used by for cache the env values.
var DefaultJWT = New()

// claims struct will be encoded to a JWT.
type claims struct {
	jwt.StandardClaims
}

// Make generates a token.
func (j JWT) Make(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * time.Duration(j.expires)).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   id,
		},
	})
	return token.SignedString([]byte(j.secret))
}

// Make is a wrapper around defaultJWT.Make.
func Make(id string) (string, error) {
	return DefaultJWT.Make(id)
}
