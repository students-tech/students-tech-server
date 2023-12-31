package students

import (
	"students-tech-server/application"
	"students-tech-server/shared"
	"students-tech-server/shared/common"
	"students-tech-server/shared/dto"
)

type (
	ViewService interface {
		RegisterStudents(req dto.CreateStudentsRequest) (dto.CreateStudentsResponse, error)
	}

	viewService struct {
		shared      shared.Holder
		application application.Holder
	}
)

func (v *viewService) RegisterStudents(req dto.CreateStudentsRequest) (dto.CreateStudentsResponse, error) {
	var (
		res dto.CreateStudentsResponse
	)

	err := v.application.StudentsService.InsertStudents(req)
	if err != nil {
		return res, err
	}

	res = dto.CreateStudentsResponse{
		Status: "Success",
	}

	if req.UniversityEmail != "" {
		go common.PublishEmailTopic(dto.EmailMessage{
			Type:  dto.USER,
			Name:  req.Name,
			Email: req.UniversityEmail,
		}, v.shared.Env.GCPProjectId, v.shared.Env.GCPPubSubTopic)
	}

	return res, nil
}

func NewViewService(shared shared.Holder, apapplication application.Holder) ViewService {
	return &viewService{
		shared:      shared,
		application: apapplication,
	}
}
