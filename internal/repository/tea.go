package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/levchenki/tea-app/internal/entity"
	"github.com/pkg/errors"
)

type TeaRepository struct {
	db *sqlx.DB
}

func NewTeaRepository(db *sqlx.DB) *TeaRepository {
	return &TeaRepository{db: db}
}

func (r *TeaRepository) Get() ([]entity.Tea, error) {
	teas := make([]entity.Tea, 0)
	err := r.db.Select(
		&teas,
		`select id, name, description, price, updated_at, id_category from teas`,
	)
	if err != nil {
		return nil, errors.Wrap(err, "get all teas")
	}
	return teas, nil
}

func (r *TeaRepository) GetById(id int) (entity.Tea, error) {
	tea := entity.Tea{}
	err := r.db.Get(
		&tea,
		"select id, name, description, price, updated_at, id_category from teas where id = $1",
		id,
	)
	if err != nil {
		return entity.Tea{}, errors.Wrap(err, "get tea by id")
	}
	return tea, nil
}

func (r *TeaRepository) GetByCategoryId(id int) ([]entity.Tea, error) {
	teas := make([]entity.Tea, 0)
	err := r.db.Select(
		&teas,
		"select id, name, description, price, updated_at, id_category from teas where id_category = $1",
		id,
	)
	if err != nil {
		return nil, errors.Wrap(err, "get all teas by category id")
	}
	return teas, nil
}
