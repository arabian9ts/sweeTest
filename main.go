package main

import (
	"fmt"
	"github.com/arabian9ts/sweeTest/app/builder"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/interactor"
)

func main() {
	repo, _ := builder.InitializeUserRepository()
	interactor, err := interactor.NewCreateUserInteractor(repo)
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
	output := interactor.CreateStudent(form)
	fmt.Println(output)

	student, err := repo.GetStudentById(output.LastCreatedUserId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(student)
}
