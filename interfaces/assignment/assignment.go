package assignment

import (
	"students-tech-server/application"
	"students-tech-server/shared"
	"students-tech-server/shared/dto"
)

type (
	ViewService interface {
		CreateAssignment(req dto.CreateAssignmentRequest) (dto.CreateAssignmentResponse, error)
	}

	viewService struct {
		shared      shared.Holder
		application application.Holder
	}
)

func (v *viewService) CreateAssignment(req dto.CreateAssignmentRequest) (dto.CreateAssignmentResponse, error) {
	var (
		res dto.CreateAssignmentResponse
	)
	students, err := v.application.StudentsService.GetStudentsByUserID(req.UserID)
	if err != nil {
		return res, err
	}

	err = v.application.AssignmentService.CreateAssignment(dto.CreateAssignmentRepositoryRequest{
		StudentsID: students.ID,
		ProjectID:  req.ProjectID,
	})
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewViewService(shared shared.Holder, application application.Holder) ViewService {
	return &viewService{
		shared:      shared,
		application: application,
	}
}
