package project

import (
	"students-tech-server/shared"
	"students-tech-server/shared/dto"
)

type (
	ViewService interface {
		SendEmail(request dto.SendEmailRequest) error
	}

	viewService struct {
		shared shared.Holder
	}
)

func NewViewService(shared shared.Holder) ViewService {
	return &viewService{
		shared: shared,
	}
}
