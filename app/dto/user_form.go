package dto

// Input
type CreateStudentInputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UpdateStudentInputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type CreateTaInputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UpdateTaInputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type CreateTeacherInputForm struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UpdateTeacherInputForm struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// Output
type GetStudentByIdOutputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
}

type GetTaByIdOutputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
}

type GetTeacherByIdOutputForm struct {
	FirstName string
	LastName  string
	Email     string
}

type CreateStudentOutputForm struct {
	LastInsertedId int64
}

type UpdateStudentOutputForm struct {
	Updated bool
}

type DeleteStudentOutputForm struct {
	Deleted bool
}

type CreateTaOutputForm struct {
	LastInsertedId int64
}

type UpdateTaOutputForm struct {
	Updated bool
}

type DeleteTaOutputForm struct {
	Deleted bool
}

type CreateTeacherOutputForm struct {
	LastInsertedId int64
}

type UpdateTeacherOutputForm struct {
	Updated bool
}

type DeleteTeacherOutputForm struct {
	Deleted bool
}
