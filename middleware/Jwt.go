package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secret := "dslldmaslkda090erejfei"
		author := ctx.Request.Header.Get("Authorization")
		tokenString := strings.Replace(author, "Bearer ", "", -1)

		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorize"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, errors.New("Unauthorize")
			}
			return []byte(secret), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !token.Valid || !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Error": gin.H{"Error": http.StatusText(http.StatusForbidden)}})
			return
		}

		exp := claims["exp"].(float64)

		if time.Now().After(time.Unix(int64(exp), 0)) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Error": "Token is Expired"})
			return
		}

		no := claims["Nim"].(string)

		if no == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Error": http.StatusText(http.StatusForbidden)})
			return
		}

		uid := claims["Uid"].(float64)

		ctx.Set("Uid", uid)
		ctx.Next()
	}
}
