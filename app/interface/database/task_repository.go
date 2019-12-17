package database

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type TaskRepository struct {
	SqlHandler
}

func NewTaskRepository(sqlHandler SqlHandler) (repository.TaskRepository, error) {
	return &TaskRepository{SqlHandler: sqlHandler}, nil
}

func (repo *TaskRepository) GetTasksByClassId(classID int64, limit int, offset int) (model.Tasks, error) {
	tasks := model.Tasks{}
	rows, err := repo.SqlHandler.Query("SELECT * FROM `tasks` WHERE `tasks`.`class_id` = ? LIMIT ? OFFSET ?", classID, limit, offset)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		task := model.Task{}
		if err = rows.Scan(&task.ID, &task.ClassID, &task.Title, &task.Desc, &task.Deadline, &task.CreatedAt, &task.UpdatedAt); err != nil {
			continue
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (repo *TaskRepository) CreateTask(task *model.Task) (*model.Task, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO `tasks` (`class_id`, `title`, `desc`, `deadline`) VALUES (?,?,?,?)",
		task.ClassID,
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
		"UPDATE `tasks` SET `class_id` = ?, `title` = ?, `desc` = ?, `deadline` = ? WHERE `id` = ?",
		task.ClassID,
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

func (repo *TaskRepository) DeleteTask(id int64, classID int64) (int64, error) {
	result, err := repo.SqlHandler.Execute("DELETE FROM `tasks` WHERE `id` = ? AND `class_id` = ?", id, classID)
	if err != nil {
		return 0, err
	}

	count, err := result.RowAffected()
	if err != nil {
		return 0, err
	}

	return int64(count), nil
}
