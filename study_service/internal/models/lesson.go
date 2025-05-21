package models

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type Lesson struct {
	Id          uuid.UUID        `db:"id"`
	Name        string           `db:"name"`
	Description string           `db:"description"`
	Type        bool             `db:"type"`
	Files       []multipart.File `db:"files"`
	Homework    uuid.UUID        `db:"homework"`
}
