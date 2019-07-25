package dto

// Input
type CreateStudentInputForm struct {
	StudentNo string `validate:"required,lte=20" json:"student_no"`
	FirstName string `validate:"required,lte=15" json:"first_name"`
	LastName  string `validate:"required,lte=15" json:"last_name"`
	Email     string `validate:"required,lte=100,email" json:"email"`
	Password  string `validate:"required,lte=100" json:"password"`
}

type UpdateStudentInputForm struct {
	ID        int64  `validate:"required,gt=0"`
	StudentNo string `validate:"required,lte=20" json:"student_no"`
	FirstName string `validate:"required,lte=15" json:"first_name"`
	LastName  string `validate:"required,lte=15" json:"last_name"`
	Email     string `validate:"required,lte=100,email" json:"email"`
	Password  string `validate:"required,lte=100" json:"password"`
}

type CreateTaInputForm struct {
	StudentNo string `validate:"required,lte=20" json:"student_no"`
	FirstName string `validate:"required,lte=15" json:"first_name"`
	LastName  string `validate:"required,lte=15" json:"last_name"`
	Email     string `validate:"required,lte=100,email" json:"email"`
	Password  string `validate:"required,lte=100" json:"password"`
}

type UpdateTaInputForm struct {
	ID        int64  `validate:"required,gt=0"`
	StudentNo string `validate:"required,lte=20" json:"student_no"`
	FirstName string `validate:"required,lte=15" json:"first_name"`
	LastName  string `validate:"required,lte=15" json:"last_name"`
	Email     string `validate:"required,lte=100,email" json:"email"`
	Password  string `validate:"required,lte=100" json:"password"`
}

type CreateTeacherInputForm struct {
	FirstName string `validate:"required,lte=15" json:"first_name"`
	LastName  string `validate:"required,lte=15" json:"last_name"`
	Email     string `validate:"required,lte=100,email" json:"email"`
	Password  string `validate:"required,lte=100" json:"password"`
}

type UpdateTeacherInputForm struct {
	ID        int64  `validate:"required,gt=0"`
	FirstName string `validate:"required,lte=15" json:"first_name"`
	LastName  string `validate:"required,lte=15" json:"last_name"`
	Email     string `validate:"required,lte=100,email" json:"email"`
	Password  string `validate:"required,lte=100" json:"password"`
}

// Output
type GetStudentByIdOutputForm struct {
	StudentNo string `json:"student_no"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type GetTaByIdOutputForm struct {
	StudentNo string `json:"student_no"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type GetTeacherByIdOutputForm struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreateStudentOutputForm struct {
	LastInsertedId int64 `json:"inserted_id"`
}

type UpdateStudentOutputForm struct {
	Updated bool `json:"updated"`
}

type DeleteStudentOutputForm struct {
	Deleted bool `json:"deleted"`
}

type CreateTaOutputForm struct {
	LastInsertedId int64 `json:"inserted_id"`
}

type UpdateTaOutputForm struct {
	Updated bool `json:"updated"`
}

type DeleteTaOutputForm struct {
	Deleted bool `json:"deleted"`
}

type CreateTeacherOutputForm struct {
	LastInsertedId int64 `json:"inserted_id"`
}

type UpdateTeacherOutputForm struct {
	Updated bool `json:"updated"`
}

type DeleteTeacherOutputForm struct {
	Deleted bool `json:"deleted"`
}
