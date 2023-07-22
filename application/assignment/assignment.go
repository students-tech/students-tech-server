package assignment

import (
	"context"
	"students-tech-server/shared"
	"students-tech-server/shared/dto"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
)

type (
	Service interface {
		CreateAssignment(data dto.CreateAssignmentRepositoryRequest) error
	}

	service struct {
		shared shared.Holder
	}
)

func (s *service) CreateAssignment(data dto.CreateAssignmentRepositoryRequest) error {
	query, args, _ := sq.Insert("students_project_assignment").
		Columns("students_id", "project_id").
		Values(data.StudentsID, data.ProjectID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	log.Debug().Str("query", query).Any("args", args).Send()

	_, err := s.shared.DB.Exec(context.Background(), query, args...)
	if err != nil {
		log.Error().Err(err).Msg("error inserting new bounty")
		return err
	}

	return nil
}

func NewService(shared shared.Holder) Service {
	return &service{
		shared: shared,
	}
}
