package models

import "github.com/google/uuid"

type StudentProfile struct {
	Id             uuid.UUID `db:"id"`
	UniversityName string    `db:"university_name"`
	FacultatyName  string    `db:"facultaty_name"`
	ProgramName    string    `db:"program_name"`
	GroupNumber    int64     `db:"group_number"`
	CourseNumber   uint8     `db:"course_number"`
	UserProfileId  uuid.UUID `db:"profile_id"`
}
