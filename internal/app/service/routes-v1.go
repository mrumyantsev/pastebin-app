package service

import (
	"github.com/mrumyantsev/pastebin-app/internal/auth"
	"github.com/mrumyantsev/pastebin-app/internal/middleware"
	"github.com/mrumyantsev/pastebin-app/internal/paste"
	"github.com/mrumyantsev/pastebin-app/internal/user"
)

func (a *App) initRoutesV1() error {
	// Domain Handlers
	authHandler := auth.NewHttpHandler(a.services.Auth)
	userHandler := user.NewHttpHandler(a.services.User)
	pasteHandler := paste.NewHttpHandler(a.services.Paste)

	// Middleware
	logger := middleware.NewLogger()
	userAuthorization := middleware.NewUserAuthorization()

	// Routes
	auth := a.server.AlignedGroup(
		logger.Middleware(),
	)
	{
		auth.POST____("/v1/auth/sign-up", authHandler.SignUp)
		auth.POST____("/v1/auth/sign-in", authHandler.SignIn)
	}

	users := a.server.AlignedGroup(
		logger.Middleware(),
		userAuthorization.Middleware(),
	)
	{
		users.POST____("/v1/users", userHandler.CreateUser)
		users.GET_____("/v1/users", userHandler.GetAllUsers)
		users.GET_____("/v1/users/:id", userHandler.GetUserById)
		users.PUT_____("/v1/users/:id", userHandler.UpdateUserById)
		users.DELETE__("/v1/users/:id", userHandler.DeleteUserById)
		users.GET_____("/v1/users/is-username-exists/:username", userHandler.IsUserExistsByUsername)
	}

	pastes := a.server.AlignedGroup(
		logger.Middleware(),
		userAuthorization.Middleware(),
	)
	{
		pastes.POST____("/v1/pastes", pasteHandler.CreatePaste)
		pastes.GET_____("/v1/pastes", pasteHandler.GetAllPastes)
		pastes.GET_____("/v1/pastes/:base64-id", pasteHandler.GetPasteById)
		pastes.PUT_____("/v1/pastes/:base64-id", pasteHandler.UpdatePasteById)
		pastes.DELETE__("/v1/pastes/:base64-id", pasteHandler.DeletePasteById)
		pastes.GET_____("/v1/pastes/is-content-exists/:base64-id", pasteHandler.IsPasteContentExistsById)
	}

	return nil
}
