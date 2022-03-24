package teacher_models

type Teacher struct {
	ID          int    `json:"id" db:"teacher_id"`
	Firstname   string `json:"firstname" db:"first_name" validate:"required,min=4,max=15"`
	Lastname    string `json:"lastname" db:"last_name" validate:"required,min=4,max=15"`
	DateOfBirth int    `json:"dateofbirth" db:"date_of_birth" validate:"numeric"`
}
