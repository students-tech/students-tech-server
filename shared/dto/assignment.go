package dto

type CreateAssignmentRequest struct {
	UserID    string `json:"userID" validate:"required"`
	ProjectID int    `json:"projectID" validate:"required"`
}

type CreateAssignmentRepositoryRequest struct {
	StudentsID int
	ProjectID  int
}
type CreateAssignmentResponse struct {
	Message string `json:"message" validate:"required"`
}
