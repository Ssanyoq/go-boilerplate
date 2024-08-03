package utility

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ssanyoq/go-boilerplate/src/models"
	"os"
	"strconv"
)

// SetTokens sets necessary tokens for further authentication.
// It requires *gin.Context to set the token in cookies
//
// Sooner or later I'll add second token for complete OAuth thing ._.
func SetTokens(c *gin.Context, user models.User) error {
	payload := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedString, err := token.SignedString(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return err
	}
	jwtAge, _ := strconv.ParseInt(os.Getenv("JWT_AGE_S"), 10, 64)
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
