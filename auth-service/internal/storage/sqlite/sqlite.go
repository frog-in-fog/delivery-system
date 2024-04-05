package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"github.com/frog-in-fog/delivery-system/auth-service/internal/models"
	"github.com/frog-in-fog/delivery-system/auth-service/internal/storage"
)

type sqliteDB struct {
	db *sql.DB
}

func NewSQLiteStorage(storagePath string) (storage.UserStorage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	return &sqliteDB{db: db}, nil
}

func (s *sqliteDB) CreateUser(ctx context.Context, user *models.User) error {
	stmt, err := s.db.Prepare("INSERT INTO users(email, password_hash) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, user.Email, user.PasswordHash)
	if err != nil {
		//var sqliteErr sqlite3.Error
		//if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
		//	return errors.New("user already exists")
		//
		return err
	}

	return nil
}

func (s *sqliteDB) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	stmt, err := s.db.Prepare("SELECT id, email, password_hash FROM users WHERE email = ?")
	if err != nil {
		return &models.User{}, err
	}

	row := stmt.QueryRowContext(ctx, email)
	var user *models.User

	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.User{}, err
		}

		return &models.User{}, err
	}

	return user, nil
}

func (s *sqliteDB) GetUserById(ctx context.Context, id string) (*models.User, error) {
	stmt, err := s.db.Prepare("SELECT id, email, password_hash FROM users WHERE id = ?")
	if err != nil {
		return &models.User{}, err
	}

	row := stmt.QueryRowContext(ctx, id)
	var user *models.User

	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.User{}, err
		}

		return &models.User{}, err
	}

	return user, nil
}