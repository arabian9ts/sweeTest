package database

import (
	"time"

	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type TaskRepository struct {
	SqlHandler
}

func NewTaskRepository(sqlHandler SqlHandler) (repository.TaskRepository, error) {
	return &TaskRepository{SqlHandler: sqlHandler}, nil
}

func (repo *TaskRepository) GetTasksByLectureId(lectureId int64, limit int, offset int) (model.Tasks, error) {
	tasks := model.Tasks{}
	rows, err := repo.SqlHandler.Query("SELECT * FROM `tasks` WHERE `tasks`.`lecture_id` = ? LIMIT ? OFFSET ?", lectureId, limit, offset)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	var id int64
	var lectureID int64
	var title string
	var desc string
	var deadline time.Time
	var createdAt time.Time
	var updatedAt time.Time
	for rows.Next() {
		if err = rows.Scan(&id, &lectureID, &title, &desc, &deadline, &createdAt, &updatedAt); err != nil {
			continue
		}
		tasks = append(tasks, &model.Task{
			ID:        id,
			LectureID: lectureID,
			Title:     title,
			Desc:      desc,
			Deadline:  deadline,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return tasks, nil
}

func (repo *TaskRepository) CreateTask(task *model.Task) (*model.Task, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO `tasks` (`lecture_id`, `title`, `desc`, `deadline`) VALUES (?,?,?,?)",
		task.LectureID,
		task.Title,
		task.Desc,
		task.Deadline,
	)
	if err != nil {
		return nil, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	task.ID = id64
	return task, nil
}

func (repo *TaskRepository) UpdateTask(task *model.Task) (*model.Task, error) {
	result, err := repo.SqlHandler.Execute(
		"UPDATE `tasks` SET `lecture_id` = ?, `title` = ?, `desc` = ?, `deadline` = ? WHERE `id` = ?",
		task.LectureID,
		task.Title,
		task.Desc,
		task.Deadline,
		task.ID,
	)
	if err != nil {
		return nil, err
	}

	_, err = result.RowAffected()
	if err != nil {
		return nil, err
	}

	return task, err
}

func (repo *TaskRepository) DeleteTask(id int64, lectureID int64) (int64, error) {
	result, err := repo.SqlHandler.Execute("DELETE FROM `tasks` WHERE `id` = ? AND `lecture_id` = ?", id, lectureID)
	if err != nil {
		return 0, err
	}

	count, err := result.RowAffected()
	if err != nil {
		return 0, err
	}

	return int64(count), nil
}
