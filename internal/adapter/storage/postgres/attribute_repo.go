package psql

import (
	"auth/internal/core/domain/entity"
	"auth/internal/core/port/db"
	"context"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	_                      attribute.AttributeRepositoryPort = (*AttributeRepository)(nil)
	AttributeRepositorySet                                   = wire.NewSet(NewAttributeRepository)
)

type AttributeRepository struct {
	dbPool *pgxpool.Pool
}

func NewAttributeRepository(em db.PostgresEngineMaker) attribute.AttributeRepositoryPort {
	return &AttributeRepository{
		dbPool: em.GetDB(),
	}
}

func (r *AttributeRepository) GetAll(ctx context.Context) ([]entity.Attribute, error) {
	query := `SELECT * FROM attributes`
	rows, err := r.dbPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attributes []entity.Attribute
	for rows.Next() {
		var attribute entity.Attribute
		err = rows.Scan(&attribute.ID, &attribute.Name, &attribute.CreatedAt)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, attribute)
	}

	return attributes, nil
}
