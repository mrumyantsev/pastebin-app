package auth

import (
	"net/http"

	"github.com/mrumyantsev/pastebin-app/internal/reqerrors"
)

var (
	ErrIncorrectUsernameOrPassword = reqerrors.New(http.StatusBadRequest, "incorrect username or password")
)
