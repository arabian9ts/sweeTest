package dto

type CreateStudentInputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type CreateStudentOutputForm struct {
	LastCreatedUserId int64
}

type CreateTaInputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type CreateTaOutputForm struct {
	LastCreatedUserId int64
}

type CreateTeacherInputForm struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type CreateTeacherOutputForm struct {
	LastCreatedUserId int64
}