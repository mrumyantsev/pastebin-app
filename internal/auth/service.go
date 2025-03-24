package auth

import (
	"context"

	"github.com/mrumyantsev/pastebin-app/internal/jwttokens"
	"github.com/mrumyantsev/pastebin-app/internal/passhash"
	"github.com/mrumyantsev/pastebin-app/internal/user"
)

type Service struct {
	userDatabase user.DatabaseAdapterer
}

func NewService(userDatabaseAdapter user.DatabaseAdapterer) *Service {
	return &Service{
		userDatabase: userDatabaseAdapter,
	}
}

func (s *Service) SignUp(ctx context.Context, outerUser user.OuterUser) (string, error) {
	user, err := outerUser.ToUser()
	if err != nil {
		return "", err
	}

	id, err := s.userDatabase.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	token, err := jwttokens.NewUserIdToken(id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) SignIn(ctx context.Context, outerAuth OuterAuth) (string, error) {
	id, hashedPassword, err := s.userDatabase.GetIdAndPasswordByUsername(ctx, outerAuth.Username)
	if err != nil {
		return "", err
	}

	if id < 0 {
		return "", ErrIncorrectUsernameOrPassword
	}

	isPasswordMatch, err := passhash.IsPasswordsMatch(hashedPassword, outerAuth.Password)
	if err != nil {
		return "", err
	}

	if !isPasswordMatch {
		return "", ErrIncorrectUsernameOrPassword
	}

	token, err := jwttokens.NewUserIdToken(id)
	if err != nil {
		return "", err
	}

	return token, nil
}
