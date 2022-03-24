package course_models

type Course struct {
	ID          int    `json:"id" db:"course_id"`
	Name        string `json:"coursename" db:"course_name" validate:"required,min=4,max=15"`
	Description string `json:"description" db:"course_description"`
	TeacherID   int    `json:"teacherid" db:"teacher_id" validate:"numeric"`
}
