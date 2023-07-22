package project

import (
	"context"
	"students-tech-server/shared"
	"students-tech-server/shared/constant"
	"students-tech-server/shared/dto"
	chatgpt2 "students-tech-server/shared/thirdparty/chatgpt"
)

type (
	ViewService interface {
		CreateProject(request dto.CreateProjectRequest) error
		CompleteRequirement(request dto.CompleteRequirementRequest) (resp dto.CompleteRequirementResponse, err error)
	}

	viewService struct {
		shared  shared.Holder
		chatGPT chatgpt2.ChatGPTProvider
	}
)

func (v *viewService) CreateProject(request dto.CreateProjectRequest) error {

	return nil
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

func NewViewService(shared shared.Holder, gpt chatgpt2.ChatGPTProvider) ViewService {
	return &viewService{
		shared:  shared,
		chatGPT: gpt,
	}
}
