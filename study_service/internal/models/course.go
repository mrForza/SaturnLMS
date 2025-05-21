package models

import "github.com/google/uuid"

type Course struct {
	Id          uuid.UUID   `db:"id"`
	Name        string      `db:"name"`
	Description string      `db:"description"`
	Formula     string      `db:"formula"`
	Languages   []string    `db:"languages"`
	Teachers    []uuid.UUID `db:"teachers"`
	Students    []uuid.UUID `db:"students"`
}
