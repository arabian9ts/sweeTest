package controllers

import "github.com/arabian9ts/sweeTest/app/domain/model"

func getCurrentStudent(ctx Context) (student *model.Student) {
	user, exists := ctx.Get("currentUser")
	if !exists {
		return &model.Student{}
	}
	student = user.(*model.Student)
	return
}

func getCurrentAssistant(ctx Context) (assistant *model.Assistant) {
	user, exists := ctx.Get("currentUser")
	if !exists {
		return &model.Assistant{}
	}
	assistant = user.(*model.Assistant)
	return
}

func getCurrentTeacher(ctx Context) (teacher *model.Teacher) {
	user, exists := ctx.Get("currentUser")
	if !exists {
		return &model.Teacher{}
	}
	teacher = user.(*model.Teacher)
	return
}

func getCurrentAdmin(ctx Context) (admin *model.Admin) {
	user, exists := ctx.Get("currentUser")
	if !exists {
		return &model.Admin{}
	}
	admin = user.(*model.Admin)
	return
}
