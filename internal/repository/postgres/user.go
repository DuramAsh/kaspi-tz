package postgres

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"kaspi-tz/internal/domain/person"
)

type PersonRepository struct {
	DB *sqlx.DB
}

func NewPersonRepository(db *sqlx.DB) *PersonRepository {
	return &PersonRepository{
		DB: db,
	}
}

func (r *PersonRepository) InsertPerson(ctx context.Context, data person.Entity) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		INSERT INTO people (name, iin, phone)
		VALUES ($1, $2, $3)`

	args := []any{data.Name, data.IIN, data.Phone}

	_, err = r.DB.ExecContext(ctx, query, args...)

	return
}

func (r *PersonRepository) GetPersonByIIN(ctx context.Context, iin string) (dest person.Entity, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		SELECT name, iin, phone FROM people
		WHERE iin = $1`

	args := []any{iin}

	err = r.DB.GetContext(ctx, &dest, query, args...)

	return
}

func (r *PersonRepository) GetPeopleByNamePart(ctx context.Context, name string) (dest []person.Entity, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		SELECT name, iin, phone
		FROM people
		WHERE REPLACE(LOWER(name), ' ', '') LIKE '%' || LOWER(REPLACE($1, ' ', '')) || '%'`

	args := []any{name}

	err = r.DB.SelectContext(ctx, &dest, query, args...)

	return
}
