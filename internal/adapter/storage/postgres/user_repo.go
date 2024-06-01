package psql

import (
	"auth/internal/core/domain/entity"
	"auth/internal/core/port/db"
	"auth/internal/core/port/user"
	"context"
	"fmt"
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

func (r *UserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	query := `SELECT 
	u.id AS user_id,
	u.role,
	u.name,
	u.surname,
	u.email,
	u.phone_number,
	u.created_at AS user_created_at,
	u.updated_at AS user_updated_at,
	d.id AS department_id,
	d.name AS department_name,
	d.created_at AS department_created_at,
	d.updated_at AS department_updated_at,
	a.id AS attribute_id,
	a.name AS attribute_name,
	p.view,
	p.search,
	p.detail,
	p.add,
	p.update,
	p.delete,
	p.export,
	p.import,
	p.can_see_price,
	p.created_at AS permission_created_at,
	p.updated_at AS permission_updated_at
    FROM 
	users u
    LEFT JOIN 
	permissions p ON u.id = p.user_id
    LEFT JOIN 
	attributes a ON p.attribute_id = a.id
    LEFT JOIN 
	departments d ON a.department_id = d.id
    WHERE 
	u.id = $1;`
	queryRow := r.dbPool.QueryRow(ctx, query, id)
	userM := new(entity.User)
	err := queryRow.Scan(&userM.ID, &userM.Role, &userM.Name, &userM.Surname, &userM.Email, &userM.PhoneNumber, &userM.Password, &userM.CreatedAt, &userM.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return userM, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, id string) (*entity.User, error) {
	const query = `SELECT 
	u.id, u.role, u.name, u.surname, u.email, u.phone_number, u.password, u.created_at, u.updated_at,
	d.id AS department_id, d.name AS department_name, d.created_at AS department_created_at,
	ua.attribute, ua.view, ua.search, ua.detail, ua.add, ua.update, ua.delete, ua.export, ua.upload, ua.can_see_price
    FROM users u
    JOIN departments d ON u.department_id = d.id
    LEFT JOIN user_attributes ua ON u.id = ua.user_id
    WHERE u.email = $1;
    `

	queryRows, err := r.dbPool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	var user entity.User
	user.UserPermissions = make(map[string]entity.Permission)

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
			return &entity.User{}, err
		}

		user.ID = userID
		user.Role = role
		user.Name = name
		user.Surname = surname
		user.Email = email
		user.PhoneNumber = phoneNumber
		user.Password = password
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
		user.Department = entity.Department{
			ID:        departmentID,
			Name:      departmentName,
			CreatedAt: departmentCreatedAt,
		}
		user.UserPermissions[attribute] = entity.Permission{
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
	fmt.Println("user1", user)
	return &user, nil
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
