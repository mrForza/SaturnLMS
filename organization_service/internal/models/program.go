package models

type Program struct {
	Name          string `db:"name"`
	Description   string `db:"description"`
	Type          string `db:"type"`
	Languages     string `db:"languages"`
	FacultatyName string `db:"facultaty_name"`
}
