package fixture

import "github.com/arabian9ts/sweeTest/app/dto"

func NewCreateStudentInputFormMock() dto.CreateStudentInputForm {
	return dto.CreateStudentInputForm{
		StudentNo: "121312312",
		FirstName: "taku",
		LastName:  "maeda",
		Email:     "aaa@bbb.ccc",
		Password:  "password",
	}
}

func NewCreateTaInputFormMock() dto.CreateTaInputForm {
	return dto.CreateTaInputForm{
		StudentNo: "121312312",
		FirstName: "taku",
		LastName:  "maeda",
		Email:     "aaa@bbb.ccc",
		Password:  "password",
	}
}

func NewCreateTeacherInputFormMock() dto.CreateTeacherInputForm {
	return dto.CreateTeacherInputForm{
		FirstName: "taku",
		LastName:  "maeda",
		Email:     "aaa@bbb.ccc",
		Password:  "password",
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
