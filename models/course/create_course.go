package course_models

import (
	"course-teacher/db"
)

func (t Course) Create(b *Course) (*Course, error) {

	db := db.GetDB()
	course := Course{}
	err := db.Get(&course, "INSERT INTO courses(course_name, course_description, teacher_id) VALUES ($1, $2, $3) RETURNING *;", b.Name, b.Description, b.TeacherID)
	return &course, err
}
