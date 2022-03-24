package teacher_models

import (
	"course-teacher/db"
)

func (t Teacher) UpdateName(id string, b *Teacher) (*Teacher, error) {

	db := db.GetDB()
	teacher := Teacher{}
	err := db.Get(&teacher, "UPDATE teachers SET first_name=$1,last_name=$2 WHERE teacher_id=$3 RETURNING *;", b.Firstname, b.Lastname, id)
	return &teacher, err
}
