package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	configurationvar "survivor-coders-students/configuration"
	"survivor-coders-students/models"
)

func GetAllStudents(c echo.Context) error {
	var students []models.Student
	configurationvar.Db.Find(&students)
	return c.JSON(http.StatusOK, students)
}

func PutStudent(c echo.Context) error {
	var student models.Student
	var saved models.Student
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&student)
	configurationvar.Db.Find(&saved, id)
	saved.Firstname = student.Firstname
	saved.Lastname = student.Lastname
	configurationvar.Db.Save(&saved)
	return c.JSON(http.StatusCreated, saved)
}

func PostStudent(c echo.Context) error {
	var student models.Student
	c.Bind(&student)
	configurationvar.Db.Create(&student)
	return c.JSON(http.StatusCreated, student)
}

func PatchStudent(c echo.Context) error {
	var student models.Student
	var saved models.Student
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&student)
	configurationvar.Db.Find(&saved, id)
	if student.Firstname != "" {
		saved.Firstname = student.Firstname
	}
	if student.Lastname != "" {
		saved.Lastname = student.Lastname
	}
	configurationvar.Db.Save(&saved)
	return c.JSON(http.StatusCreated, saved)
}

func DeleteStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	configurationvar.Db.Delete(&models.Student{}, id)
	return c.JSON(http.StatusNoContent, nil)
}
