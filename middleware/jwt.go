package middleware

import (
	"log"
	"strings"
	"weather_forecast/infrastructure"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret []byte

func init() {
	jwtSecret = []byte(infrastructure.GetSecret())
}
func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		var claims *jwt.StandardClaims
		log.Println(token[:7])
		if strings.ToLower(token[:7]) != "bearer " {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
		tokkenValue := token[7:]
		jwtClaims, err := jwt.ParseWithClaims(tokkenValue, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if jwtClaims == nil || err != nil {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		if err := jwtClaims.Claims.Valid(); err != nil {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
		claims, _ = jwtClaims.Claims.(*jwt.StandardClaims)
		c.Set("user", claims.Audience)
		c.Next()
	}
}
