package models

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type Homework struct {
	Id          uuid.UUID        `db:"id"`
	Name        string           `db:"name"`
	Description string           `db:"description"`
	Files       []multipart.File `db:"files"`
}
