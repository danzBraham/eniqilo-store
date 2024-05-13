package postgres_repositories

import (
	"context"

	"github.com/danzBraham/eniqilo-store/internal/domains/entities"
	"github.com/danzBraham/eniqilo-store/internal/domains/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StaffRepositoryDB struct {
	DB *pgxpool.Pool
}

func NewStaffRepositoryDB(db *pgxpool.Pool) repositories.StaffRepository {
	return &StaffRepositoryDB{DB: db}
}

func (r *StaffRepositoryDB) CreateStaff(staff *entities.RegisterStaff) (id string, err error) {
	query := "INSERT INTO staff (phone_number, name, password) VALUES ($1, $2, $3) RETURNING id"
	err = r.DB.QueryRow(context.Background(), query, staff.PhoneNumber, staff.Name, staff.Password).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, err
}

func (r *StaffRepositoryDB) VerifyPhoneNumber(phoneNumber string) (bool, error) {
	var isPhoneNumberExists bool
	query := "SELECT EXISTS (SELECT phone_number FROM staff WHERE phone_number = $1)"
	if err := r.DB.QueryRow(context.Background(), query, phoneNumber).Scan(&isPhoneNumberExists); err != nil {
		return false, err
	}

	return isPhoneNumberExists, nil
}

func (r *StaffRepositoryDB) FindByPhoneNumber(phoneNumber string) (*entities.Staff, error) {
	var id, name string
	query := "SELECT id, name FROM staff WHERE phone_number = $1 LIMIT 1"
	if err := r.DB.QueryRow(context.Background(), query, phoneNumber).Scan(&id, &name); err != nil {
		return nil, err
	}

	return &entities.Staff{
		ID:          id,
		Name:        name,
		PhoneNumber: phoneNumber,
	}, nil
}
