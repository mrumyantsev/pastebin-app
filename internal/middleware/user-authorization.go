package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mrumyantsev/pastebin-app/internal/jwttokens"
	"github.com/mrumyantsev/pastebin-app/internal/server"
)

type UserAuthorization struct {
}

func NewUserAuthorization() *UserAuthorization {
	return &UserAuthorization{}
}

func (a *UserAuthorization) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if header == "" {
			server.HttpErrorResponse(c, errors.New("header 'Authorization' is empty"))

			return
		}

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			server.HttpErrorResponse(c, errors.New("authorization header has invalid data"))

			return
		}

		if headerParts[1] == "" {
			server.HttpErrorResponse(c, errors.New("token is empty"))

			return
		}

		userId, err := jwttokens.ParseUserIdToken(headerParts[1])
		if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
			server.HttpErrorResponse(c, err)

			return
		}

		if userId >= 0 {
			server.SetUserId(c, userId)
		}

		c.Next()
	}
}
