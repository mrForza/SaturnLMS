package models

import "github.com/google/uuid"

type File struct {
	BucketId  uuid.UUID `db:"bucket_id"`
	Name      string    `db:"name"`
	Extension string    `db:"extension"`
}
