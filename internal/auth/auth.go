package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/go-errlib"
	"github.com/mrumyantsev/pastebin-app/internal/user"
)

type Servicer interface {
	SignUp(ctx context.Context, outerUser user.OuterUser) (string, error)
	SignIn(ctx context.Context, outerAuth OuterAuth) (string, error)
}

type OuterAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewOuterAuth(c *gin.Context) (OuterAuth, error) {
	var auth OuterAuth

	err := c.Bind(&auth)
	if err != nil {
		return OuterAuth{}, errlib.Wrap(err, "could not parse json body")
	}

	if auth.Username == "" {
		return OuterAuth{}, user.ErrFieldUsernameIsEmpty
	}

	if auth.Password == "" {
		return OuterAuth{}, user.ErrFieldPasswordIsEmpty
	}

	passwordLen := len(auth.Password)

	if passwordLen < 8 || passwordLen > 60 {
		return OuterAuth{}, user.ErrInvalidPasswordLength
	}

	return auth, nil
}
