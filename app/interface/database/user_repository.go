package database

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
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

func (repo *UserRepository) GetStudents(limit, offset int) (students model.Students, err error) {
	rows, err := repo.SqlHandler.Query("SELECT * FROM `students` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		student := &model.Student{}
		err = rows.Scan(&student.ID, &student.StudentNo, &student.FirstName, &student.LastName, &student.Email, &student.Digest, &student.CreatedAt, &student.UpdatedAt)
		if err != nil {
			continue
		}
		students = append(students, student)
	}

	return
}

func (repo *UserRepository) GetAssistants(limit, offset int) (assistants model.Assistants, err error) {
	rows, err := repo.SqlHandler.Query("SELECT * FROM `assistants` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		assistant := &model.Assistant{}
		err = rows.Scan(&assistant.ID, &assistant.StudentNo, &assistant.FirstName, &assistant.LastName, &assistant.Email, &assistant.Digest, &assistant.CreatedAt, &assistant.UpdatedAt)
		if err != nil {
			continue
		}
		assistants = append(assistants, assistant)
	}

	return
}

func (repo *UserRepository) GetTeachers(limit, offset int) (teachers model.Teachers, err error) {
	rows, err := repo.SqlHandler.Query("SELECT * FROM `teachers` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		teacher := &model.Teacher{}
		err = rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Digest, &teacher.CreatedAt, &teacher.UpdatedAt)
		if err != nil {
			continue
		}
		teachers = append(teachers, teacher)
	}

	return
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

func (repo *UserRepository) GetAssistantById(assistantId int64) (assistant *model.Assistant, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM assistants WHERE id = ?", assistantId)
	if err != nil {
		return &model.Assistant{}, err
	}
	defer row.Close()
	row.Next()

	assistant = &model.Assistant{}
	err = row.Scan(&assistant.ID, &assistant.StudentNo, &assistant.FirstName, &assistant.LastName, &assistant.Email, &assistant.Digest, &assistant.CreatedAt, &assistant.UpdatedAt)
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

func (repo *UserRepository) GetStudentByStudentNo(studentNo string) (student *model.Student, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM `students` WHERE `student_no` = ?", studentNo)
	if err != nil {
		return &model.Student{}, err
	}
	defer row.Close()
	row.Next()

	student = &model.Student{}
	err = row.Scan(&student.ID, &student.StudentNo, &student.FirstName, &student.LastName, &student.Email, &student.Digest, &student.CreatedAt, &student.UpdatedAt)
	return
}

func (repo *UserRepository) GetAssistantByStudentNo(studentNo string) (assistant *model.Assistant, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM `assistants` WHERE `student_no` = ?", studentNo)
	if err != nil {
		return &model.Assistant{}, err
	}
	defer row.Close()
	row.Next()

	assistant = &model.Assistant{}
	err = row.Scan(&assistant.ID, &assistant.StudentNo, &assistant.FirstName, &assistant.LastName, &assistant.Email, &assistant.Digest, &assistant.CreatedAt, &assistant.UpdatedAt)
	return
}

func (repo *UserRepository) GetTeacherByEmail(email string) (teacher *model.Teacher, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM `teachers` WHERE `email` = ?", email)
	if err != nil {
		return &model.Teacher{}, err
	}
	defer row.Close()
	row.Next()

	teacher = &model.Teacher{}
	err = row.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Digest, &teacher.CreatedAt, &teacher.UpdatedAt)
	return
}

func (repo *UserRepository) GetParticipatingStudentsOfLecture(lectureID int64, limit int, offset int) (students model.Students, err error) {
	rows, err := repo.SqlHandler.Query("SELECT students.* FROM students INNER JOIN students_lectures ON students_lectures.student_id = students.id WHERE lecture_id = ? limit ? offset ?",
		lectureID,
		limit,
		offset,
	)
	if err != nil {
		return nil, nil
	}
	defer rows.Close()

	for rows.Next() {
		student := &model.Student{}
		err = rows.Scan(&student.ID, &student.StudentNo, &student.FirstName, &student.LastName, &student.Email, &student.Digest, &student.CreatedAt, &student.UpdatedAt)
		if err != nil {
			continue
		}
		students = append(students, student)
	}

	return
}

