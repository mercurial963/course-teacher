package teacher_models

import (
	"course-teacher/db"
)

func (h Teacher) Delete(id string) (*Teacher, error) {
	db := db.GetDB()
	teachers := Teacher{}
	err := db.Get(&teachers, "DELETE FROM teachers WHERE teacher_id=$1 RETURNING *", id)
	return &teachers, err
}
