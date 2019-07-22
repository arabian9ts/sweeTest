package main

import (
	"fmt"
	"github.com/arabian9ts/sweeTest/app/builder"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/interface/presenter"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
)

func main() {
	repo, _ := builder.InitializeUserRepository()
	output := &presenter.UserPresenter{}
	interactor, err := interactor.NewUserInteractor(repo, output)
	if err != nil {
		fmt.Println(err)
	}

	form := dto.CreateStudentInputForm{
		StudentNo: "121312312",
		FirstName: "taku",
		LastName:  "maeda",
		Email:     "aaa@bbb.ccc",
		Password:  "password",
	}
	outputForm, _ := interactor.CreateStudent(form)
	fmt.Println(outputForm)

	student, err := repo.GetStudentById(outputForm.LastCreatedUserId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(student)
}
