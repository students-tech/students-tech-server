package chatgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"students-tech-server/shared/config"
	"time"
)

// ChatGPTService is the struct handling all requests to OpenAI API
type ChatGPTService struct {
	host   string
	apiKey string
	client http.Client
}

// NewOpenAiService returns a new ChatGPTService with the specified configurations
func NewOpenAiService(env *config.Env) ChatGPTProvider {

	// set default timeout to be 30s
	clientTimeout := time.Duration(30) * time.Second

	httpClient := http.Client{
		Timeout: clientTimeout,
	}

	return &ChatGPTService{
		host:   env.ChatGPTHost,
		apiKey: env.ChatGPTAPIKey,
		client: httpClient,
	}
}

func (o *ChatGPTService) GetTextCompletions(ctx context.Context, request GetTextCompletionRequest) (GetTextCompletionResponse, error) {
	var responseBody GetTextCompletionResponse
	requestUrl := fmt.Sprintf("%s%s", o.host, "/v1/chat/completions")

	requestBody, _ := json.Marshal(request)
	bodyBytes := bytes.NewBuffer(requestBody)

	req, _ := http.NewRequestWithContext(ctx, "POST", requestUrl, bodyBytes)

	err := o.doRequest(ctx, req, &responseBody)
	if err != nil {
		return responseBody, err
	}

	return responseBody, nil
}

func (o *ChatGPTService) doRequest(ctx context.Context, req *http.Request, response interface{}) error {
	log.Debug().Msg("[ChatGPTService] starting request to OpenAI")

	o.initializeRequestHeader(req)

	resp, err := o.client.Do(req)
	if err != nil {
		log.Debug().Msg("[ChatGPTService] Failed request to OpenAI")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// read the whole response
		errDetail, _ := io.ReadAll(resp.Body)
		log.Debug().Msgf("[ChatGPTService] Request to OpenAI returns error %s", string(errDetail))

		var returnedErr error
		switch resp.StatusCode {
		case http.StatusBadRequest:
			returnedErr = ErrorInvalidParameter
		case http.StatusServiceUnavailable:
			returnedErr = ErrorServiceUnavailable
		case http.StatusUnauthorized:
			returnedErr = ErrorAPIKeyInvalid
		case http.StatusTooManyRequests:
			returnedErr = ErrorTooManyRequest
		default:
			returnedErr = ErrorUnknown
		}

		return returnedErr
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Debug().Msg("[ChatGPTService] Invalid response from OpenAI")
		return ErrorInvalidResponseStructure
	}

	return nil
}

func (o *ChatGPTService) initializeRequestHeader(req *http.Request) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", o.apiKey),
		"Content-Type":  "application/json",
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}
}
