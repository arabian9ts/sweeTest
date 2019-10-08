package fixture

import "github.com/arabian9ts/sweeTest/app/dto"

func NewLectureInputFormFixture() *dto.CreateLectureInputForm {
	return &dto.CreateLectureInputForm{
		Name: "lecture input form fixture",
	}
}

func NewUpdateLectureInputFormFixture() *dto.UpdateLectureInputForm {
	return &dto.UpdateLectureInputForm{
		ID:   1,
		Name: "lecture input form fixture",
	}
}
