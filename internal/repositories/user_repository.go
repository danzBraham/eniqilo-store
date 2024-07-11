package repositories

import (
	"context"
	"errors"

	"github.com/danzBraham/eniqilo-store/internal/entities/userentity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	IsPhoneNumberByItsRoleExists(ctx context.Context, phoneNumber string, role userentity.Role) (bool, error)
	CreateUser(ctx context.Context, user *userentity.User) error
}

type UserRepositoryImpl struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) IsPhoneNumberByItsRoleExists(ctx context.Context, phoneNumber string, role userentity.Role) (bool, error) {
	query := `
		SELECT
			1
		FROM
			users
		WHERE
			phone_number = $1
			AND role = $2
			AND is_deleted = false
	`
	var exists int
	err := r.DB.QueryRow(ctx, query, phoneNumber, role).Scan(&exists)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *userentity.User) error {
	query := `
		INSERT INTO
			users (id, phone_number, name, password, role)
		VALUES
			($1, $2, $3, $4, $5)
	`
	_, err := r.DB.Exec(ctx, query,
		&user.Id,
		&user.PhoneNumber,
		&user.Name,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return err
	}
	return nil
}
