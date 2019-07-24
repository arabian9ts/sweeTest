package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type TaskRepository interface {
	GetTasksByLectureId(lectureId int64, limit int, offset int) (model.Tasks, error)
	CreateTask(task *model.Task) (int64, error)
	UpdateTask(task *model.Task) (int64, error)
	DeleteTask(id int64, lectureID int64) (int64, error)
}
