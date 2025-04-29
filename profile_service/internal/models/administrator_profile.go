package models

import "github.com/google/uuid"

type AdministratorProfile struct {
	Id             uuid.UUID
	Education      string
	WorkExperience string
	Achievements   string
	Languages      string
	UserProfileId  uuid.UUID
}
