package models

import (
	"time"

	"github.com/google/uuid"
)

type SystemUserProfile struct {
	Id               uuid.UUID `gorm:"primaryKey"`
	Email            string
	Password         string
	LastLogin        time.Time
	LastActive       time.Time
	RegistrationTime time.Time
}
