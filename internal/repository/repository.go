package repository

import (
	"kaspi-tz/internal/repository/postgres"
	"kaspi-tz/pkg/store"
)

// Configuration is an alias for a function that will take in a pointer to a Repository and modify it
type Configuration func(r *Repository) error

// Repository is an implementation of the Repository
type Repository struct {
	Person *postgres.PersonRepository
}

// New takes a variable amount of Configuration functions and returns a new Repository
// Each Configuration will be called in the order they are passed in
func New(configs ...Configuration) (s *Repository, err error) {
	// Insert the repository
	s = &Repository{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the repository into the configuration function
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

func (r *Repository) Close() (err error) {
	if r.Person != nil && r.Person.DB != nil {
		return r.Person.DB.Close()
	}

	return
}

func WithPostgresStore(dsn string) Configuration {
	return func(r *Repository) (err error) {
		db, err := store.NewSQL(dsn)
		if err != nil {
			return
		}

		if err = store.Migrate(dsn); err != nil {
			return
		}

		r.Person = postgres.NewPersonRepository(db)

		return
	}
}
