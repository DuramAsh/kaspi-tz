package contragent

import (
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"

	"kaspi-tz/internal/domain/person"
	"kaspi-tz/pkg/log"
)

func (s *Service) CreatePerson(ctx context.Context, req person.CreatePersonRequest) (err error) {
	logger := log.LoggerFromContext(ctx).Named("CreatePerson")

	data := person.ParseToEntity(req)

	if err = s.personRepository.InsertPerson(ctx, data); err != nil {
		logger.Error("Failed to create person", zap.Error(err))
		return
	}

	return
}

func (s *Service) GetPersonByIIN(ctx context.Context, iin string) (dest person.GetPersonResponse, err error) {
	logger := log.LoggerFromContext(ctx).Named("GetPersonByIIN")

	personSrc, err := s.personRepository.GetPersonByIIN(ctx, iin)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			err = person.ErrPersonNotFound
		default:
			logger.Error("Failed to get person by iin", zap.Error(err))
		}

		return
	}

	dest = person.ParseFromEntity(personSrc)

	return
}

func (s *Service) GetPeopleByNamePart(ctx context.Context, name string) (dest []person.GetPersonResponse, err error) {
	logger := log.LoggerFromContext(ctx).Named("GetPersonByIIN")

	peopleSrc, err := s.personRepository.GetPeopleByNamePart(ctx, name)
	if err != nil {
		logger.Error("Failed to get person by iin", zap.Error(err))
		return
	}

	dest = person.ParseFromEntities(peopleSrc)

	return
}
