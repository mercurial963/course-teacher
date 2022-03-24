package course_models

import (
	"course-teacher/db"
)

func (h Course) Delete(id string) (*Course, error) {
	db := db.GetDB()
	courses := Course{}
	err := db.Get(&courses, "DELETE FROM courses WHERE course_id=$1 RETURNING *", id)
	return &courses, err
}
