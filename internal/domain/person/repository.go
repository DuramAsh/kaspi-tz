package person

import "context"

type Repository interface {
	InsertPerson(ctx context.Context, data Entity) (err error)
	GetPersonByIIN(ctx context.Context, iin string) (dest Entity, err error)
	GetPeopleByNamePart(ctx context.Context, name string) (dest []Entity, err error)
}
