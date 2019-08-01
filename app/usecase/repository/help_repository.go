package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type HelpRepository interface {
	GetHelpsByStudentID(studentID int64, limit int, offset int) (model.Helps, error)
	GetHelpsByLectureID(lectureID int64, limit int, offset int) (model.Helps, error)
	CreateHelp(help *model.Help) (int64, error)
	UpdateHelp(help *model.Help) (int64, error)
	DeleteHelp(id int64, lectureID int64, studentID int64) (int64, error)
}
