package main

import (
	"fmt"
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/builder")

func main() {
	repo, _ := builder.InitializeLectureRepository()
	fmt.Println(repo)
	lecture := &model.Lecture{ID: 9, Name: "aaa"}

	id, err := repo.UpdateLecture(lecture)
	fmt.Println(id)
	fmt.Println(err)
}
