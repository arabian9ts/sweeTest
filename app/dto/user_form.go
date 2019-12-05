package dto

import "time"

// Input
type CreateStudentInputForm struct {
	StudentNo string `json:"student_no" validate:"required,lte=20"`
	FirstName string `json:"first_name" validate:"required,lte=15"`
	LastName  string `json:"last_name"  validate:"required,lte=15"`
	Email     string `json:"email"      validate:"required,lte=100,email"`
	Password  string `json:"password"   validate:"required,lte=100"`
}

type CreateAssistantInputForm struct {
	StudentNo string `json:"student_no" validate:"required,lte=20"`
	FirstName string `json:"first_name" validate:"required,lte=15"`
	LastName  string `json:"last_name"  validate:"required,lte=15"`
	Email     string `json:"email"      validate:"required,lte=100,email"`
	Password  string `json:"password"   validate:"required,lte=100"`
}

type CreateTeacherInputForm struct {
	FirstName string `json:"first_name" validate:"required,lte=15"`
	LastName  string `json:"last_name"  validate:"required,lte=15"`
	Email     string `json:"email"      validate:"required,lte=100,email"`
	Password  string `json:"password"   validate:"required,lte=100"`
}

type UpdateStudentInputForm struct {
	ID        int64  `json:"-"          validate:"required,gt=0"`
	StudentNo string `json:"student_no" validate:"required,lte=20"`
	FirstName string `json:"first_name" validate:"required,lte=15"`
	LastName  string `json:"last_name"  validate:"required,lte=15"`
	Email     string `json:"email"      validate:"required,lte=100,email"`
	Password  string `json:"password"   validate:"required,lte=100"`
}

type UpdateAssistantInputForm struct {
	ID        int64  `json:"-"          validate:"required,gt=0"`
	StudentNo string `json:"student_no" validate:"required,lte=20"`
	FirstName string `json:"first_name" validate:"required,lte=15"`
	LastName  string `json:"last_name"  validate:"required,lte=15"`
	Email     string `json:"email"      validate:"required,lte=100,email"`
	Password  string `json:"password"   validate:"required,lte=100"`
}

type UpdateTeacherInputForm struct {
	ID        int64  `json:"-"          validate:"required,gt=0"`
	FirstName string `json:"first_name" validate:"required,lte=15"`
	LastName  string `json:"last_name"  validate:"required,lte=15"`
	Email     string `json:"email"      validate:"required,lte=100,email"`
}

// Output
type GetStudentsOutputForm []*GetStudentOutputForm

type GetAssistantsOutputForm []*GetAssistantOutputForm

type GetTeachersOutputForm []*GetTeacherOutputForm

type GetStudentOutputForm struct {
	ID        int64     `json:"id"`
	StudentNo string    `json:"student_no"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAssistantOutputForm struct {
	ID        int64     `json:"id"`
	StudentNo string    `json:"student_no"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetTeacherOutputForm struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateStudentOutputForm struct {
	ID        int64     `json:"id"`
	StudentNo string    `json:"student_no" validate:"required,lte=20"`
	FirstName string    `json:"first_name" validate:"required,lte=15"`
	LastName  string    `json:"last_name"  validate:"required,lte=15"`
	Email     string    `json:"email"      validate:"required,lte=100,email"`
	Password  string    `json:"password"   validate:"required,lte=100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAssistantOutputForm struct {
	ID        int64     `json:"id"`
	StudentNo string    `json:"student_no" validate:"required,lte=20"`
	FirstName string    `json:"first_name" validate:"required,lte=15"`
	LastName  string    `json:"last_name"  validate:"required,lte=15"`
	Email     string    `json:"email"      validate:"required,lte=100,email"`
	Password  string    `json:"password"   validate:"required,lte=100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTeacherOutputForm struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateStudentOutputForm struct {
	ID        int64     `json:"id"`
	StudentNo string    `json:"student_no" validate:"required,lte=20"`
	FirstName string    `json:"first_name" validate:"required,lte=15"`
	LastName  string    `json:"last_name"  validate:"required,lte=15"`
	Email     string    `json:"email"      validate:"required,lte=100,email"`
	Password  string    `json:"password"   validate:"required,lte=100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateAssistantOutputForm struct {
	ID        int64     `json:"id"`
	StudentNo string    `json:"student_no" validate:"required,lte=20"`
	FirstName string    `json:"first_name" validate:"required,lte=15"`
	LastName  string    `json:"last_name"  validate:"required,lte=15"`
	Email     string    `json:"email"      validate:"required,lte=100,email"`
	Password  string    `json:"password"   validate:"required,lte=100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateTeacherOutputForm struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name" validate:"required,lte=15"`
	LastName  string    `json:"last_name"  validate:"required,lte=15"`
	Email     string    `json:"email"      validate:"required,lte=100,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteStudentOutputForm struct {
	AffectedRowsCount int64 `json:"count"`
}

type DeleteAssistantOutputForm struct {
	AffectedRowsCount int64 `json:"count"`
}

type DeleteTeacherOutputForm struct {
	AffectedRowsCount int64 `json:"count"`
}
