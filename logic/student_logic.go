package logic

import (
	"student-service/dbx"
	"student-service/model"

	"golang.org/x/crypto/bcrypt"
)

func CreateStudent(s model.Student) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(s.Password), 14)
	_, err := dbx.DB.Exec("INSERT INTO students (name, email, password) VALUES (?, ?, ?)", s.Name, s.Email, hash)
	return err
}

func GetStudentByEmail(email string) (model.Student, error) {
	var s model.Student
	err := dbx.DB.QueryRow("SELECT id, name, email, password FROM students WHERE email = ?", email).
		Scan(&s.ID, &s.Name, &s.Email, &s.Password)
	if err != nil {
		return s, err
	}
	return s, nil
}

func GetAllStudents() ([]model.Student, error) {
	rows, err := dbx.DB.Query("SELECT id, name, email FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.Student
	for rows.Next() {
		var s model.Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Email); err == nil {
			students = append(students, s)
		}
	}
	return students, nil
}

func DeleteStudentByID(id int) error {
	_, err := dbx.DB.Exec("DELETE FROM students WHERE id = ?", id)
	return err
}

func UpdateStudentName(id int, name string) error {
	_, err := dbx.DB.Exec("UPDATE students SET name = ? WHERE id = ?", name, id)
	return err
}
