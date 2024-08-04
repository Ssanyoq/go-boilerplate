package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ssanyoq/go-boilerplate/src/utility"
	"net/http"
	"os"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Unauthorized"})
		return
	}
	var claims jwt.MapClaims
	_, err = jwt.ParseWithClaims(
		tokenString,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
	if err != nil {
		// oof tried to forge this token
		utility.ClearTokens(c)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Unauthorized, bad token"})
		return
	}
	c.Set("jwt", claims)
	c.Next()
}
