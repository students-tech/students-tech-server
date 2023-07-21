package health

import (
	"students-tech-server/shared"
	"students-tech-server/shared/dto"
)

type (
	ViewService interface {
		SystemHealth() ([]dto.Status, error)
	}

	viewService struct {
		shared shared.Holder
	}
)

func (v *viewService) SystemHealth() ([]dto.Status, error) {
	status := make([]dto.Status, 0)

	status = append(status, dto.Status{
		Name:   "Http",
		Status: dto.OK,
		Data:   v.shared.Http.HandlersCount(),
	})

	return status, nil
}

func NewViewService(shared shared.Holder) ViewService {
	return &viewService{
		shared: shared,
	}
}
