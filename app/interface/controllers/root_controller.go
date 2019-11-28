package controllers

type RootController struct {
	StudentsController            *StudentsController
	AssistantsController          *AssistantsController
	TeachersController            *TeachersController
	LecturesController            *LecturesController
	TasksController               *TasksController
	LoginController               *LoginController
	HelpsController               *HelpsController
	CommentsController            *CommentsController
	ParticipationController       *ParticipationController
	ParticipatedLectureController *ParticipatingLecturesController
	ParticipantsController        *ParticipantsController
}

func NewRootController(
	studentsController *StudentsController,
	assistantsController *AssistantsController,
	teachersController *TeachersController,
	lecturesController *LecturesController,
	tasksController *TasksController,
	loginController *LoginController,
	helpsController *HelpsController,
	commentsController *CommentsController,
	participationController *ParticipationController,
	participatedLectureController *ParticipatingLecturesController,
	participantsController *ParticipantsController,
) (*RootController, error) {
	return &RootController{
		StudentsController:            studentsController,
		AssistantsController:          assistantsController,
		TeachersController:            teachersController,
		LecturesController:            lecturesController,
		TasksController:               tasksController,
		LoginController:               loginController,
		HelpsController:               helpsController,
		CommentsController:            commentsController,
		ParticipationController:       participationController,
		ParticipatedLectureController: participatedLectureController,
		ParticipantsController:        participantsController,
	}, nil
}
