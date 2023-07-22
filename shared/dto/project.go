package dto

type SendEmailRequest struct {
	Name        string           `json:"name"`
	Email       string           `json:"email"`
	CallbackURL string           `json:"callback_url"`
	Project     SendEmailProject `json:"project"`
}

type SendEmailProject struct {
	ProjectName     string `json:"project_name"`
	ProjectStack    string `json:"project_stack"`
	ProjectOverview string `json:"project_overview"`
	ProjectPlatform string `json:"project_platform"`
	IsHaveDesign    bool   `json:"is_have_design"`
}

type CreateProjectRequest struct {
	EmailOwner            string `json:"email_owner"`
	PhoneOwner            string `json:"phone_owner"`
	RawProjectDescription string `json:"raw_project_description"`
	ProjectName           string `json:"project_name"`
	ProjectStack          string `json:"project_stack"`
	ProjectPlatform       string `json:"project_platform"`
	IsHaveDesign          bool   `json:"is_have_design"`
	BudgetRange           int    `json:"budget_range"`
	TimelineRange         string `json:"timeline_range"`
}

type CompleteRequirementMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompleteRequirementRequest struct {
	Messages []CompleteRequirementMessage `json:"messages"`
}

type CompleteRequirementResponse struct {
	Message CompleteRequirementMessage `json:"message"`
}
