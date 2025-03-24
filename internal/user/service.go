package user

import "context"

type Service struct {
	database DatabaseAdapterer
}

func NewService(databaseAdapter DatabaseAdapterer) *Service {
	return &Service{
		database: databaseAdapter,
	}
}

func (s *Service) CreateUser(ctx context.Context, outerUser OuterUser) (int64, error) {
	user, err := outerUser.ToUser()
	if err != nil {
		return 0, err
	}

	id, err := s.database.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) GetAllUsers(ctx context.Context) ([]OuterUser, error) {
	users, err := s.database.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	outerUsers := make([]OuterUser, len(users))

	for i := range outerUsers {
		outerUsers[i], err = users[i].ToOuterUser()
		if err != nil {
			return nil, err
		}
	}

	return outerUsers, nil
}

func (s *Service) GetUserById(ctx context.Context, id int64) (OuterUser, error) {
	user, err := s.database.GetUserById(ctx, id)
	if err != nil {
		return OuterUser{}, err
	}

	outerUser, err := user.ToOuterUser()
	if err != nil {
		return OuterUser{}, err
	}

	return outerUser, nil
}

func (s *Service) UpdateUserById(ctx context.Context, id int64, outerUser OuterUser) error {
	user, err := outerUser.ToUser()
	if err != nil {
		return err
	}

	if err = s.database.UpdateUserById(ctx, id, user); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUserById(ctx context.Context, id int64) error {
	err := s.database.DeleteUserById(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) IsUserExistsByUsername(ctx context.Context, username string) (bool, error) {
	isExists, err := s.database.IsUserExistsByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	return isExists, nil
}

func (s *Service) GetIdAndPasswordByUsername(ctx context.Context, username string) (int64, string, error) {
	id, password, err := s.database.GetIdAndPasswordByUsername(ctx, username)
	if err != nil {
		return -1, "", err
	}

	return id, password, nil
}
