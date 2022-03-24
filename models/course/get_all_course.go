package course_models

import (
	"course-teacher/db"
)

func (h Course) GetByID(id string) (*Course, error) {
	db := db.GetDB()
	courses := Course{}
	err := db.Get(&courses, "SELECT * FROM courses WHERE course_id=$1", id)
	return &courses, err
}
