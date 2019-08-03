package controllers

type RootController struct {
	StudentsController          *StudentsController
	AssistantsController        *AssistantsController
	TeachersController          *TeachersController
	LecturesController          *LecturesController
	TasksController             *TasksController
	StudentLoginController      *StudentLoginController
	AssistantLoginController    *AssistantLoginController
	TeacherLoginController      *TeacherLoginController
	AdminLoginController        *AdminLoginController
	HelpsController             *HelpsController
	StudentCommentsController   *StudentCommentsController
	AssistantCommentsController *AssistantCommentsController
	TeacherCommentsController   *TeacherCommentsController
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
	helpsCOntroller *HelpsController,
	studentCommentsController *StudentCommentsController,
	assistantCommentsController *AssistantCommentsController,
	teacherCommentsController *TeacherCommentsController,
) (*RootController, error) {
	return &RootController{
		StudentsController:          studentsController,
		AssistantsController:        assistantsController,
		TeachersController:          teachersController,
		LecturesController:          lecturesController,
		TasksController:             tasksController,
		StudentLoginController:      studentLoginController,
		AssistantLoginController:    assistantLoginController,
		TeacherLoginController:      teacherLoginController,
		AdminLoginController:        adminLoginController,
		HelpsController:             helpsCOntroller,
		StudentCommentsController:   studentCommentsController,
		AssistantCommentsController: assistantCommentsController,
		TeacherCommentsController:   teacherCommentsController,
	}, nil
}
