package chatgpt

import "students-tech-server/shared/dto"

type GetTextCompletionRequest struct {
	Model    string                           `json:"model"`
	Messages []dto.CompleteRequirementMessage `json:"messages"`
}

type GetTextCompletionResponse struct {
	Id      string                    `json:"id"`
	Object  string                    `json:"object"`
	Created int                       `json:"created"`
	Model   string                    `json:"model"`
	Choices []GetTextCompletionChoice `json:"choices"`
	Usage   GetTextCompletionUsage    `json:"usage"`
}

type GetTextCompletionUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type GetTextCompletionChoice struct {
	Index        int                            `json:"index"`
	Message      dto.CompleteRequirementMessage `json:"message"`
	FinishReason string                         `json:"finish_reason"`
}
