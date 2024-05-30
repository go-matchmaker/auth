package psql

import (
	"auth/internal/core/domain/entity"
	"auth/internal/core/port/db"
	"auth/internal/core/port/user"
	"context"
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

var (
	_                 user.UserRepositoryPort = (*UserRepository)(nil)
	UserRepositorySet                         = wire.NewSet(NewUserRepository)
)

type UserRepository struct {
	dbPool *pgxpool.Pool
}

func NewUserRepository(em db.PostgresEngineMaker) user.UserRepositoryPort {
	return &UserRepository{
		dbPool: em.GetDB(),
	}
}

func (r *UserRepository) Insert(ctx context.Context, userModel *entity.User) (*uuid.UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	userModel.ID = id
	query := `INSERT INTO users (id, role, name, surname, email, phone_number, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, role, name, surname, email, phone_number, password_hash, created_at, updated_at`
	queryRow := r.dbPool.QueryRow(ctx, query, userModel.ID, userModel.Role, userModel.Name, userModel.Surname, userModel.Email, userModel.PhoneNumber, userModel.PasswordHash, userModel.CreatedAt, userModel.UpdatedAt)
	err = queryRow.Scan(&userModel.ID, &userModel.Role, &userModel.Name, &userModel.Surname, &userModel.Email, &userModel.PhoneNumber, &userModel.PasswordHash, &userModel.CreatedAt, &userModel.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *UserRepository) Update(ctx context.Context, userModel *entity.User) (*entity.User, error) {
	query := `
        UPDATE users
        SET
            name = COALESCE($1, name),
            surname = COALESCE($2, surname),
			role = COALESCE($3, role),
            email = COALESCE($4, email),
            phone_number = COALESCE($5, phone_number),
            password_hash = COALESCE($6, password_hash),
            updated_at = COALESCE($7, updated_at)
        WHERE
            id = $8
        RETURNING *
    `

	queryRow := r.dbPool.QueryRow(ctx, query, userModel.Name, userModel.Surname, userModel.Role, userModel.Email, userModel.PhoneNumber, userModel.PasswordHash, userModel.UpdatedAt, userModel.ID)

	updatedUser := new(entity.User)
	err := queryRow.Scan(&updatedUser.ID, &updatedUser.Role, &updatedUser.Name, &updatedUser.Surname, &updatedUser.Email, &updatedUser.PhoneNumber, &updatedUser.PasswordHash, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (r *UserRepository) UpdatePassword(ctx context.Context, id uuid.UUID, password string) (*entity.User, error) {
	query := `
		UPDATE users
		SET
			password_hash = $1,
			updated_at = $2
		WHERE
			id = $3
		RETURNING *
	`

	queryRow := r.dbPool.QueryRow(ctx, query, password, time.Now(), id)

	updatedUser := new(entity.User)
	err := queryRow.Scan(&updatedUser.ID, &updatedUser.Role, &updatedUser.Name, &updatedUser.Surname, &updatedUser.Email, &updatedUser.PhoneNumber, &updatedUser.PasswordHash, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (entity.User, error) {
	query := `SELECT id, role, name, surname, email, phone_number, password_hash, created_at, updated_at FROM users WHERE id = $1 LIMIT 1`
	queryRow := r.dbPool.QueryRow(ctx, query, id)
	userM := new(entity.User)
	err := queryRow.Scan(&userM.ID, &userM.Role, &userM.Name, &userM.Surname, &userM.Email, &userM.PhoneNumber, &userM.PasswordHash, &userM.CreatedAt, &userM.UpdatedAt)
	if err != nil {
		return entity.User{}, err
	}

	return *userM, nil

}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	query := `SELECT id, role, name, surname, email, phone_number, password_hash, created_at, updated_at FROM users WHERE email = $1 LIMIT 1`
	queryRow := r.dbPool.QueryRow(ctx, query, email)
	userM := new(entity.User)
	err := queryRow.Scan(&userM.ID, &userM.Role, &userM.Name, &userM.Surname, &userM.Email, &userM.PhoneNumber, &userM.PasswordHash, &userM.CreatedAt, &userM.UpdatedAt)
	if err != nil {
		return entity.User{}, err
	}

	return *userM, nil
}

func (r *UserRepository) DeleteOne(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.dbPool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteAll(ctx context.Context) error {
	query := `DELETE FROM users`
	_, err := r.dbPool.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
