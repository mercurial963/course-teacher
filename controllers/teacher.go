package controllers

import (
	teacher_models "course-teacher/models/teacher"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

type TeacherController struct{}

var teacherModel = new(teacher_models.Teacher)

func (u TeacherController) Create(c *gin.Context) {

	var body teacher_models.Teacher
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if errs2 := validator.Validate(body); errs2 != nil {
		fmt.Println(errs2)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to save teacher", "error": errs2})
		return
	}
	teachers, err := teacherModel.Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to save teacher", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success!", "teacher": teachers})
	return
}

func (u TeacherController) GetByID(c *gin.Context) {
	id := c.Param("id")
	teachers, err := teacherModel.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve teacher", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success!", "teacher": teachers})
	return
}

func (u TeacherController) UpdateName(c *gin.Context) {
	id := c.Param("id")
	var body teacher_models.Teacher

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if errs2 := validator.Validate(body); errs2 != nil {
		fmt.Println(errs2)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to save teacher", "error": errs2})
		return
	}
	teachers, err := teacherModel.UpdateName(id, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve teacher", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success!", "teacher": teachers})
	return
}

func (u TeacherController) Delete(c *gin.Context) {
	id := c.Param("id")
	teachers, err := teacherModel.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve teacher", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success!", "teacher": teachers})
	return
}
