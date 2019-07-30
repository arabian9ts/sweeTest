package dto

// Input
type AuthorizeStudentInputForm struct {
	StudentNo string `validate:"required,lte=20" json:"student_no"`
	Password  string `validate:"required,lte=100" json:"password"`
}

type AuthorizeAssistantInputForm struct {
	StudentNo string `validate:"required,lte=20" json:"student_no"`
	Password  string `validate:"required,lte=100" json:"password"`
}

type AuthorizeTeacherInputForm struct {
	Email    string `validate:"required,lte=100,email" json:"email"`
	Password string `validate:"required,lte=100" json:"password"`
}

type AuthorizeAdminInputForm struct {
	Email    string `validate:"required,lte=100,email" json:"email"`
	Password string `validate:"required,lte=100" json:"password"`
}

// Output
type AuthorizeStudentOutputForm struct {
	JwtToken string `json:"jwt_token"`
}

type AuthorizeAssistantOutputForm struct {
	JwtToken string `json:"jwt_token"`
}

type AuthorizeTeacherOutputForm struct {
	JwtToken string `json:"jwt_token"`
}

type AuthorizeAdminOutputForm struct {
	JwtToken string `json:"jwt_token"`
}