func (repo *UserRepository) GetParticipatingAssistantsOfLecture(lectureID int64, limit int, offset int) (assistants model.Assistants, err error) {
	rows, err := repo.SqlHandler.Query("SELECT assistants.* FROM assistants INNER JOIN assistants_lectures ON assistants_lectures.student_id = assistants.id WHERE lecture_id = ? limit ? offset ?",
		lectureID,
		limit,
		offset,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		assistant := &model.Assistant{}
		err = rows.Scan(&assistant.ID, &assistant.StudentNo, &assistant.FirstName, &assistant.LastName, &assistant.Email, &assistant.Digest, &assistant.CreatedAt, &assistant.UpdatedAt)
		if err != nil {
			continue
		}
		assistants = append(assistants, assistant)
	}

	return
}

func (repo *UserRepository) GetParticipatingTeachersOfLecture(lectureID int64, limit int, offset int) (teachers model.Teachers, err error) {
	rows, err := repo.SqlHandler.Query("SELECT teachers.* FROM teachers INNER JOIN teachers_lectures ON teachers_lectures.teacher_id = teachers.id WHERE lecture_id = ? limit ? offset ?",
		lectureID,
		limit,
		offset,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		teacher := &model.Teacher{}
		err = rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Digest, &teacher.CreatedAt, &teacher.UpdatedAt)
		if err != nil {
			continue
		}
		teachers = append(teachers, teacher)
	}

	return
}

func (repo *UserRepository) GetAdminByEmail(email string) (admin *model.Admin, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM `admins` WHERE `email` = ?", email)
	if err != nil {
		return &model.Admin{}, err
	}
	defer row.Close()
	row.Next()

	admin = &model.Admin{}
	err = row.Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.Email, &admin.Digest, &admin.CreatedAt, &admin.UpdatedAt)
	return
}

func (repo *UserRepository) InsertStudent(student *model.Student) (*model.Student, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO students (student_no, first_name, last_name, email, digest) VALUES (?,?,?,?,?)",
		student.StudentNo, student.FirstName, student.LastName, student.Email, student.Digest,
	)
	if err != nil {
		return nil, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	student.ID = id64
	return student, nil
}

func (repo *UserRepository) InsertAssistant(assistant *model.Assistant) (*model.Assistant, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO assistants (student_no, first_name, last_name, email, digest) VALUES (?,?,?,?,?)",
		assistant.StudentNo, assistant.FirstName, assistant.LastName, assistant.Email, assistant.Digest,
	)
	if err != nil {
		return nil, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	assistant.ID = id64
	return assistant, nil
}

func (repo *UserRepository) InsertTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO teachers (first_name, last_name, email, digest) VALUES (?,?,?,?)",
		teacher.FirstName, teacher.LastName, teacher.Email, teacher.Digest,
	)
	if err != nil {
		return nil, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	teacher.ID = id64
	return teacher, nil
}

func (repo *UserRepository) InsertAdmin(admin *model.Admin) (*model.Admin, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO admins (first_name, last_name, email, digest) VALUES (?,?,?,?)",
		admin.FirstName, admin.LastName, admin.Email, admin.Digest,
	)
	if err != nil {
		return nil, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	admin.ID = id64
	return admin, nil
}

func (repo *UserRepository) UpdateStudent(student *model.Student) (*model.Student, error) {
	result, err := repo.SqlHandler.Execute(
		"UPDATE students SET student_no = ?, first_name = ?, last_name = ?, email = ? WHERE id = ?",
		student.StudentNo,
		student.FirstName,
		student.LastName,
		student.Email,
		student.ID,
	)
	if err != nil {
		return nil, err
	}

	_, err = result.RowAffected()
	if err != nil {
		return nil, err
	}

	return student, err
}

func (repo *UserRepository) UpdateAssistant(assistant *model.Assistant) (*model.Assistant, error) {
	result, err := repo.SqlHandler.Execute(
		"UPDATE assistants SET assistant_no = ?, first_name = ?, last_name = ?, Email = ? WHERE id = ?",
		assistant.StudentNo,
		assistant.FirstName,
		assistant.LastName,
		assistant.Email,
		assistant.ID,
	)
	if err != nil {
		return nil, err
	}

	_, err = result.RowAffected()
	if err != nil {
		return nil, err
	}

	return assistant, err
}

func (repo *UserRepository) UpdateTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	result, err := repo.SqlHandler.Execute(
		"UPDATE teachers SET first_name = ?, last_name = ?, Email = ? WHERE id = ?",
		teacher.FirstName,
		teacher.LastName,
		teacher.Email,
		teacher.ID,
	)
	if err != nil {
		return nil, err
	}

	_, err = result.RowAffected()
	if err != nil {
		return nil, err
	}

	return teacher, err
}
