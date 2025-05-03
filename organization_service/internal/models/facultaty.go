package models

type Facultaty struct {
	Name           string `db:"name"`
	Description    string `db:"description"`
	UniversityName string `db:"university_name"`
}
