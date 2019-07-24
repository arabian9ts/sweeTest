package controllers

type RootController struct {
	StudentsController *StudentsController
	TasController      *TasController
	TeachersController *TeachersController
	LecturesController *LecturesController
}

func NewRootController(
	studentsController *StudentsController,
	tasController *TasController,
	teachersController *TeachersController,
	lecturesController *LecturesController,
) (*RootController, error) {
	return &RootController{
		StudentsController:studentsController,
		TasController:tasController,
		TeachersController:teachersController,
		LecturesController:lecturesController,
	}, nil
}
