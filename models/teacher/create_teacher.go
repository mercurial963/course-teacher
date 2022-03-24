package teacher_models

import (
	"course-teacher/db"
)

func (t Teacher) Create(b *Teacher) (*Teacher, error) {

	db := db.GetDB()
	teacher := Teacher{}
	err := db.Get(&teacher, "INSERT INTO teachers(first_name, last_name, date_of_birth) VALUES ($1, $2, $3) RETURNING *;", b.Firstname, b.Lastname, b.DateOfBirth)
	return &teacher, err
}
