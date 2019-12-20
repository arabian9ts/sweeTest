package fixture

import "github.com/arabian9ts/sweeTest/app/dto"

func NewLectureInputFormFixture() *dto.CreateLectureInputForm {
	return &dto.CreateLectureInputForm{
		Name:        "lecture input form fixture",
		TeacherName: "lecture input form fixture",
	}
}

func NewUpdateLectureInputFormFixture() *dto.UpdateLectureInputForm {
	return &dto.UpdateLectureInputForm{
		ID:          1,
		Name:        "update lecture input form fixture",
		TeacherName: "update lecture input form fixture",
	}
}
