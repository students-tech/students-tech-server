package dto

type (
	// CreateStudentsRequest CreateStudentsRequest
	CreateStudentsRequest struct {
		UserID             string `json:"userID" validate:"required"`
		Email              string `json:"email" validate:"required,email"`
		Name               string `json:"name" validate:"required"`
		Residency          string `json:"residency" validate:"required"`
		PhoneNumber        string `json:"phoneNumber" validate:"required"`
		LinkedInUrl        string `json:"linkedInUrl" validate:"omitempty"`
		GithubUsername     string `json:"githubUsername" validate:"omitempty"`
		GitlabUsername     string `json:"gitlabUsername" validate:"omitempty"`
		ProjectCount       string `json:"projectCount" validate:"omitempty"`
		UniversityName     string `json:"universityName" validate:"required"`
		Major              string `json:"major" validate:"required"`
		CurrentSemester    int    `json:"currentSemester" validate:"required"`
		RelevantCoursework string `json:"relevantCoursework" validate:"omitempty"`
		HoursAvailability  string `json:"hoursAvailability" validate:"required"`
		Role               string `json:"role" validate:"required"`
		UniversityEmail    string `json:"universityEmail,omitempty"`
		KTMUrl             string `json:"ktmUrl" validate:"required"`
		ResumeUrl          string `json:"resumeUrl" validate:"required"`
	}

	// CreateStudentsResponse CreateStudentsResponse
	CreateStudentsResponse struct {
		Status string `json:"status"`
	}

	Students struct {
		ID    int
		Name  string
		Email string
	}
)
