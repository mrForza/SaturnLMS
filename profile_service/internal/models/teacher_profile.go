package models

import "github.com/google/uuid"

type TeacherProfile struct {
	Id                    uuid.UUID
	Education             string
	ScientificExperience  string
	TeachingExperience    string
	ProfessionalInterests string
	Achievements          string
	Languages             string
	UserProfileId         uuid.UUID
}
