package repositories

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/danzBraham/eniqilo-store/internal/entities/userentity"
	"github.com/danzBraham/eniqilo-store/internal/errors/usererror"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	IsPhoneNumberByItsRoleExists(ctx context.Context, phoneNumber string, role userentity.Role) (bool, error)
	CreateUser(ctx context.Context, user *userentity.User) error
	GetUserByPhoneNumberAndRole(ctx context.Context, phoneNumber string, role userentity.Role) (*userentity.User, error)
	GetCustomers(ctx context.Context, params *userentity.CustomerQueryParams) ([]*userentity.GetCustomerResponse, error)
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
		&user.ID,
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

func (r *UserRepositoryImpl) GetUserByPhoneNumberAndRole(ctx context.Context, phoneNumber string, role userentity.Role) (*userentity.User, error) {
	query := `
		SELECT
			id,
			phone_number,
			name,
			password,
			role,
			created_at
		FROM
			users
		WHERE
			phone_number = $1
			AND role = $2
			AND is_deleted = false
	`
	var user userentity.User
	var createdAt time.Time
	err := r.DB.QueryRow(ctx, query, phoneNumber, role).Scan(
		&user.ID,
		&user.PhoneNumber,
		&user.Name,
		&user.Password,
		&user.Role,
		&createdAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, usererror.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	user.CreatedAt = createdAt.Format(time.RFC3339)
	return &user, nil
}

func (r *UserRepositoryImpl) GetCustomers(ctx context.Context, params *userentity.CustomerQueryParams) ([]*userentity.GetCustomerResponse, error) {
	query := `
		SELECT
			id,
			phone_number,
			name
		FROM
			users
		WHERE
			role = 'Customer'
			AND is_deleted = false
	`
	args := []interface{}{}
	argID := 1

	if params.PhoneNumber != "" {
		query += ` AND phone_number LIKE $` + strconv.Itoa(argID)
		args = append(args, `+`+params.PhoneNumber+`%`)
		argID++
	}

	if params.Name != "" {
		query += ` AND name ILIKE $` + strconv.Itoa(argID)
		args = append(args, `%`+params.Name+`%`)
		argID++
	}

	query += ` ORDER BY created_at DESC`
	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := make([]*userentity.GetCustomerResponse, 0)
	for rows.Next() {
		var customer userentity.GetCustomerResponse
		rows.Scan(
			&customer.UserID,
			&customer.PhoneNumber,
			&customer.Name,
		)
		customers = append(customers, &customer)
	}

	return customers, nil
}
