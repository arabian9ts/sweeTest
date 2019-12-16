package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type HelpRepository interface {
	GetHelpsByStudentID(studentID int64, limit int, offset int) (model.Helps, error)
	GetHelpsByLectureID(lectureID int64, limit int, offset int) (model.Helps, error)
	CreateHelp(help *model.Help) (*model.Help, error)
	UpdateHelp(help *model.Help) (*model.Help, error)
	DeleteHelp(id int64, lectureID int64, studentID int64) (int64, error)
	DeleteHelpWithoutStudentId(id int64, lectureID int64) (int64, error)
}
