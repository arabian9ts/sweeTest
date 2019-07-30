package controllers

type RootController struct {
	StudentsController       *StudentsController
	AssistantsController     *AssistantsController
	TeachersController       *TeachersController
	LecturesController       *LecturesController
	TasksController          *TasksController
	StudentLoginController   *StudentLoginController
	AssistantLoginController *AssistantLoginController
	TeacherLoginController   *TeacherLoginController
	AdminLoginController     *AdminLoginController
}

func NewRootController(
	studentsController *StudentsController,
	assistantsController *AssistantsController,
	teachersController *TeachersController,
	lecturesController *LecturesController,
	tasksController *TasksController,
	studentLoginController *StudentLoginController,
	assistantLoginController *AssistantLoginController,
	teacherLoginController *TeacherLoginController,
	adminLoginController *AdminLoginController,
) (*RootController, error) {
	return &RootController{
		StudentsController:       studentsController,
		AssistantsController:     assistantsController,
		TeachersController:       teachersController,
		LecturesController:       lecturesController,
		TasksController:          tasksController,
		StudentLoginController:   studentLoginController,
		AssistantLoginController: assistantLoginController,
		TeacherLoginController:   teacherLoginController,
		AdminLoginController:     adminLoginController,
	}, nil
}
