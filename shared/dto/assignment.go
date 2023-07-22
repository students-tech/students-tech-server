package dto

type CreateAssignmentRequest struct {
	UserID    string `json:"userID" validate:"required"`
	ProjectID int    `json:"projectID" validate:"required"`
}

type GetAssignmentRequest struct {
	AssignmentID uint32 `json:"assignmentID" validate:"required"`
}

type CreateAssignmentRepositoryRequest struct {
	StudentsID int
	ProjectID  int
}
type CreateAssignmentResponse struct {
	Message string `json:"message" validate:"required"`
}

type GetAssignmentQuestionAndAnswerResponse struct {
	Data []QuestionAndAnswerData `json:"data"`
}

type QuestionAndAnswerData struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}
type Answer struct {
	Text   string `json:"text"`
	IsTrue bool   `json:"isTrue"`
}

type QuestionQueryJoin struct {
	QuestionID     uint32
	QuestionText   string
	QuestionNumber uint32
	Answer         string
	IsTrue         bool
}
