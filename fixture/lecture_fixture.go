package fixture

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"time"
)

func NewValidLecture() *model.Lecture {
	return &model.Lecture{
		ID:        1,
		Name:      "test lecture",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
