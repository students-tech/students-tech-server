package students

import (
	"context"
	"fmt"
	"students-tech-server/shared"
	"students-tech-server/shared/dto"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
)

type (
	Service interface {
		InsertStudents(data dto.CreateStudentsRequest) error
		GetStudentsByUserID(userID string) (resp dto.Students, err error)
	}

	service struct {
		shared shared.Holder
	}
)

func (s *service) GetStudentsByUserID(userID string) (resp dto.Students, err error) {
	var data dto.Students
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Question)

	query := psql.Select("s.id, s.name, s.email").
		From("students as s").
		Where(fmt.Sprintf("s.user_id = '%s'", userID))

	sql, args, err := query.ToSql()
	if err != nil {
		log.Error().Err(err).Msg("error make query")
		return data, err
	}

	row := s.shared.DB.QueryRow(context.Background(), sql, args...)

	if err := row.Scan(&data.ID, &data.Name, &data.Email); err != nil {
		log.Error().Err(err).Msg("error scan row")
		return data, err
	}

	return data, nil
}

func (s *service) InsertStudents(data dto.CreateStudentsRequest) error {
	query, args, _ := sq.Insert("students").
		Columns("user_id", "email", "name", "residency", "phone_number", "linkedin_url", "github_username", "gitlab_username", "project_count", "university_name", "major", "current_semester", "relevant_coursework", "hours_availability_per_week", "role", "university_email", "ktm_url", "resume_url").
		Values(data.UserID, data.Email, data.Name, data.Residency, data.PhoneNumber, data.LinkedInUrl, data.GithubUsername, data.GitlabUsername, data.ProjectCount, data.UniversityName, data.Major, data.CurrentSemester, data.RelevantCoursework, data.HoursAvailability, data.Role, data.UniversityEmail, data.KTMUrl, data.ResumeUrl).
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
