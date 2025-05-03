package models

import "github.com/google/uuid"

type University struct {
	Name          string    `db:"name"`
	Description   string    `db:"description"`
	LegalAddress  string    `db:"legal_address"`
	ActualAddress string    `db:"actual_address"`
	Inn           string    `db:"inn"`
	BankName      string    `db:"bank_name"`
	OwnerId       uuid.UUID `db:"owner_id"`
}
