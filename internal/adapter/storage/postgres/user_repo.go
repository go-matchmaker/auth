package psql

import (
	"auth/internal/core/domain/aggregate"
	"auth/internal/core/domain/entity"
	"auth/internal/core/port/db"
	"auth/internal/core/port/user"
	"context"
	"time"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
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

func (r *UserRepository) GetByID(ctx context.Context, id string) (*aggregate.UserAggregate, error) {
	const query = `SELECT 
	u.id, u.role, u.name, u.surname, u.email, u.phone_number, u.password, u.created_at, u.updated_at,
	d.id AS department_id, d.name AS department_name, d.created_at AS department_created_at,
	ua.attribute, ua.view, ua.search, ua.detail, ua.add, ua.update, ua.delete, ua.export, ua.upload, ua.can_see_price
    FROM users u
    JOIN departments d ON u.department_id = d.id
    LEFT JOIN user_attributes ua ON u.id = ua.user_id
    WHERE u.id = $1;
    `
	queryRows, err := r.dbPool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	var userModel *aggregate.UserAggregate
	userModel.Permissions = make(map[string]*entity.Permission)

	for queryRows.Next() {
		var (
			role                                                                              entity.Role
			departmentID, departmentName, userID, name, surname, email, phoneNumber, password string
			createdAt, updatedAt, departmentCreatedAt                                         time.Time
			attribute                                                                         string
			view, search, detail, add, update, delete, export, upload, canSeePrice            bool
		)

		err := queryRows.Scan(
			&userID, &role, &name, &surname, &email, &phoneNumber, &password, &createdAt, &updatedAt,
			&departmentID, &departmentName, &departmentCreatedAt,
			&attribute, &view, &search, &detail, &add, &update, &delete, &export, &upload, &canSeePrice,
		)
		if err != nil {
			return nil, err
		}

		userModel.User.ID = userID
		userModel.User.Role = role
		userModel.User.Name = name
		userModel.User.Surname = surname
		userModel.User.Email = email
		userModel.User.PhoneNumber = phoneNumber
		userModel.User.Password = password
		userModel.User.CreatedAt = createdAt
		userModel.User.UpdatedAt = updatedAt
		userModel.Department = &entity.Department{
			ID:        departmentID,
			Name:      departmentName,
			CreatedAt: departmentCreatedAt,
		}
		userModel.Permissions[attribute] = &entity.Permission{
			View:        view,
			Search:      search,
			Detail:      detail,
			Add:         add,
			Update:      update,
			Delete:      delete,
			Export:      export,
			Import:      upload,
			CanSeePrice: canSeePrice,
		}
	}
	return userModel, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*aggregate.UserAggregate, error) {
	const query = `SELECT 
	u.id, u.role, u.name, u.surname, u.email, u.phone_number, u.password, u.created_at, u.updated_at,
	d.id AS department_id, d.name AS department_name, d.created_at AS department_created_at,
	ua.attribute, ua.view, ua.search, ua.detail, ua.add, ua.update, ua.delete, ua.export, ua.upload, ua.can_see_price
    FROM users u
    JOIN departments d ON u.department_id = d.id
    LEFT JOIN user_attributes ua ON u.id = ua.user_id
    WHERE u.email = $1;
    `

	queryRows, err := r.dbPool.Query(ctx, query, email)
	if err != nil {
		return nil, err
	}
	var userModel *aggregate.UserAggregate
	userModel.Permissions = make(map[string]*entity.Permission)

	for queryRows.Next() {
		var (
			role                                                                              entity.Role
			departmentID, departmentName, userID, name, surname, email, phoneNumber, password string
			createdAt, updatedAt, departmentCreatedAt                                         time.Time
			attribute                                                                         string
			view, search, detail, add, update, delete, export, upload, canSeePrice            bool
		)

		err := queryRows.Scan(
			&userID, &role, &name, &surname, &email, &phoneNumber, &password, &createdAt, &updatedAt,
			&departmentID, &departmentName, &departmentCreatedAt,
			&attribute, &view, &search, &detail, &add, &update, &delete, &export, &upload, &canSeePrice,
		)
		if err != nil {
			return nil, err
		}

		userModel.User.ID = userID
		userModel.User.Role = role
		userModel.User.Name = name
		userModel.User.Surname = surname
		userModel.User.Email = email
		userModel.User.PhoneNumber = phoneNumber
		userModel.User.Password = password
		userModel.User.CreatedAt = createdAt
		userModel.User.UpdatedAt = updatedAt
		userModel.Department = &entity.Department{
			ID:        departmentID,
			Name:      departmentName,
			CreatedAt: departmentCreatedAt,
		}
		userModel.Permissions[attribute] = &entity.Permission{
			View:        view,
			Search:      search,
			Detail:      detail,
			Add:         add,
			Update:      update,
			Delete:      delete,
			Export:      export,
			Import:      upload,
			CanSeePrice: canSeePrice,
		}
	}
	return userModel, nil
}

func (r *UserRepository) GetUserPassword(ctx context.Context, email string) (string, error) {
	query := `SELECT password FROM users where email = $1`
	queryRow := r.dbPool.QueryRow(ctx, query, email)
	password := ""
	err := queryRow.Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}
