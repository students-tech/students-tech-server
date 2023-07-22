package project

import (
	"context"
	"students-tech-server/application"
	"students-tech-server/shared"
	"students-tech-server/shared/constant"
	"students-tech-server/shared/dto"
	chatgpt2 "students-tech-server/shared/thirdparty/chatgpt"
)

type (
	ViewService interface {
		CreateProject(request dto.CreateProjectRequest) (dto.CreateProjectResponse, error)
		ListProject() (dto.GetListProjectResponse, error)
		GetDetailProject(ID uint32) (dto.GetListProjectData, error)
		CompleteRequirement(request dto.CompleteRequirementRequest) (resp dto.CompleteRequirementResponse, err error)
	}

	viewService struct {
		shared      shared.Holder
		chatGPT     chatgpt2.ChatGPTProvider
		application application.Holder
	}
)

func (v *viewService) GetDetailProject(ID uint32) (dto.GetListProjectData, error) {
	var (
		res dto.GetListProjectData
	)

	res, err := v.application.ProjectService.GetProjectByID(ID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (v *viewService) ListProject() (dto.GetListProjectResponse, error) {
	var (
		res dto.GetListProjectResponse
	)

	res, err := v.application.ProjectService.GetProjects()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (v *viewService) CreateProject(request dto.CreateProjectRequest) (dto.CreateProjectResponse, error) {
	var (
		res dto.CreateProjectResponse
	)

	err := v.application.ProjectService.InsertProject(request)
	if err != nil {
		return res, err
	}

	res = dto.CreateProjectResponse{
		Status: "Success",
	}
	return res, nil

}

func (v *viewService) CompleteRequirement(request dto.CompleteRequirementRequest) (resp dto.CompleteRequirementResponse, err error) {
	ctx := context.Background()
	gptResp, err := v.chatGPT.GetTextCompletions(ctx, chatgpt2.GetTextCompletionRequest{
		Model:    constant.DefaultModel,
		Messages: request.Messages,
	})

	resp.Message = gptResp.Choices[0].Message
	return resp, nil
}

func NewViewService(shared shared.Holder, gpt chatgpt2.ChatGPTProvider, application application.Holder) ViewService {
	return &viewService{
		shared:      shared,
		chatGPT:     gpt,
		application: application,
	}
}
