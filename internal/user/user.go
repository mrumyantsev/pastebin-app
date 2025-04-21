package user

import (
	"context"
	"net/mail"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/go-errlib"
	"github.com/mrumyantsev/pastebin-app/internal/passhash"
)

type Servicer interface {
	CreateUser(ctx context.Context, outerUser OuterUser) (int64, error)
	GetAllUsers(ctx context.Context) ([]OuterUser, error)
	GetUserById(ctx context.Context, id int64) (OuterUser, error)
	UpdateUserById(ctx context.Context, id int64, outerUser OuterUser) error
	DeleteUserById(ctx context.Context, id int64) error

	IsUserExistsByUsername(ctx context.Context, username string) (bool, error)
	GetIdAndPasswordByUsername(ctx context.Context, username string) (int64, string, error)
}

type DatabaseAdapterer interface {
	CreateUser(ctx context.Context, user User) (int64, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserById(ctx context.Context, id int64) (User, error)
	UpdateUserById(ctx context.Context, id int64, user User) error
	DeleteUserById(ctx context.Context, id int64) error

	IsUserExistsByUsername(ctx context.Context, username string) (bool, error)
	GetIdAndPasswordByUsername(ctx context.Context, username string) (int64, string, error)
}

type User struct {
	Id        int64
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
}

func (u User) ToOuterUser() (OuterUser, error) {
	return OuterUser{
		Id:        u.Id,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Format(time.DateTime),
	}, nil
}

type OuterUser struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

func NewOuterUser(c *gin.Context) (OuterUser, error) {
	var user OuterUser

	err := c.Bind(&user)
	if err != nil {
		return OuterUser{}, errlib.Wrap(err, "could not parse json body")
	}

	if user.Username == "" {
		return OuterUser{}, ErrFieldUsernameIsEmpty
	}

	if user.Password == "" {
		return OuterUser{}, ErrFieldPasswordIsEmpty
	}

	passwordLen := len(user.Password)

	if passwordLen < 8 || passwordLen > 60 {
		return OuterUser{}, ErrInvalidPasswordLength
	}

	if user.Email == "" {
		return OuterUser{}, ErrFieldEmailIsEmpty
	}

	if _, err = mail.ParseAddress(user.Email); err != nil {
		return OuterUser{}, ErrEmailNotMatchPattern
	}

	return user, nil
}

func (u OuterUser) ToUser() (User, error) {
	passwordHash, err := passhash.Generate(u.Password)
	if err != nil {
		return User{}, err
	}

	return User{
		Id:       u.Id,
		Username: u.Username,
		Password: passwordHash,
		Email:    u.Email,
	}, nil
}
