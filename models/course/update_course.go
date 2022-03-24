package course_models

import (
	"course-teacher/db"
	"fmt"

	"gopkg.in/validator.v2"
)

func (t Course) UpdateCourse(id string, b *Course) (*Course, error) {
	if errs := validator.Validate(b.Name); errs != nil {
		fmt.Println(errs)
	}
	db := db.GetDB()
	course := Course{}
	err := db.Get(&course, "UPDATE courses SET course_name=$1,course_description=$2 WHERE course_id=$3 RETURNING *", b.Name, b.Description, id)
	return &course, err
}
