package contragent

import (
	"kaspi-tz/internal/domain/person"
)

// Configuration is an alias for a function that will take in a pointer to a Service and modify it
type Configuration func(s *Service) error

// Service is an implementation of the Service
type Service struct {
	personRepository person.Repository
}

// New takes a variable amount of Configuration functions and returns a new Service
// Each Configuration will be called in the order they are passed in
func New(configs ...Configuration) (s *Service, err error) {
	// Insert the service
	s = &Service{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

func WithPersonRepository(personRepository person.Repository) Configuration {
	return func(s *Service) error {
		s.personRepository = personRepository
		return nil
	}
}
