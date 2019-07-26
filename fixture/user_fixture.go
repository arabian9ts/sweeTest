package fixture

import (
	"time"

	"github.com/arabian9ts/sweeTest/app/domain/model"
)

func NewValidStudent() *model.Student {
	return &model.Student{
		ID:        1,
		StudentNo: "12345678",
		FirstName: "First",
		LastName:  "Last",
		Email:     "Email@test.com",
		Digest:    "Digest",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewValidAssistant() *model.Assistant {
	return &model.Assistant{
		ID:        1,
		StudentNo: "12345678",
		FirstName: "First",
		LastName:  "Last",
		Email:     "Email@test.com",
		Digest:    "Digest",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewValidTeacher() *model.Teacher {
	return &model.Teacher{
		ID:        1,
		FirstName: "First",
		LastName:  "Last",
		Email:     "Email@test.com",
		Digest:    "Digest",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewValidAdmin() *model.Admin {
	return &model.Admin{
		ID:        1,
		FirstName: "First",
		LastName:  "Last",
		Email:     "Email@test.com",
		Digest:    "Digest",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
