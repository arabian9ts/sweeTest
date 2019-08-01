package adapter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

func ConvertCreateHelpInputFormToHelp(form *dto.CreateHelpInputForm) (help *model.Help) {
	help = &model.Help{
		LectureID: form.LectureID,
		StudentID: form.StudentID,
		Contents:  form.Contents,
	}
	return
}

func ConvertUpdateHelpInputFormToHelp(form *dto.UpdateHelpInputForm) (help *model.Help) {
	help = &model.Help{
		ID:        form.ID,
		LectureID: form.LectureID,
		StudentID: form.StudentID,
		Contents:  form.Contents,
	}
	return
}
