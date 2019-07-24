package controllers

type RootController struct {
	StudentsController *StudentsController
	TasController      *TasController
	TeachersController *TeachersController
	LecturesController *LecturesController
	TasksController    *TasksController
}

func NewRootController(
	studentsController *StudentsController,
	tasController *TasController,
	teachersController *TeachersController,
	lecturesController *LecturesController,
	tasksController *TasksController,
) (*RootController, error) {
	return &RootController{
		StudentsController: studentsController,
		TasController:      tasController,
		TeachersController: teachersController,
		LecturesController: lecturesController,
		TasksController:    tasksController,
	}, nil
}
