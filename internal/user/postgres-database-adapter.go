package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mrumyantsev/pastebin-app/internal/database"
)

type PostgresDatabaseAdapter struct {
	db *database.PostgresDatabase
}

func NewPostgresDatabaseAdapter(db *database.PostgresDatabase) *PostgresDatabaseAdapter {
	return &PostgresDatabaseAdapter{
		db: db,
	}
}

func (a *PostgresDatabaseAdapter) CreateUser(ctx context.Context, user User) (int64, error) {
	const query = "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id"

	row := a.db.QueryRow(ctx, query, user.Username, user.Password, user.Email)

	var id int64
	var pgErr *pgconn.PgError

	err := row.Scan(&id)
	if errors.As(err, &pgErr) {
		return 0, ErrUserAlreadyExists
	}
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (a *PostgresDatabaseAdapter) GetAllUsers(ctx context.Context) ([]User, error) {
	const query = "SELECT id, username, password, email, created_at FROM users"

	rows, err := a.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	users := []User{}
	var user User

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (a *PostgresDatabaseAdapter) GetUserById(ctx context.Context, id int64) (User, error) {
	const query = "SELECT id, username, password, email, created_at FROM users WHERE id = $1"

	row := a.db.QueryRow(ctx, query, id)

	var user User

	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, ErrUserNotFound
	}
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (a *PostgresDatabaseAdapter) UpdateUserById(ctx context.Context, id int64, user User) error {
	const query = "UPDATE users SET username = $1, password = $2, email = $3 WHERE id = $4"

	_, err := a.db.Exec(ctx, query, user.Username, user.Password, user.Email, id)
	if err != nil {
		return err
	}

	return nil
}

func (a *PostgresDatabaseAdapter) DeleteUserById(ctx context.Context, id int64) error {
	const query = "DELETE FROM users WHERE id = $1"

	_, err := a.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (a *PostgresDatabaseAdapter) IsUserExistsByUsername(ctx context.Context, username string) (bool, error) {
	const query = "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)"

	row := a.db.QueryRow(ctx, query, username)

	var isExists bool

	err := row.Scan(&isExists)
	if err != nil {
		return false, err
	}

	return isExists, nil
}

func (a *PostgresDatabaseAdapter) GetIdAndPasswordByUsername(ctx context.Context, username string) (int64, string, error) {
	const query = "SELECT id, password FROM users WHERE username = $1"

	row := a.db.QueryRow(ctx, query, username)

	var id int64
	var password string

	err := row.Scan(&id, &password)
	if errors.Is(err, sql.ErrNoRows) {
		return -1, "", nil
	}
	if err != nil {
		return -1, "", err
	}

	return id, password, nil
}
