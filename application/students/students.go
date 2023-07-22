package students

import (
	"context"
	"students-tech-server/shared"
	"students-tech-server/shared/dto"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
)

type (
	Service interface {
		InsertStudents(data dto.CreateStudentsRequest) error
	}

	service struct {
		shared shared.Holder
	}
	Service interface {
		InsertStudents(data dto.CreateStudentsRequest) error
	}

	service struct {
		shared shared.Holder
	}
)

func (s *service) InsertStudents(data dto.CreateStudentsRequest) error {
	query, args, _ := sq.Insert("students").
		Columns("email", "name", "residency", "phone_number", "linkedin_url", "github_username", "gitlab_username", "project_count", "university_name", "major", "current_semester", "relevant_coursework", "hours_availability_per_week", "role", "university_email", "ktm_url", "resume_url").
		Values(data.Email, data.Name, data.Residency, data.PhoneNumber, data.LinkedInUrl, data.GithubUsername, data.GitlabUsername, data.ProjectCount, data.UniversityName, data.Major, data.CurrentSemester, data.RelevantCoursework, data.HoursAvailability, data.Role, data.UniversityEmail, data.KTMUrl, data.ResumeUrl).
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
