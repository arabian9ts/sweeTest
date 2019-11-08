package controllers

type RootController struct {
	StudentsController   *StudentsController
	AssistantsController *AssistantsController
	TeachersController   *TeachersController
	LecturesController   *LecturesController
	TasksController      *TasksController
	LoginController      *LoginController
	HelpsController      *HelpsController
	CommentsController   *CommentsController
}

func NewRootController(
	studentsController *StudentsController,
	assistantsController *AssistantsController,
	teachersController *TeachersController,
	lecturesController *LecturesController,
	tasksController *TasksController,
	loginController *LoginController,
	helpsCOntroller *HelpsController,
	commentsController *CommentsController,
) (*RootController, error) {
	return &RootController{
		StudentsController:   studentsController,
		AssistantsController: assistantsController,
		TeachersController:   teachersController,
		LecturesController:   lecturesController,
		TasksController:      tasksController,
		LoginController:      loginController,
		HelpsController:      helpsCOntroller,
		CommentsController:   commentsController,
	}, nil
}
