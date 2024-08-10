package utility

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ssanyoq/go-boilerplate/src/models"
	"os"
	"strconv"
	"time"
)

type JWTClaims struct {
	id        int64
	name      string
	expiresAt int
}

// SetTokens sets necessary tokens for further authentication.
// It requires *gin.Context to set the token in cookies
//
// Sooner or later I'll add second token for complete OAuth thing ._.
func SetTokens(c *gin.Context, user models.User) error {
	jwtAge, _ := strconv.ParseInt(os.Getenv("JWT_AGE_S"), 10, 64)
	payload := jwt.MapClaims{
		"id":        user.ID,
		"name":      user.Name,
		"expiresAt": time.Now().Second() + int(jwtAge),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return err
	}
	c.SetCookie(
		"jwt",
		signedString,
		int(jwtAge),
		"/",
		"localhost",
		os.Getenv("ENV") == "PRODUCTION",
		true,
	)
	return nil
}

// ClearTokens clears tokens that are made by
// SetTokens. Use to punish for incorrect tokens
// (suspicions on token forgery) or for incorrect
// authorization data
func ClearTokens(c *gin.Context) error {
	c.SetCookie("jwt",
		"oops sorry",
		1, // minimal age possible
		"/",
		"localhost",
		os.Getenv("ENV") == "PRODUCTION",
		true)
	return nil
}
