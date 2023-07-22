package dto

type (
	// CreateStudentsRequest CreateStudentsRequest
	CreateStudentsRequest struct {
		Email             string `json:"email" validate:"required,email"`
		Name              string `json:"name" validate:"required"`
		Residency         string `json:"residency" validate:"required"`
		PhoneNumber       string `json:"phoneNumber" validate:"required"`
		LinkedInUrl       string `json:"linkedInUrl" validate:"omitempty"`
		GithubUsername    string `json:"githubUsername" validate:"omitempty"`
		GitlabUsername    string `json:"gitlabUsername" validate:"omitempty"`
		ProjectCount      int    `json:"projectCount" validate:"omitempty"`
		UniveristyName    string `json:"universityName" validate:"required"`
		Major             string `json:"major" validate:"required"`
		CurrentSemester   int    `json:"currentSemester" validate:"required"`
		RelevantCousework string `json:"relevantCoursework" validate:"omitempty"`
		HoursAvailibility string `json:"hoursAvailability" validate:"required"`
		Role              string `json:"role" validate:"required"`
		UniversityEmail   string `json:"universityEmail" validate:"required"`
		KTMUrl            string `json:"ktmUrl" validate:"required"`
		ResumeUrl         string `json:"resumeUrl" validate:"required"`
	}
)
