package main

import (
	"github.com/labstack/echo/v4"
	configurationvar "survivor-coders-students/configuration"
	"survivor-coders-students/controllers"
)

func main() {
	e := echo.New()
	configurationvar.DbConfig()
	e.GET("/students", controllers.GetAllStudents)
	e.PUT("/students/:id", controllers.PutStudent)
	e.POST("/students", controllers.PostStudent)
	e.PATCH("/students/:id", controllers.PatchStudent)
	e.DELETE("/students/:id", controllers.DeleteStudent)
	e.Start("localhost:8080")
}
