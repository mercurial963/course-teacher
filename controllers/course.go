package controllers

import (
	course_models "course-teacher/models/course"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

type CourseController struct{}

var courseModel = new(course_models.Course)

func (u CourseController) Create(c *gin.Context) {
	var body course_models.Course
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if errs2 := validator.Validate(body); errs2 != nil {
		fmt.Println(errs2)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to save teacher", "error": errs2})
		return
	}
	courses, err := courseModel.Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to save teacher", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success!", "course": courses})
	return
}

func (u CourseController) GetByID(c *gin.Context) {
	id := c.Param("id")
	courses, err := courseModel.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve course", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success!", "course": courses})
	return
}

func (u CourseController) UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var body course_models.Course
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if errs2 := validator.Validate(body); errs2 != nil {
		fmt.Println(errs2)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to save teacher", "error": errs2})
		return
	}
	courses, err := courseModel.UpdateCourse(id, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to save course", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success!", "course": courses})
	return
}

func (u CourseController) Delete(c *gin.Context) {
	id := c.Param("id")
	courses, err := courseModel.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve course", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success!", "course": courses})
	return
}
