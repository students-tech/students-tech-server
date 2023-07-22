package dto

type SendEmailRequest struct {
	Name        string           `json:"name"`
	Email       string           `json:"email"`
	CallbackURL string           `json:"callbackURL"`
	Project     SendEmailProject `json:"project"`
}

type SendEmailProject struct {
	ProjectName     string `json:"projectName"`
	ProjectStack    string `json:"projectStack"`
	ProjectOverview string `json:"projectOverview"`
	ProjectPlatform string `json:"projectPlatform"`
	IsHaveDesign    bool   `json:"isHaveDesign"`
}

type CreateProjectRequest struct {
	EmailOwner            string `json:"emailOwner"`
	PhoneOwner            string `json:"phoneOwner"`
	RawProjectDescription string `json:"rawProjectDescription"`
	ProjectName           string `json:"projectName"`
	ProjectStack          string `json:"projectStack"`
	ProjectPlatform       string `json:"projectPlatform"`
	ProjectObjective      string `json:"projectObjective"`
	IsHaveDesign          string `json:"isHaveDesign"`
	BudgetRange           string `json:"budgetRange"`
	TimelineRange         string `json:"timelineRange"`
}

type GetListProjectResponse struct {
	Data []GetListProjectData `json:"data"`
}
type GetListProjectData struct {
	ID          int    `json:"id"`
	EmailOwner  string `json:"emailOwner"`
	PhoneOwner  string `json:"phoneOwner"`
	ProjectName string `json:"projectName"`

	RawProjectDescription string `json:"rawProjectDescription"`
	ProjectStack          string `json:"projectStack"`
	ProjectPlatform       string `json:"projectPlatform"`
	ProjectObjective      string `json:"projectObjective"`
	IsHaveDesign          bool   `json:"isHaveDesign"`
	BudgetRange           int    `json:"budgetRange"`
	TimelineRange         string `json:"timelineRange"`
}

type GetProjectDetail struct {
	ProjectID uint32 `json:"projectID"`
}

type CreateProjectResponse struct {
	Status string `json:"status"`
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
