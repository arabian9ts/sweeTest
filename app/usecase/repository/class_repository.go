package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type ClassRepository interface {
	GetClassesByLectureID(lectureID int64, limit int, offset int) (model.Classes, error)
	GetClassById(id int64) (class *model.Class, err error)
	CreateClass(class *model.Class) (*model.Class, error)
	UpdateClass(help *model.Class) (*model.Class, error)
	DeleteClass(id int64, lectureID int64) (int64, error)
}
