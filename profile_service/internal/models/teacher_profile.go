package models

import "github.com/google/uuid"

type TeacherProfile struct {
	Id                    uuid.UUID `db:"id"`
	Education             string    `db:"education"`
	ScientificExperience  string    `db:"scientific_experience"`
	TeachingExperience    string    `db:"teaching_experience"`
	ProfessionalInterests string    `db:"professional_interests"`
	Achievements          string    `db:"achievements"`
	Languages             string    `db:"languages"`
	UserProfileId         uuid.UUID `db:"profile_id"`
}
