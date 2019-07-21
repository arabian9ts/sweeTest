package database

import (
	"github.com/arabian9ts/sweeTest/app/repository"
	"github.com/arabian9ts/sweeTest/app/domain/model"
)

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) (repository.UserRepository, error) {
	if sqlHandler == nil {
		return nil, nil
	}
	repo := &UserRepository{SqlHandler: sqlHandler}
	return repo, nil
}

func (repo *UserRepository) GetStudentById(studentId int64) (student *model.Student, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM students WHERE id = ?", studentId)
	if err != nil {
		return &model.Student{}, err
	}
	defer row.Close()
	row.Next()


	student = &model.Student{}
	err = row.Scan(&student.ID, &student.StudentNo, &student.FirstName, &student.LastName, &student.Email, &student.Digest, &student.CreatedAt, &student.UpdatedAt)
	return
}

func (repo *UserRepository) GetTaById(taId int64) (ta *model.Ta, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM tas WHERE id = ?", taId)
	if err != nil {
		return &model.Ta{}, err
	}
	defer row.Close()
	row.Next()

	ta = &model.Ta{}
	err = row.Scan(&ta.ID, &ta.StudentNo, &ta.FirstName, &ta.LastName, &ta.Email, &ta.Digest, &ta.CreatedAt, &ta.UpdatedAt)
	return
}

func (repo *UserRepository) GetTeacherById(teacherId int64) (teacher *model.Teacher, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM teachers WHERE id = ?", teacherId)
	if err != nil {
		return &model.Teacher{}, err
	}
	defer row.Close()
	row.Next()

	teacher = &model.Teacher{}
	err = row.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Digest, &teacher.CreatedAt, &teacher.UpdatedAt)
	return
}

func (repo *UserRepository) GetAdminById(adminId int64) (admin *model.Admin, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM admins WHERE id = ?", adminId)
	if err != nil {
		return &model.Admin{}, err
	}
	defer row.Close()
	row.Next()

	admin = &model.Admin{}
	err = row.Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.Email, &admin.Digest, &admin.CreatedAt, &admin.UpdatedAt)
	return
}

func (repo *UserRepository) InsertStudent(student *model.Student) (int64, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO students (student_no, first_name, last_name, email, digest) VALUES (?,?,?,?,?)",
		student.StudentNo, student.FirstName, student.LastName, student.Email, student.Digest,
	)
	if err != nil {
		return 0, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(id64), nil
}

func (repo *UserRepository) InsertTa(ta *model.Ta) (int64, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO tas (student_no, first_name, last_name, email, digest) VALUES (?,?,?,?,?)",
		ta.StudentNo, ta.FirstName, ta.LastName, ta.Email, ta.Digest,
	)
	if err != nil {
		return 0, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(id64), nil
}

func (repo *UserRepository) InsertTeacher(teacher *model.Teacher) (int64, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO teachers (first_name, last_name, email, digest) VALUES (?,?,?,?)",
		teacher.FirstName, teacher.LastName, teacher.Email, teacher.Digest,
	)
	if err != nil {
		return 0, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(id64), nil
}

func (repo *UserRepository) InsertAdmin(admin *model.Admin) (int64, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO admins (first_name, last_name, email, digest) VALUES (?,?,?,?)",
		admin.FirstName, admin.LastName, admin.Email, admin.Digest,
	)
	if err != nil {
		return 0, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(id64), nil
}
