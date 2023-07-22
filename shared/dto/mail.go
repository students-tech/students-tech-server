package dto

const (
	USER    EmailMessageType = "USER"
	PROJECT EmailMessageType = "PROJECT"
)

type (
	EmailMessageType string

	EmailMessage struct {
		Type            EmailMessageType `json:"type"`
		Name            string           `json:"name"`
		Email           string           `json:"email"`
		CallbackUrl     string           `json:"callbackUrl"`
		ProjectName     string           `json:"projectName"`
		ProjectStack    string           `json:"projectStack"`
		ProjectPlatform string           `json:"projectPlatform"`
		IsHaveDesign    bool             `json:"isHaveDesign"`
		ProjectOverview string           `json:"projectOverview"`
	}
)
