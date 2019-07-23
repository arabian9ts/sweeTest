package dto

type ReadStudentOutputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
}

type ReadTaOutputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
}

type ReadTeacherOutputForm struct {
	FirstName string
	LastName  string
	Email     string
}

type CreateStudentInputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type WriteStudentOutputForm struct {
	LastChangedUserId int64
}

type CreateTaInputForm struct {
	StudentNo string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type WriteTaOutputForm struct {
	LastChangedUserId int64
}

type CreateTeacherInputForm struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type WriteTeacherOutputForm struct {
	LastChangedUserId int64
}
