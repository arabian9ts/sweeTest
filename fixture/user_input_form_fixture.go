package fixture

import (
	"github.com/arabian9ts/sweeTest/app/dto"
)

func NewValidCreateStudentInputForm() *dto.CreateStudentInputForm {
	return &dto.CreateStudentInputForm{
		StudentNo: "12345678",
		FirstName: "First",
		LastName:  "Last",
		Email:     "Email@test.com",
	}
}

func NewValidCreateAssistantInputForm() *dto.CreateAssistantInputForm {
	return &dto.CreateAssistantInputForm{
		StudentNo: "121312312",
		FirstName: "taku",
		LastName:  "maeda",
		Email:     "aaa@bbb.ccc",
		Password:  "password",
	}
}

func NewValidCreateTeacherInputForm() *dto.CreateTeacherInputForm {
	return &dto.CreateTeacherInputForm{
		FirstName: "taku",
		LastName:  "maeda",
		Email:     "aaa@bbb.ccc",
		Password:  "password",
	}
}

func NewValidUpdateStudentInputForm() *dto.UpdateStudentInputForm {
	return &dto.UpdateStudentInputForm{
		ID:        1,
		StudentNo: "121312312",
		FirstName: "taku",
		LastName:  "maeda",
		Email:     "aaa@bbb.ccc",
	}
}

func NewValidUpdateAssistantInputForm() *dto.UpdateAssistantInputForm {
	return &dto.UpdateAssistantInputForm{
		ID:        1,
		StudentNo: "121312312",
		FirstName: "taku",
		LastName:  "maeda",
		Email:     "aaa@bbb.ccc",
	}
}

func NewValidUpdateTeacherInputForm() *dto.UpdateTeacherInputForm {
	return &dto.UpdateTeacherInputForm{
		ID:        1,
		FirstName: "taku",
		LastName:  "maeda",
		Email:     "aaa@bbb.ccc",
	}
}

func NewInValidUpdateStudentInputForm() *dto.UpdateStudentInputForm {
	return &dto.UpdateStudentInputForm{
		ID:        0,
		StudentNo: "121312312",
		FirstName: "taku",
		LastName:  "maeda",
		Email:     "aaa@bbb.ccc",
	}
}

func NewInValidUpdateAssistantInputForm() *dto.UpdateAssistantInputForm {
	return &dto.UpdateAssistantInputForm{
		ID:        0,
		StudentNo: "121312312",
		FirstName: "taku",
		LastName:  "maeda",
		Email:     "aaa@bbb.ccc",
	}
}

func NewInValidUpdateTeacherInputForm() *dto.UpdateTeacherInputForm {
	return &dto.UpdateTeacherInputForm{
		ID:        0,
		FirstName: "taku",
		LastName:  "maeda",
		Email:     "aaa@bbb.ccc",
	}
}

//func NewCreateAdminInputFormMock() dto.CreateAdminInputForm {
//	return dto.CreateAdminInputForm{
//		FirstName: "taku",
//		LastName: "maeda",
//		Email: "aaa@bbb.ccc",
//		Password: "password",
//	}
//}
