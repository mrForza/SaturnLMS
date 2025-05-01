package models

import "github.com/google/uuid"

type Gender bool

const (
	Male   Gender = false
	Female Gender = true
)

type UserProfile struct {
	Id                   uuid.UUID       `db:"id"`
	FirstName            string          `db:"first_name"`
	LastName             string          `db:"last_name"`
	FatherName           string          `db:"father_name"`
	Age                  uint8           `db:"age"`
	Gender               Gender          `db:"gender"`
	AboutMe              string          `db:"about_me"`
	Interests            string          `db:"interests"`
	StudentProfile       *StudentProfile `db:"-"`
	TeacherProfile       *TeacherProfile `db:"-"`
	AdministratorProfile *AdminProfile   `db:"-"`
}
