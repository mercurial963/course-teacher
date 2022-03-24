package teacher_models

import (
	"course-teacher/db"
)

func (h Teacher) GetByID(id string) (*Teacher, error) {
	db := db.GetDB()
	teachers := Teacher{}
	err := db.Get(&teachers, "SELECT * FROM teachers WHERE teacher_id=$1", id)
	return &teachers, err
}
