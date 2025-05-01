package models

import (
	"time"

	"github.com/google/uuid"
)

type SystemUserProfile struct {
	Id               uuid.UUID `db:"id"`
	Email            string    `db:"email"`
	Password         string    `db:"password"`
	LastLogin        time.Time `db:"last_login"`
	LastActive       time.Time `db:"last_active"`
	RegistrationTime time.Time `db:"registration_time"`
}
