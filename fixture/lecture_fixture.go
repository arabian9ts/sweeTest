package fixture

import (
	"time"

	"github.com/arabian9ts/sweeTest/app/domain/model"
)

func NewValidLecture() *model.Lecture {
	return &model.Lecture{
		ID:          1,
		Name:        "test lecture",
		TeacherName: "test name",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
