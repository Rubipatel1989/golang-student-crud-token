package handler

import (
	"os"
	"strconv"
	"time"

	"student-service/gtservices/responsex"
	"student-service/logic"
	"student-service/model"
	"student-service/request"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	var req request.StudentSignupRequest
	if err := c.BodyParser(&req); err != nil {
		return responsex.BadRequest(c, "Invalid request body")
	}

	student := model.Student{Name: req.Name, Email: req.Email, Password: req.Password}
	if err := logic.CreateStudent(student); err != nil {
		return responsex.GTError(c, 500, "Signup failed")
	}
	return responsex.GTSuccess(c, true, "Signup successful", nil)
}

func Login(c *fiber.Ctx) error {
	var req request.StudentLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return responsex.BadRequest(c, "Invalid input")
	}

	s, err := logic.GetStudentByEmail(req.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(s.Password), []byte(req.Password)) != nil {
		return responsex.Unauthorized(c, "Invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id": s.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return responsex.GTSuccess(c, true, "Login successful", fiber.Map{"token": t})
}

func ListStudents(c *fiber.Ctx) error {
	students, err := logic.GetAllStudents()
	if err != nil {
		return responsex.GTError(c, 500, "Failed to fetch students")
	}
	return responsex.GTSuccess(c, true, "Students retrieved", students)
}

func UpdateStudent(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var data map[string]string
	c.BodyParser(&data)
	err := logic.UpdateStudentName(id, data["name"])
	if err != nil {
		return responsex.GTError(c, 500, "Update failed")
	}
	return responsex.GTSuccess(c, true, "Updated", nil)
}

func DeleteStudent(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := logic.DeleteStudentByID(id)
	if err != nil {
		return responsex.GTError(c, 500, "Delete failed")
	}
	return responsex.GTSuccess(c, true, "Deleted", nil)
}
