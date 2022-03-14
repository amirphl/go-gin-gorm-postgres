package middleware

import (
	"github.com/amirphl/go-gin-gorm-postgres/helper"
	"github.com/amirphl/go-gin-gorm-postgres/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
)

// AuthorizeJWT 400 401
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			resp := helper.BuildErrResp("No token found", "", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			return
		}

		token, err := jwtService.ValidateToken(authHeader)

		if err != nil {
			resp := helper.BuildErrResp("Token is not valid", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		// token.valid
		claims := token.Claims.(jwt.MapClaims)
		log.Println("valid Claims:", claims)
	}
}
