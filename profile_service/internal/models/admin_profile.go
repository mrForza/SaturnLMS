package models

import "github.com/google/uuid"

type AdminProfile struct {
	Id             uuid.UUID `db:"id"`
	Education      string    `db:"education"`
	WorkExperience string    `db:"work_experience"`
	Achievements   string    `db:"achievements"`
	Languages      string    `db:"languages"`
	UserProfileId  uuid.UUID `db:"profile_id"`
}
