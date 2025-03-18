package models

import "github.com/google/uuid"

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type UserProfile struct {
	Id                   uuid.UUID `gorm:"primaryKey"`
	FirstName            string
	LastName             string
	FatherName           *string
	Age                  uint8
	Gender               Gender
	AboutMe              *string
	Interests            *string
	StudentProfile       *StudentProfile       `gorm:"foreignKey:Id"`
	TeacherProfile       *TeacherProfile       `gorm:"foreignKey:Id"`
	AdministratorProfile *AdministratorProfile `gorm:"foreignKey:Id"`
}
