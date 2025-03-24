package user

import (
	"net/http"

	"github.com/mrumyantsev/pastebin-app/internal/reqerrors"
)

var (
	ErrFieldUsernameIsEmpty  = reqerrors.New(http.StatusBadRequest, "field 'username' is empty")
	ErrFieldPasswordIsEmpty  = reqerrors.New(http.StatusBadRequest, "field 'password' is empty")
	ErrInvalidPasswordLength = reqerrors.New(http.StatusBadRequest, "invalid password length: min: 8, max: 60")
	ErrFieldEmailIsEmpty     = reqerrors.New(http.StatusBadRequest, "field 'email' is empty")
	ErrEmailNotMatchPattern  = reqerrors.New(http.StatusBadRequest, "email not match 'johndoe@mail.com' pattern")

	ErrUserAlreadyExists = reqerrors.New(http.StatusBadRequest, "a user with the given username already exists")
	ErrUserNotFound      = reqerrors.New(http.StatusNotFound, "user not found")
)
