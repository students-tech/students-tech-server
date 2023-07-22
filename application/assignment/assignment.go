package assignment

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
		CreateAssignment(data dto.CreateAssignmentRepositoryRequest) error
		GetAssignmentQuestionAndAnswer(assignmentID uint32) (resp dto.GetAssignmentQuestionAndAnswerResponse, err error)
	}

	service struct {
		shared shared.Holder
	}
)

func (s *service) GetAssignmentQuestionAndAnswer(assignmentID uint32) (resp dto.GetAssignmentQuestionAndAnswerResponse, err error) {
	var data dto.GetAssignmentQuestionAndAnswerResponse
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Question)

	query := psql.Select("q.id, q.question, q.question_number, a.answer , a.is_correct").
		From("question as q").Join("answer as a on a.question_id = q.id").
		Where(fmt.Sprintf("q.students_project_assignment_id = %d", assignmentID))

	sql, args, err := query.ToSql()
	if err != nil {
		log.Error().Err(err).Msg("error make query")
		return data, err
	}

	rows, err := s.shared.DB.Query(context.Background(), sql, args...)
	if err != nil {
		log.Error().Err(err).Msg("error query row")
		return data, err
	}
	defer rows.Close()

	//var datas []dto.QuestionAndAnswerData
	//for rows.Next() {
	//	var data dto.GetListProjectData
	//	if err := rows.Scan(&data.ID, &data.EmailOwner, &data.PhoneOwner, &data.ProjectName, &data.RawProjectDescription, &data.ProjectPlatform, &data.ProjectStack, &data.ProjectObjective, &data.IsHaveDesign, &data.BudgetRange, &data.TimelineRange); err != nil {
	//		log.Error().Err(err).Msg("error scan row")
	//		return res, err
	//	}
	//	datas = append(datas, data)
	//}
	//
	//res.Data = datas
	//fmt.Println(row)
	//if err := row.Scan(&data.ID, &data.EmailOwner, &data.PhoneOwner, &data.ProjectName, &data.RawProjectDescription, &data.ProjectPlatform, &data.ProjectStack, &data.ProjectObjective, &data.IsHaveDesign, &data.BudgetRange, &data.TimelineRange); err != nil {
	//	log.Error().Err(err).Msg("error scan row")
	//	return data, err
	//}

	return data, err

}

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
