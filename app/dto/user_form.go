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

type CreateAssistantInputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UpdateAssistantInputForm struct {
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

type GetAssistantByIdOutputForm struct {
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

type CreateAssistantOutputForm struct {
	LastInsertedId int64
}

type UpdateAssistantOutputForm struct {
	Updated bool
}

type DeleteAssistantOutputForm struct {
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
