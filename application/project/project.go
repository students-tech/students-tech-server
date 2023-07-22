package project

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"students-tech-server/shared"
	"students-tech-server/shared/dto"

	"github.com/rs/zerolog/log"
)

type (
	Service interface {
		InsertProject(data dto.CreateProjectRequest) error
		GetProjects() (dto.GetListProjectResponse, error)
		GetProjectByID(ID uint32) (dto.GetListProjectData, error)
	}

	service struct {
		shared shared.Holder
	}
)

func (s *service) GetProjectByID(ID uint32) (dto.GetListProjectData, error) {
	var data dto.GetListProjectData
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Question)

	query := psql.Select("p.id, p.email_owner, p.phone_owner, p.project_name, p.raw_project_description, p.platform, p.tech_stack, p.project_objective, p.is_have_design, p.budget_range, p.timeline").
		From("project as p").
		Where(fmt.Sprintf("p.id = %d", ID))

	sql, args, err := query.ToSql()
	if err != nil {
		log.Error().Err(err).Msg("error make query")
		return data, err
	}

	row := s.shared.DB.QueryRow(context.Background(), sql, args...)

	if err := row.Scan(&data.ID, &data.EmailOwner, &data.PhoneOwner, &data.ProjectName, &data.RawProjectDescription, &data.ProjectPlatform, &data.ProjectStack, &data.ProjectObjective, &data.IsHaveDesign, &data.BudgetRange, &data.TimelineRange); err != nil {
		log.Error().Err(err).Msg("error scan row")
		return data, err
	}

	return data, nil
}

// import sq "github.com/Masterminds/squirrel"
//
// users := sq.Select("*").From("users").Join("emails USING (email_id)")
//
// active := users.Where(sq.Eq{"deleted_at": nil})
//
// sql, args, err := active.ToSql()
//
// sql == "SELECT * FROM users JOIN emails USING (email_id) WHERE deleted_at IS NULL"
func (s *service) GetProjects() (dto.GetListProjectResponse, error) {
	var res dto.GetListProjectResponse
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Question)

	query := psql.Select("p.id, p.email_owner, p.phone_owner, p.project_name, p.raw_project_description, p.platform, p.tech_stack, p.project_objective, p.is_have_design, p.budget_range, p.timeline").From("project as p")

	sql, args, err := query.ToSql()
	if err != nil {
		log.Error().Err(err).Msg("error make query")
		return res, err
	}

	rows, err := s.shared.DB.Query(context.Background(), sql, args...)
	if err != nil {
		log.Error().Err(err).Msg("error query row")
		return res, err
	}
	defer rows.Close()

	var datas []dto.GetListProjectData
	for rows.Next() {
		var data dto.GetListProjectData
		if err := rows.Scan(&data.ID, &data.EmailOwner, &data.PhoneOwner, &data.ProjectName, &data.RawProjectDescription, &data.ProjectPlatform, &data.ProjectStack, &data.ProjectObjective, &data.IsHaveDesign, &data.BudgetRange, &data.TimelineRange); err != nil {
			log.Error().Err(err).Msg("error scan row")
			return res, err
		}
		datas = append(datas, data)
	}

	res.Data = datas
	return res, nil
}

func (s *service) InsertProject(data dto.CreateProjectRequest) error {
	query, args, _ := sq.Insert("project").
		Columns("email_owner", "phone_owner", "project_name", "raw_project_description", "platform", "tech_stack", "project_objective", "is_have_design", "budget_range", "timeline", "is_active").
		Values(data.EmailOwner, data.PhoneOwner, data.ProjectName, data.RawProjectDescription, data.ProjectPlatform, data.ProjectStack, data.ProjectObjective, data.IsHaveDesign, data.BudgetRange, data.TimelineRange, false).
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
