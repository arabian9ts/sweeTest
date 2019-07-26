package controllers

type RootController struct {
	StudentsController   *StudentsController
	AssistantsController *AssistantsController
	TeachersController   *TeachersController
	LecturesController   *LecturesController
	TasksController      *TasksController
}

func NewRootController(
	studentsController *StudentsController,
	assistantsController *AssistantsController,
	teachersController *TeachersController,
	lecturesController *LecturesController,
	tasksController *TasksController,
) (*RootController, error) {
	return &RootController{
		StudentsController:   studentsController,
		AssistantsController: assistantsController,
		TeachersController:   teachersController,
		LecturesController:   lecturesController,
		TasksController:      tasksController,
	}, nil
}
