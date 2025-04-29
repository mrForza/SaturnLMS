package models

import "github.com/google/uuid"

type StudentProfile struct {
	UniversityName string
	FacultatyName  string
	ProgramName    string
	GroupNumber    int64
	CourseNumber   uint8
	UserProfileId  uuid.UUID
}
