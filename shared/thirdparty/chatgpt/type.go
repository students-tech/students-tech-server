package chatgpt

import (
	"context"
)

// OpenAiProvider is the interface that provide methods to connect to OpenAI API.
//
// Behavior of GetTextCompletions should be receiving prompt string and a userID string as request identification.
// It will return the response to the answer, wrapped in TextCompletionResponse, and return any error if process failed.
//
// Behavior of GetTextModeration should be receiving an input string to be moderated.
// It will return the moderation response, wrapped in ModerationResponse, and return any error if process failed.
type ChatGPTProvider interface {
	GetTextCompletions(ctx context.Context, request GetTextCompletionRequest) (GetTextCompletionResponse, error)
}
